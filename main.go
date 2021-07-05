package main

import (
	"GinBAsic/core"

	// "./core"

	"github.com/gin-gonic/gin"
)

func main() {
	// setup router
	router := gin.Default()

	// set endpoint
	v1 := router.Group("/api/v1")
	{
		v1.GET("/", getHome)
		blogs := v1.Group("/blog")
		{
			blogs.GET("/", core.GetBlog)
			blogs.GET("/:id", core.GetBlogByID)
			blogs.POST("/", core.InsertBlog)
			blogs.PUT("/", core.UpdateBlog)
			blogs.DELETE("/:id", core.DeleteBlog)
		}
		users := v1.Group("/user")
		{
			users.GET("/", core.GetUser)
			users.GET("/:id", core.GetUserByID)
			users.POST("/", core.InsertUser)
		}
	}

	// run server
	router.Run() // listen and serve in HTTP localhost:8080
}

func getHome(c *gin.Context) {
	c.JSON(200, gin.H{
		"Message": "Welcome To Gin Framework",
	})
}
