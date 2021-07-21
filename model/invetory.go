package model

import "gorm.io/gorm"

type Inventory struct {
	gorm.Model
	Name    string
	Number  string
	Archive Archive
}
