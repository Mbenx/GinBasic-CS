package core

import (
	"GinBAsic/config"
	"GinBAsic/model"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetBlog(c *gin.Context) {
	blogData := []model.Blog{}

	config.DB.Find(&blogData)

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
	blog := model.Blog{
		Title: c.PostForm("title"),
		Desc:  c.PostForm("desc"),
		Slug:  c.PostForm("slug"),
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
