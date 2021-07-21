package model

import "gorm.io/gorm"

type Department struct {
	gorm.Model
	Name     string
	Code     string
	Position []Position
}
