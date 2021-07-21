package core

import (
	"GinBAsic/config"
	"GinBAsic/model"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetBlog(c *gin.Context) {
	blogData := []model.Blog{}
	userData := []model.User{}

	config.DB.Find(&blogData)

	// getUser := config.DB.Model(&userData).Association("Blogs").Find(&userData)
	getUser := config.DB.Preload("Blogs").Find(&userData)

	fmt.Println("getUser")
	fmt.Println(getUser)

	c.JSON(200, gin.H{
		"Message": "Welcome To Gin Framework",
		"data":    blogData,
	})
}

func GetBlogByID(c *gin.Context) {
	id := c.Param("id")
	var blogItem model.Blog

	// select * from blogs where id = id
	// DB.First(&blogItem, id)
	getData := config.DB.First(&blogItem, "id = ?", id)
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

func InsertBlog(c *gin.Context) {
	fmt.Println("jwt_user_id")
	userID := uint(c.MustGet("jwt_user_id").(float64))
	fmt.Println(userID)

	blog := model.Blog{
		Title:  c.PostForm("title"),
		Desc:   c.PostForm("desc"),
		Slug:   c.PostForm("slug"),
		UserID: userID,
	}

	config.DB.Create(&blog)

	c.JSON(http.StatusCreated, gin.H{
		"Status":  "Created",
		"Message": "Posting Success",
		"data":    blog,
	})
}

func UpdateBlog(c *gin.Context) {
	id := c.PostForm("id")

	jwtUserID := uint(c.MustGet("jwt_user_id").(float64))

	var blogItem model.Blog

	// get spesific data
	getData := config.DB.First(&blogItem, "id = ?", id)
	if getData.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"Status":  "StatusNotFound",
			"Message": "Data Not Found",
		})
		c.Abort()
		return
	}

	fmt.Println("jwtUserID")
	fmt.Println(jwtUserID)
	fmt.Println(blogItem.UserID)

	if blogItem.UserID != jwtUserID {
		c.JSON(http.StatusNotAcceptable, gin.H{
			"Status":  "Not Acceptable",
			"Message": "Your are not Owner this article",
		})
		c.Abort()
		return
	}

	// define in struct variable
	blogItem.Title = c.PostForm("title")
	blogItem.Desc = c.PostForm("desc")
	blogItem.Slug = c.PostForm("slug")

	// save / update
	config.DB.Save(&blogItem)

	c.JSON(http.StatusAccepted, gin.H{
		"Status":  "Updated",
		"Message": "Update Blog Success",
		"data":    blogItem,
	})
}

func DeleteBlog(c *gin.Context) {
	id := c.Param("id")
	var blogItem model.Blog

	// get spesific data
	getData := config.DB.First(&blogItem, "id = ?", id)
	if getData.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"Status":  "StatusNotFound",
			"Message": "Data Not Found",
		})
		c.Abort()
		return
	}

	// save / update
	config.DB.Delete(&blogItem)

	c.JSON(http.StatusAccepted, gin.H{
		"Status":  "Deleted",
		"Message": "Delete Blog Success",
	})
}
