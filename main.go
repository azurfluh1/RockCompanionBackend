package main

import (
	"log"
	"rockcompanion/config"
	"rockcompanion/routes"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load("/app/.env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	config.ConnectDB()

	r := gin.Default()
	api := r.Group("/api")
	routes.AuthRoutes(api)

	r.Run(":8080")
}
