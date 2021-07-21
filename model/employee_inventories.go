package model

import "gorm.io/gorm"

type EmployeeInventories struct {
	gorm.Model
	EmployeeID  uint
	InventoryID uint
	Description string
	Date        string
}
