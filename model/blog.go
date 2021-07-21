package model

import "gorm.io/gorm"

type Blog struct {
	gorm.Model
	Title  string `sql:"type:text`
	Slug   string `gorm:"unique_index"`
	Desc   string `sql:"type:text`
	UserID uint
}
