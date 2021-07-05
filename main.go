package main

import (
	"net/http"

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
			blogs.GET("/", getBlog)
			blogs.GET("/:id", getBlogByID)
			blogs.POST("/", insertBlog)
			blogs.PUT("/", updateBlog)
			blogs.DELETE("/:id", deleteBlog)
		}
		users := v1.Group("/user")
		{
			users.GET("/", getUser)
			users.GET("/:id", getUserByID)
			users.POST("/", insertUser)
			// users.POST("/", updateUser)
			// users.DELETE("/:id", deleteUser)
		}
	}

	v2 := router.Group("/api/v2")
	{
		v2.GET("/", getHomeV2)
	}

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

func getHomeV2(c *gin.Context) {
	c.JSON(200, gin.H{
		"Message": "this is Home in V2",
	})
}

func getBlog(c *gin.Context) {
	blogData := []Blog{}

	DB.Find(&blogData)

	c.JSON(200, gin.H{
		"Message": "Welcome To Gin Framework",
		"data":    blogData,
	})
}

func getBlogByID(c *gin.Context) {
	id := c.Param("id")
	var blogItem Blog

	// select * from blogs where id = id
	// DB.First(&blogItem, id)
	getData := DB.First(&blogItem, "id = ?", id)
	if getData.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"Status":  "StatusNotFound",
			"Message": "Data Not Found",
		})
		c.Abort()
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"Status":  "success",
		"Message": "Welcome to Our Blogs",
		"data":    blogItem,
	})
}

func insertBlog(c *gin.Context) {
	blog := Blog{
		Title: c.PostForm("title"),
		Desc:  c.PostForm("desc"),
		Slug:  c.PostForm("slug"),
	}

	DB.Create(&blog)

	c.JSON(http.StatusCreated, gin.H{
		"Status":  "Created",
		"Message": "Posting Success",
		"data":    blog,
	})
}

func updateBlog(c *gin.Context) {
	id := c.PostForm("id")
	var blogItem Blog

	// get spesific data
	getData := DB.First(&blogItem, "id = ?", id)
	if getData.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"Status":  "StatusNotFound",
			"Message": "Data Not Found",
		})
		c.Abort()
		return
	}

	// define in struct variable
	blogItem.Title = c.PostForm("title")
	blogItem.Desc = c.PostForm("desc")
	blogItem.Slug = c.PostForm("slug")

	// save / update
	DB.Save(&blogItem)

	c.JSON(http.StatusAccepted, gin.H{
		"Status":  "Updated",
		"Message": "Update Blog Success",
		"data":    blogItem,
	})
}

func deleteBlog(c *gin.Context) {
	id := c.Param("id")
	var blogItem Blog

	// get spesific data
	getData := DB.First(&blogItem, "id = ?", id)
	if getData.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"Status":  "StatusNotFound",
			"Message": "Data Not Found",
		})
		c.Abort()
		return
	}

	// save / update
	DB.Delete(&blogItem)

	c.JSON(http.StatusAccepted, gin.H{
		"Status":  "Deleted",
		"Message": "Delete Blog Success",
	})
}

func getUser(c *gin.Context) {
	userData := []User{}

	DB.Find(&userData)

	c.JSON(200, gin.H{
		"Status":  "success",
		"Message": "Welcome Our Users",
		"Data":    userData,
	})
}

func getUserByID(c *gin.Context) {
	id := c.Param("id")
	var userItem User

	// select * from user where id = id
	getData := DB.First(&userItem, "id = ?", id)
	if getData.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"Status":  "StatusNotFound",
			"Message": "Data Not Found",
		})
		c.Abort()
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"Status":  "success",
		"Message": "List User",
		"data":    userItem,
	})
}

func insertUser(c *gin.Context) {
	user := User{
		Username: c.PostForm("username"),
		Fullname: c.PostForm("fullname"),
		Email:    c.PostForm("email"),
		Address:  c.PostForm("address"),
	}

	DB.Create(&user)
	c.JSON(http.StatusCreated, gin.H{
		"Status":  "Created",
		"Message": "Insert User Success",
		"Data":    user,
	})
}

func postUser(c *gin.Context) {
	name := c.PostForm("name")
	email := c.PostForm("email")

	c.JSON(201, gin.H{
		"Status":  "Created",
		"Name":    name,
		"Email":   email,
		"Message": "Posting Success",
	})
}
