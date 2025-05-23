package utils

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
)

var jwtKey = []byte("your_secret_key")

func GenerateToken(userID uint) (string, error) {
	claims := jwt.MapClaims{
		"user_id": userID,
		"exp":     time.Now().Add(24 * time.Hour).Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(jwtKey)
}
