package core

import (
	"GinBAsic/config"
	"GinBAsic/model"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm/clause"
)

func GetInventory(c *gin.Context) {
	InventoryData := []model.Inventory{}

	// config.DB.Find(&InventoryData)
	// config.DB.Preload("Department").Find(&InventoryData)
	config.DB.Preload(clause.Associations).Find(&InventoryData)

	c.JSON(200, gin.H{
		"Message": "Welcome To Gin Framework",
		"data":    InventoryData,
	})
}

func GetInventoryByID(c *gin.Context) {
	id := c.Param("id")
	var InventoryItem model.Inventory

	// select * from Inventorys where id = id
	// DB.First(&InventoryItem, id)
	getData := config.DB.First(&InventoryItem, "id = ?", id)
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
		"Message": "Welcome to Our Inventorys",
		"data":    InventoryItem,
	})
}

func InsertInventory(c *gin.Context) {
	fmt.Println("jwt_user_id")
	userID := uint(c.MustGet("jwt_user_id").(float64))
	fmt.Println(userID)

	Inventory := model.Inventory{
		Name:   c.PostForm("name"),
		Number: c.PostForm("number"),
	}

	config.DB.Create(&Inventory)

	lastInventory := model.Inventory{}

	config.DB.Where("number = ? ", Inventory.Number).First(&lastInventory)

	Archive := model.Archive{
		Name:        c.PostForm("archive_name"),
		Description: c.PostForm("archive_description"),
		InventoryID: lastInventory.ID,
	}

	config.DB.Create(&Archive)

	c.JSON(http.StatusCreated, gin.H{
		"Status":  "Created",
		"Message": "Posting Success",
		"data":    Inventory,
	})
}

func UpdateInventory(c *gin.Context) {
	id := c.PostForm("id")

	var InventoryItem model.Inventory

	// get spesific data
	getData := config.DB.First(&InventoryItem, "id = ?", id)
	if getData.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"Status":  "StatusNotFound",
			"Message": "Data Not Found",
		})
		c.Abort()
		return
	}

	// define in struct variable
	InventoryItem.Name = c.PostForm("name")
	InventoryItem.Number = c.PostForm("number")

	// save / update
	config.DB.Save(&InventoryItem)

	c.JSON(http.StatusAccepted, gin.H{
		"Status":  "Updated",
		"Message": "Update Inventory Success",
		"data":    InventoryItem,
	})
}

func DeleteInventory(c *gin.Context) {
	id := c.Param("id")
	var InventoryItem model.Inventory

	// get spesific data
	getData := config.DB.First(&InventoryItem, "id = ?", id)
	if getData.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"Status":  "StatusNotFound",
			"Message": "Data Not Found",
		})
		c.Abort()
		return
	}

	// save / update
	config.DB.Delete(&InventoryItem)

	c.JSON(http.StatusAccepted, gin.H{
		"Status":  "Deleted",
		"Message": "Delete Inventory Success",
	})
}
