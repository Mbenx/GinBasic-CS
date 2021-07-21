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

func GetPosition(c *gin.Context) {
	PositionData := []model.Position{}

	// config.DB.Find(&PositionData)
	// config.DB.Preload("Department").Find(&PositionData)
	config.DB.Preload(clause.Associations).Find(&PositionData)

	c.JSON(200, gin.H{
		"Message": "Welcome To Gin Framework",
		"data":    PositionData,
	})
}

func GetPositionByID(c *gin.Context) {
	id := c.Param("id")
	var PositionItem model.Position

	// select * from Positions where id = id
	// DB.First(&PositionItem, id)
	getData := config.DB.First(&PositionItem, "id = ?", id)
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
		"Message": "Welcome to Our Positions",
		"data":    PositionItem,
	})
}

func InsertPosition(c *gin.Context) {
	fmt.Println("jwt_user_id")
	userID := uint(c.MustGet("jwt_user_id").(float64))
	fmt.Println(userID)

	DepartmentData := []model.Department{}

	u, _ := strconv.ParseUint(c.PostForm("department_id"), 10, 64)
	DepartmentID := uint(u)

	getDepartment := config.DB.First(&DepartmentData, "id = ?", DepartmentID)
	if getDepartment.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"Status":  "StatusNotFound",
			"Message": "Department Not Found",
		})
		c.Abort()
		return
	}

	Position := model.Position{
		Name:         c.PostForm("name"),
		Description:  c.PostForm("desc"),
		DepartmentID: DepartmentID,
	}

	config.DB.Create(&Position)

	c.JSON(http.StatusCreated, gin.H{
		"Status":  "Created",
		"Message": "Posting Success",
		"data":    Position,
	})
}

func UpdatePosition(c *gin.Context) {
	id := c.PostForm("id")

	var PositionItem model.Position

	// get spesific data
	getData := config.DB.First(&PositionItem, "id = ?", id)
	if getData.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"Status":  "StatusNotFound",
			"Message": "Data Not Found",
		})
		c.Abort()
		return
	}

	// define in struct variable
	PositionItem.Name = c.PostForm("name")
	PositionItem.Description = c.PostForm("description")

	// save / update
	config.DB.Save(&PositionItem)

	c.JSON(http.StatusAccepted, gin.H{
		"Status":  "Updated",
		"Message": "Update Position Success",
		"data":    PositionItem,
	})
}

func DeletePosition(c *gin.Context) {
	id := c.Param("id")
	var PositionItem model.Position

	// get spesific data
	getData := config.DB.First(&PositionItem, "id = ?", id)
	if getData.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"Status":  "StatusNotFound",
			"Message": "Data Not Found",
		})
		c.Abort()
		return
	}

	// save / update
	config.DB.Delete(&PositionItem)

	c.JSON(http.StatusAccepted, gin.H{
		"Status":  "Deleted",
		"Message": "Delete Position Success",
	})
}
