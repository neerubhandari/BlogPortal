package database

import (
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/neerubhandari/BlogPortal/models"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("error to load .env file")
	}
	dsn := os.Getenv("DNS")
	database, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("Could not connect to the database")
	}
	DB = database
	database.AutoMigrate(&models.User{})
}
