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
			blogs.PUT("/", midleware.IsUser(), core.UpdateBlog)        // IsAdmin || user_id created
			blogs.DELETE("/:id", midleware.IsAdmin(), core.DeleteBlog) // IsAdmin || user_id created
		}
		users := v1.Group("/user")
		{
			users.GET("/", midleware.IsAdmin(), core.GetUser)        // IsAdmin
			users.GET("/:id", midleware.IsAdmin(), core.GetUserByID) // IsAdmin || user_id spesific
			users.POST("/", midleware.IsAdmin(), core.InsertUser)    // IsAdmin
		}
	}

	hr := router.Group("/api/hr")
	{
		department := hr.Group("/department")
		{
			department.GET("/", midleware.IsAdmin(), core.GetDepartment)
			department.GET("/:id", midleware.IsAdmin(), core.GetDepartmentByID)
			department.POST("/", midleware.IsAdmin(), core.InsertDepartment)
			department.PUT("/", midleware.IsAdmin(), core.UpdateDepartment)
			department.DELETE("/:id", midleware.IsAdmin(), core.DeleteDepartment)
		}

		position := hr.Group("/position")
		{
			position.GET("/", midleware.IsAdmin(), core.GetPosition)
			position.GET("/:id", midleware.IsAdmin(), core.GetPositionByID)
			position.POST("/", midleware.IsAdmin(), core.InsertPosition)
			position.PUT("/", midleware.IsAdmin(), core.UpdatePosition)
			position.DELETE("/:id", midleware.IsAdmin(), core.DeletePosition)
		}

		employee := hr.Group("/employee")
		{
			employee.GET("/", midleware.IsAdmin(), core.GetEmployee)
			employee.GET("/:id", midleware.IsAdmin(), core.GetEmployeeByID)
			employee.POST("/", midleware.IsAdmin(), core.InsertEmployee)
			employee.PUT("/", midleware.IsAdmin(), core.UpdateEmployee)
			employee.DELETE("/:id", midleware.IsAdmin(), core.DeleteEmployee)
		}

		inventory := hr.Group("/inventory")
		{
			inventory.GET("/", midleware.IsAdmin(), core.GetInventory)
			inventory.GET("/:id", midleware.IsAdmin(), core.GetInventoryByID)
			inventory.POST("/", midleware.IsAdmin(), core.InsertInventory)
			inventory.PUT("/", midleware.IsAdmin(), core.UpdateInventory)
			inventory.DELETE("/:id", midleware.IsAdmin(), core.DeleteInventory)
		}
	}

	// run server
	router.Run() // listen and serve in HTTP localhost:8080
}
