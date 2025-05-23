package routes

import (
	"rockcompanion/controllers"

	"github.com/gin-gonic/gin"
)

func AuthRoutes(r *gin.RouterGroup) {
	r.POST("/login", controllers.Login)
	r.POST("/register", controllers.Register)
	r.POST("/logout", controllers.Logout)
}
