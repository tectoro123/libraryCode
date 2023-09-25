package db

import (
	"BookStore/models"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var Connection *gorm.DB

func ConnectToDB() {
	var dsn = "root:Tectoro@123@tcp(127.0.0.1:3306)/join?charset=utf8mb4&parseTime=True&loc=Local"
	Connection, _ = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	Connection.AutoMigrate(&models.Book{})
}
