package main

import (
	"GinBAsic/config"
	"GinBAsic/core"
	"GinBAsic/midleware"

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

		v1.GET("/checkToken", midleware.IsAuth()) // IsAuth

		auth := v1.Group("/auth")
		{
			auth.GET("/", core.IndexHandler)
			auth.GET("/:provider", core.RedirectHandler)
			auth.GET("/:provider/callback", core.CallbackHandler)
			auth.POST("/login", core.Login)
			auth.POST("/register", core.Register)
		}

		blogs := v1.Group("/blog")
		{
			blogs.GET("/", midleware.IsAuth(), core.GetBlog)           // IsAuth
			blogs.GET("/:id", midleware.IsAuth(), core.GetBlogByID)    // IsAuth
			blogs.POST("/", midleware.IsUser(), core.InsertBlog)       // IsUser
			blogs.PUT("/", midleware.IsAdmin(), core.UpdateBlog)       // IsAdmin || user_id created
			blogs.DELETE("/:id", midleware.IsAdmin(), core.DeleteBlog) // IsAdmin || user_id created
		}
		users := v1.Group("/user")
		{
			users.GET("/", midleware.IsAdmin(), core.GetUser)        // IsAdmin
			users.GET("/:id", midleware.IsAdmin(), core.GetUserByID) // IsAdmin || user_id spesific
			users.POST("/", midleware.IsAdmin(), core.InsertUser)    // IsAdmin
		}
	}

	// run server
	router.Run() // listen and serve in HTTP localhost:8080
}
