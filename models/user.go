package models

import (
	"github.com/google/uuid"
)

type User struct {
	Id        uuid.UUID `gorm:"type:uuid;"`
	Username  string    `gorm:"unique" json:"username"`
	Firstname string    `json:"firstname"`
	Email     string    `gorm:"unique" json:"email"`
	Password  string    `json:"password"`
}

type UserRegistrationRequest struct {
	Username  string `json:"username" binding:"required"`
	Firstname string `json:"firstname" binding:"required"`
	Email     string `json:"email" binding:"required,email"`
	Password  string `json:"password" binding:"required,min=6"`
}
