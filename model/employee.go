package model

import "gorm.io/gorm"

type Employee struct {
	gorm.Model
	Name                string
	Phone               string
	PositionID          uint
	Position            Position
	Inventories         []Inventory `gorm:"many2many:employee_inventories;"`
	EmployeeInventories []EmployeeInventories
}
