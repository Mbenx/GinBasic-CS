package main

import (
	"GinBAsic/config"
	"GinBAsic/core"

	// "./core"

	"github.com/gin-gonic/gin"
	"github.com/subosito/gotenv"
)

func main() {

	// connect DB
	config.InitDB()

	// Get generic database object sql.DB to use its functions
	sqlDB, err := config.DB.DB()
	if err != nil {
		panic("failed Get generic database object")
	}

	// Close
	defer sqlDB.Close()

	gotenv.Load()

	// setup router
	router := gin.Default()

	// set endpoint
	v1 := router.Group("/api/v1")
	{
		v1.GET("/", core.GetHome)

		auth := v1.Group("/auth")
		{
			auth.GET("/", core.IndexHandler)
			auth.GET("/:provider", core.RedirectHandler)
			auth.GET("/:provider/callback", core.CallbackHandler)
		}

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
