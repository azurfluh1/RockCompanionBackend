package main

import (
	"log"
	"rockcompanion/config"
	"rockcompanion/middleware"
	"rockcompanion/routes"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	config.ConnectDB()

	r := gin.Default()
	r.Use(middleware.CORSMiddleware())
	api := r.Group("/api")
	routes.AuthRoutes(api)

	r.Run(":8080")
}
