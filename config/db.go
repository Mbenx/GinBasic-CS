package config

import (
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

	// Get generic database object sql.DB to use its functions
	sqlDB, err := DB.DB()
	if err != nil {
		panic("failed Get generic database object")
	}

	// Close
	defer sqlDB.Close()

	// run migration
	DB.AutoMigrate(&Blog{})
	DB.AutoMigrate(&User{})
}
