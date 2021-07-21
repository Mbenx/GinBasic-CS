package model

import "gorm.io/gorm"

type Archive struct {
	gorm.Model
	Name        string
	Description string
	InventoryID uint
}
