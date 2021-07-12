package midleware

import (
	"fmt"
	"net/http"
	"os"
	"strings"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

func IsAuth() gin.HandlerFunc {
	return CheckJWT(9)
}

func IsAdmin() gin.HandlerFunc {
	return CheckJWT(0)
}

func IsUser() gin.HandlerFunc {
	return CheckJWT(1)
}

func CheckJWT(role uint) gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.Request.Header.Get("Authorization")
		splitToken := strings.Split(authHeader, " ")

		if len(splitToken) == 2 {
			token, err := jwt.Parse(splitToken[1], func(token *jwt.Token) (interface{}, error) {
				return []byte(os.Getenv("JWT_SECRET")), nil
			})

			if token.Valid {
				claim := token.Claims.(jwt.MapClaims)

				user_id := uint(claim["user_id"].(float64))
				role_id := uint(claim["user_role"].(float64))

				fmt.Println("Role")
				fmt.Println(role_id)
				fmt.Println(role)

				if role == role_id {
					fmt.Println("Role as Admin")
					c.Set("jwt_user_id", user_id)
					c.Set("jwt_role_id", role_id)
				} else if role == 9 {
					c.Set("jwt_user_id", user_id)
					c.Set("jwt_role_id", role_id)
				} else {
					c.JSON(http.StatusUnprocessableEntity, gin.H{
						"Status":  "UnprocessableEntity",
						"Message": "Your Role Cant Allowed",
					})
					c.Abort()
					return
				}
				fmt.Println("You look nice today")
			} else if ve, ok := err.(*jwt.ValidationError); ok {
				if ve.Errors&jwt.ValidationErrorMalformed != 0 {
					c.JSON(http.StatusUnauthorized, gin.H{
						"Status":  "Unauthorized",
						"Message": "That's not even a token",
					})
					fmt.Println("That's not even a token")
					c.Abort()
					return
				} else if ve.Errors&(jwt.ValidationErrorExpired|jwt.ValidationErrorNotValidYet) != 0 {
					// Token is either expired or not active yet
					c.JSON(http.StatusUnauthorized, gin.H{
						"Status":  "Unauthorized",
						"Message": "Token is either expired or not active yet",
					})

					fmt.Println("Token is either expired or not active yet")

					c.Abort()
					return
				} else {
					c.JSON(http.StatusUnauthorized, gin.H{
						"Status":  "Unauthorized",
						"Message": "Token Invalid",
					})
					fmt.Println("Couldn't handle this token:", err)

					c.Abort()
					return
				}
			} else {
				c.JSON(http.StatusUnauthorized, gin.H{
					"Status":  "Unauthorized",
					"Message": "Token Invalid",
				})

				fmt.Println("Couldn't handle this token:", err)

				c.Abort()
				return
			}

		} else {
			c.JSON(http.StatusUnauthorized, gin.H{
				"Status":  "Unauthorized",
				"Message": "Unauthorized",
			})
			c.Abort()
			return
		}

	}

}
