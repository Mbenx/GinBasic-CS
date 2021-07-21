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
	DB.AutoMigrate(&model.User{})
	DB.AutoMigrate(&model.Blog{})

	DB.AutoMigrate(&model.Department{})
	DB.AutoMigrate(&model.Position{})
	DB.AutoMigrate(&model.Employee{})
	DB.AutoMigrate(&model.Inventory{})
	DB.AutoMigrate(&model.Archive{})
	DB.AutoMigrate(&model.EmployeeInventories{})

	// DB.Table("employee_inventories").AddForeignKey("employee_id", "employees(id)", "CASCADE", "CASCADE")
	// DB.Table("employee_inventories").AddForeignKey("inventory_id", "inventories(id)", "CASCADE", "CASCADE")

	// DB.Migrator().CurrentDatabase()
	// DB.Migrator().CreateConstraint(&model.User{}, "Blogs")
	// DB.Migrator().CreateConstraint(&model.User{}), "fk_users_blogs")

	// DB.Model(&model.Blog{}).AddForeignKey("user_id", "users(id)", "CASCADE", "CASCADE")
	// DB.Model(&model.User{}).Related(&model.Blog{})

}
