package main

import (
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/neerubhandari/BlogPortal/database"
)

func main() {
	database.Connect()

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	port := os.Getenv("PORT")
	// gin.SetMode(gin.ReleaseMode)
	httpRouter := gin.Default()
	httpRouter.Run(":3002" + port)
}
