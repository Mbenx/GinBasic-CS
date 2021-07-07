package config

import (
	"GinBAsic/model"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDB() {
	var err error
	dsn := "root:h3ru@mysql@tcp(127.0.0.1:3306)/golang1?charset=utf8mb4&parseTime=True&loc=Local"
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		panic("failed to connect database")
	}

	// run migration
	DB.AutoMigrate(&model.Blog{})
	DB.AutoMigrate(&model.User{})
}
