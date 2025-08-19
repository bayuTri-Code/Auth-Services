package models


import (
	"github.com/google/uuid"
)


type RegisterRequest struct {
	ID       uuid.UUID `gorm:"type:char(36);primary_key" json:"id"`
	Name     string    `json:"name" binding:"required"`
	Email    string    `json:"email" binding:"required"`
	Password string    `json:"password" binding:"required"`
}

type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type RegisterResponse struct {
	ID    uuid.UUID `json:"id"`
	Name  string    `json:"name"`
	Email string    `json:"email"`
}

type LoginResponse struct {
	Token string           `json:"token"`
	User  RegisterResponse `json:"user"`
}
