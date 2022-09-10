package database

import (
	"go-server-jwt/models"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect() {
	connection, err := gorm.Open(mysql.Open("root:DucDV@31@/gin-server-jwt-db"), &gorm.Config{})

	if err != nil {
		panic(any("Could not connect to the database"))
	}

	DB = connection

	err = connection.AutoMigrate(&models.User{})
	if err != nil {
		return
	}
}
