package controllers

import (
	"net/http"
	"rockcompanion/config"
	"rockcompanion/models"
	"rockcompanion/utils"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

func Register(c *gin.Context) {
	var body models.User
	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid body"})
		return
	}

	hashed, _ := bcrypt.GenerateFromPassword([]byte(body.Password), 14)
	user := models.User{Email: body.Email, Password: string(hashed)}

	if err := config.DB.Create(&user).Error; err != nil {
		c.JSON(http.StatusConflict, gin.H{"error": "User already exists"})
		return
	}

	token, _ := utils.GenerateToken(user.ID)
	c.JSON(http.StatusOK, gin.H{"user": user, "token": token})
}

func Login(c *gin.Context) {
	var body models.User
	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid login"})
		return
	}

	var user models.User
	if err := config.DB.Where("email = ?", body.Email).First(&user).Error; err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "No such user"})
		return
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(body.Password)); err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Wrong password"})
		return
	}

	token, _ := utils.GenerateToken(user.ID)
	c.JSON(http.StatusOK, gin.H{"user": user, "token": token})
}

func Logout(c *gin.Context) {
	// Client handles token removal; backend can optionally blacklist tokens (not shown here)
	c.JSON(http.StatusOK, gin.H{"message": "Logged out"})
}
