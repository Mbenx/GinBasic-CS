package core

import "github.com/gin-gonic/gin"

func GetHome(c *gin.Context) {
	c.JSON(200, gin.H{
		"Message": "Welcome To Gin Framework",
	})
}
