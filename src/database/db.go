package database

import (
	"log"

	"github.com/OrlandoRomo/go-ambassador/src/model"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

// Connect creates a new mysql instance
// TODO: Use env variables instead of hardcoding the database credentials
func Connect() {
	var err error
	DB, err = gorm.Open(mysql.Open("root:root@tcp(db:3306)/ambassador"), &gorm.Config{})
	if err != nil {
		log.Println("could not connect to db", "status", "ERROR")
		return
	}
}

// AutoMigrate migrates the models
func AutoMigrate() {
	DB.AutoMigrate(model.User{}, model.Product{})
}
