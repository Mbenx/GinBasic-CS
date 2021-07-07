package core

import (
	"fmt"
	"net/http"
	"os"

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
	user, token, err := gocial.Handle(state, code)
	if err != nil {
		c.Writer.Write([]byte("Error: " + err.Error()))
		return
	}

	// Print in terminal user information
	fmt.Printf("%#v", token)
	fmt.Printf("%#v", user)
	fmt.Printf("%#v", provider)

	// If no errors, show provider name
	c.Writer.Write([]byte("Hi, " + user.FullName))
}
