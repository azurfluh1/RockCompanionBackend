package main

import (
	"os"
	"rockcompanion/config"
	"rockcompanion/middleware"
	"rockcompanion/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	// err := godotenv.Load()
	// if os.Getenv("FLY_REGION") == "" {
	// 	err := godotenv.Load()
	// 	if err != nil {
	// 		log.Fatal("Error loading .env file")
	// 	}
	// } else {
	os.Getenv("DB_HOST")
	// }
	config.ConnectDB()

	r := gin.Default()
	r.Use(middleware.CORSMiddleware())
	api := r.Group("/api")
	routes.AuthRoutes(api)

	r.Run(":8080")
}
