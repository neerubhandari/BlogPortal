package main

import (
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/neerubhandari/BlogPortal/database"
	"github.com/neerubhandari/BlogPortal/routes"
)

func main() {
	database.Connect()

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	port := os.Getenv("PORT")
	if port == "" {
		port = "3002" // default port if not set
	}
	router := routes.SetupRouter()

	router.Run(":" + port)
}
