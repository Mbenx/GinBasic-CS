package core

import (
	"GinBAsic/config"
	"GinBAsic/model"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetDepartment(c *gin.Context) {
	DepartmentData := []model.Department{}

	// config.DB.Find(&DepartmentData)

	config.DB.Preload("Position").Find(&DepartmentData)
	// config.DB.Preload(clause.Associations).Find(&DepartmentData)

	c.JSON(200, gin.H{
		"Message": "Welcome To Gin Framework",
		"data":    DepartmentData,
	})
}

func GetDepartmentByID(c *gin.Context) {
	id := c.Param("id")
	var DepartmentItem model.Department

	// select * from Departments where id = id
	// DB.First(&DepartmentItem, id)
	getData := config.DB.First(&DepartmentItem, "id = ?", id)
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
		"Message": "Welcome to Our Departments",
		"data":    DepartmentItem,
	})
}

func InsertDepartment(c *gin.Context) {
	fmt.Println("jwt_user_id")
	userID := uint(c.MustGet("jwt_user_id").(float64))
	fmt.Println(userID)

	Department := model.Department{
		Name: c.PostForm("name"),
		Code: c.PostForm("code"),
	}

	config.DB.Create(&Department)

	c.JSON(http.StatusCreated, gin.H{
		"Status":  "Created",
		"Message": "Posting Success",
		"data":    Department,
	})
}

func UpdateDepartment(c *gin.Context) {
	id := c.PostForm("id")

	var DepartmentItem model.Department

	// get spesific data
	getData := config.DB.First(&DepartmentItem, "id = ?", id)
	if getData.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"Status":  "StatusNotFound",
			"Message": "Data Not Found",
		})
		c.Abort()
		return
	}

	// define in struct variable
	DepartmentItem.Name = c.PostForm("name")
	DepartmentItem.Code = c.PostForm("code")

	// save / update
	config.DB.Save(&DepartmentItem)

	c.JSON(http.StatusAccepted, gin.H{
		"Status":  "Updated",
		"Message": "Update Department Success",
		"data":    DepartmentItem,
	})
}

func DeleteDepartment(c *gin.Context) {
	id := c.Param("id")
	var DepartmentItem model.Department

	// get spesific data
	getData := config.DB.First(&DepartmentItem, "id = ?", id)
	if getData.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"Status":  "StatusNotFound",
			"Message": "Data Not Found",
		})
		c.Abort()
		return
	}

	// save / update
	config.DB.Delete(&DepartmentItem)

	c.JSON(http.StatusAccepted, gin.H{
		"Status":  "Deleted",
		"Message": "Delete Department Success",
	})
}
