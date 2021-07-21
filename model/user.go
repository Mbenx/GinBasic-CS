package model

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Blog     []Blog
	Username string `gorm:"unique_index"`
	Fullname string
	Email    string
	Password string
	Address  string
	SocialID string
	Provider string
	Role     uint8 `gorm:"default:0"`
}
