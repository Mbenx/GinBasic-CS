package core

import (
	"GinBAsic/config"
	"GinBAsic/model"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetUser(c *gin.Context) {
	userData := []model.User{}

	config.DB.Find(&userData)

	c.JSON(200, gin.H{
		"Status":  "success",
		"Message": "Welcome Our Users",
		"Data":    userData,
	})
}

func GetUserByID(c *gin.Context) {
	id := c.Param("id")
	var userItem model.User

	// select * from user where id = id
	getData := config.DB.First(&userItem, "id = ?", id)
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

func InsertUser(c *gin.Context) {
	user := model.User{
		Username: c.PostForm("username"),
		Fullname: c.PostForm("fullname"),
		Email:    c.PostForm("email"),
		Address:  c.PostForm("address"),
	}

	config.DB.Create(&user)
	c.JSON(http.StatusCreated, gin.H{
		"Status":  "Created",
		"Message": "Insert User Success",
		"Data":    user,
	})
}
