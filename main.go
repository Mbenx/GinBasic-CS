package main

import (
	"GinBAsic/core"

	// "./core"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

type Blog struct {
	gorm.Model
	Title string `sql:"type:text`
	Slug  string `gorm:"unique_index"`
	Desc  string `sql:"type:text`
}

type User struct {
	gorm.Model
	Username string `gorm:"unique_index"`
	Fullname string
	Email    string
	Address  string
}

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

	// v2 := router.Group("/api/v2")
	// {
	// 	v2.GET("/", getHomeV2)
	// }

	var err error
	dsn := "root:h3ru@mysql@tcp(127.0.0.1:3306)/golang1?charset=utf8mb4&parseTime=True&loc=Local"
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		panic("failed to connect database")
	}

	// Get generic database object sql.DB to use its functions
	sqlDB, err := DB.DB()
	if err != nil {
		panic("failed Get generic database object")
	}

	// Close
	defer sqlDB.Close()

	// run migration
	DB.AutoMigrate(&Blog{})
	DB.AutoMigrate(&User{})

	// run server
	router.Run() // listen and serve in HTTP localhost:8080
}

func getHome(c *gin.Context) {
	c.JSON(200, gin.H{
		"Message": "Welcome To Gin Framework",
	})
}

// func getHomeV2(c *gin.Context) {
// 	c.JSON(200, gin.H{
// 		"Message": "this is Home in V2",
// 	})
// }
