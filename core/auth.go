package core

import (
	"GinBAsic/config"
	"GinBAsic/model"
	"crypto/sha256"
	"encoding/base64"
	"fmt"
	"net/http"
	"os"
	"strconv"
	"time"

	gocialStruct "github.com/danilopolani/gocialite/structs"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"gopkg.in/danilopolani/gocialite.v1"
)

var gocial = gocialite.NewDispatcher()

// Show homepage with login URL
func IndexHandler(c *gin.Context) {
	c.Writer.Write([]byte("<html><head><title>Gocialite example</title></head><body>" +
		"<a href='/auth/github'><button>Login with GitHub</button></a><br>" +
		"<a href='/auth/linkedin'><button>Login with LinkedIn</button></a><br>" +
		"<a href='/auth/facebook'><button>Login with Facebook</button></a><br>" +
		"<a href='/auth/google'><button>Login with Google</button></a><br>" +
		"<a href='/auth/bitbucket'><button>Login with Bitbucket</button></a><br>" +
		"<a href='/auth/amazon'><button>Login with Amazon</button></a><br>" +
		"<a href='/auth/amazon'><button>Login with Slack</button></a><br>" +
		"</body></html>"))
}

// Redirect to correct oAuth URL
func RedirectHandler(c *gin.Context) {
	// Retrieve provider from route
	provider := c.Param("provider")

	fmt.Println("AUTH_URL")
	fmt.Println(os.Getenv("AUTH_URL"))

	// In this case we use a map to store our secrets, but you can use dotenv or your framework configuration
	// for example, in revel you could use revel.Config.StringDefault(provider + "_clientID", "") etc.
	providerSecrets := map[string]map[string]string{
		"github": {
			"clientID":     os.Getenv("CLIENT_ID_GITHUB"),
			"clientSecret": os.Getenv("CLIENT_SECRET_GITHUB"),
			"redirectURL":  os.Getenv("AUTH_URL") + "github/callback",
		},
		"google": {
			"clientID":     os.Getenv("CLIENT_ID_GOOGLE"),
			"clientSecret": os.Getenv("CLIENT_SECRET_GOOGLE"),
			"redirectURL":  os.Getenv("AUTH_URL") + "google/callback",
		},
	}

	providerScopes := map[string][]string{
		"github": []string{"public_repo"},
		"google": []string{},
	}

	providerData := providerSecrets[provider]
	actualScopes := providerScopes[provider]
	authURL, err := gocial.New().
		Driver(provider).
		Scopes(actualScopes).
		Redirect(
			providerData["clientID"],
			providerData["clientSecret"],
			providerData["redirectURL"],
		)

	// Check for errors (usually driver not valid)
	if err != nil {
		c.Writer.Write([]byte("Error: " + err.Error()))
		return
	}

	// Redirect with authURL
	c.Redirect(http.StatusFound, authURL)
}

// Handle callback of provider
func CallbackHandler(c *gin.Context) {
	// Retrieve query params for state and code
	state := c.Query("state")
	code := c.Query("code")
	provider := c.Param("provider")

	// Handle callback and check for errors
	gocialUser, _, err := gocial.Handle(state, code)
	if err != nil {
		c.Writer.Write([]byte("Error: " + err.Error()))
		return
	}

	// check User
	dataUser, msg := SocialLoginOrRegister(provider, gocialUser)
	jwtToken := createToken(&dataUser)

	c.JSON(200, gin.H{
		"Message":  msg,
		"User":     dataUser,
		"Token":    jwtToken,
		"Provider": provider,
	})
}

func SocialLoginOrRegister(provider string, gocialUser *gocialStruct.User) (user model.User, msg string) {
	var dataUser model.User

	// config.DB.First(&dataUser, "provider = ? AND social_id = ?", provider, gocialUser.ID)
	config.DB.Where("provider = ? AND social_id = ?", provider, gocialUser.ID).First(&dataUser)

	fmt.Println("dataUser")
	fmt.Println(dataUser)

	if dataUser.ID == 0 {
		newUser := model.User{
			Username: gocialUser.Username,
			Fullname: gocialUser.FullName,
			Email:    gocialUser.Email,
			SocialID: gocialUser.ID,
			Provider: provider,
			Role:     1,
		}

		config.DB.Create(&newUser)

		msg = "Register Success"
		return newUser, msg
	} else {
		msg = "Login Success"
		return dataUser, msg
	}
}

func createToken(user *model.User) string {
	// jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{})

	KEY := []byte(os.Getenv("JWT_SECRET"))

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id":        user.ID,
		"user_social_id": user.SocialID,
		"user_role":      user.Role,
		"exp":            time.Now().AddDate(0, 0, 1).Unix(),
		"iat":            time.Now().Unix(),
	})

	tokenString, err := token.SignedString(KEY)
	fmt.Printf("%v %v", tokenString, err)

	return tokenString
}

func Login(c *gin.Context) {
	username := c.PostForm("username")
	password := c.PostForm("password")

	hashPass := encryptPassword(password)

	fmt.Println("hashPass")
	fmt.Println(hashPass)

	var dataUser model.User

	getData := config.DB.First(&dataUser, "username = ? AND password = ?", username, hashPass)
	if getData.Error != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"Status":  "Unauthorized",
			"Message": "Data Not Found, please check your Username & Password",
		})
		c.Abort()
		return
	}

	jwtToken := createToken(&dataUser)

	c.JSON(200, gin.H{
		"Message": "Login Berhasil",
		"User":    dataUser,
		"Token":   jwtToken,
	})
}

func Register(c *gin.Context) {
	username := c.PostForm("username")
	password := c.PostForm("password")
	email := c.PostForm("email")

	hashPass := encryptPassword(password)
	u, _ := strconv.ParseUint(c.PostForm("role"), 10, 64)
	var dataUser model.User

	config.DB.Where("username = ? OR email = ?", username, email).First(&dataUser)
	if dataUser.ID == 0 {
		newUser := model.User{
			Username: username,
			Fullname: c.PostForm("fullname"),
			Email:    email,
			Address:  c.PostForm("address"),
			Password: hashPass,
			Role:     uint8(u),
		}

		config.DB.Create(&newUser)

		jwtToken := createToken(&newUser)

		c.JSON(200, gin.H{
			"Message": "Register Berhasil",
			"User":    newUser,
			"Token":   jwtToken,
		})
	} else {
		c.JSON(http.StatusUnprocessableEntity, gin.H{
			"Status":  "Unprocessable Entity",
			"Message": "YourUsername or Password is already exist",
		})
		c.Abort()
		return
	}
}

func encryptPassword(password string) string {
	h := sha256.Sum256([]byte(password))
	return base64.StdEncoding.EncodeToString(h[:])
}
