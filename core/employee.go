package core

import (
	"GinBAsic/config"
	"GinBAsic/model"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm/clause"
)

func GetEmployee(c *gin.Context) {
	EmployeeData := []model.Employee{}

	// config.DB.Find(&EmployeeData)
	// config.DB.Preload("Department").Find(&EmployeeData)
	config.DB.Preload(clause.Associations).Find(&EmployeeData)

	c.JSON(200, gin.H{
		"Message": "Welcome To Gin Framework",
		"data":    EmployeeData,
	})
}

func GetEmployeeByID(c *gin.Context) {
	id := c.Param("id")
	var EmployeeItem model.Employee

	// select * from Employees where id = id
	// DB.First(&EmployeeItem, id)
	getData := config.DB.First(&EmployeeItem, "id = ?", id)
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
		"Message": "Welcome to Our Employees",
		"data":    EmployeeItem,
	})
}

func InsertEmployee(c *gin.Context) {
	fmt.Println("jwt_user_id")
	userID := uint(c.MustGet("jwt_user_id").(float64))
	fmt.Println(userID)

	PositionData := []model.Position{}

	u, _ := strconv.ParseUint(c.PostForm("position_id"), 10, 64)
	PositionID := uint(u)

	getPosition := config.DB.First(&PositionData, "id = ?", PositionID)
	if getPosition.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"Status":  "StatusNotFound",
			"Message": "Position Not Found",
		})
		c.Abort()
		return
	}

	Employee := model.Employee{
		Name:       c.PostForm("name"),
		Phone:      c.PostForm("phone"),
		PositionID: PositionID,
	}

	config.DB.Create(&Employee)

	c.JSON(http.StatusCreated, gin.H{
		"Status":  "Created",
		"Message": "Posting Success",
		"data":    Employee,
	})
}

func UpdateEmployee(c *gin.Context) {
	id := c.PostForm("id")

	var EmployeeItem model.Employee

	// get spesific data
	getData := config.DB.First(&EmployeeItem, "id = ?", id)
	if getData.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"Status":  "StatusNotFound",
			"Message": "Data Not Found",
		})
		c.Abort()
		return
	}

	// define in struct variable
	EmployeeItem.Name = c.PostForm("name")
	EmployeeItem.Phone = c.PostForm("phone")

	// save / update
	config.DB.Save(&EmployeeItem)

	c.JSON(http.StatusAccepted, gin.H{
		"Status":  "Updated",
		"Message": "Update Employee Success",
		"data":    EmployeeItem,
	})
}

func DeleteEmployee(c *gin.Context) {
	id := c.Param("id")
	var EmployeeItem model.Employee

	// get spesific data
	getData := config.DB.First(&EmployeeItem, "id = ?", id)
	if getData.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"Status":  "StatusNotFound",
			"Message": "Data Not Found",
		})
		c.Abort()
		return
	}

	// save / update
	config.DB.Delete(&EmployeeItem)

	c.JSON(http.StatusAccepted, gin.H{
		"Status":  "Deleted",
		"Message": "Delete Employee Success",
	})
}
