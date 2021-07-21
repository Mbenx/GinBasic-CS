package model

import "gorm.io/gorm"

type Position struct {
	gorm.Model
	Name         string
	Description  string
	DepartmentID uint
	Department   Department
	Employee     []Employee
}
