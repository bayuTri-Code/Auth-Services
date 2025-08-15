package services

import (
	"context"
	"errors"
	"fmt"
	"strings"

	"github.com/bayuTri-Code/Auth-Services/database"
	"github.com/bayuTri-Code/Auth-Services/internal/models"
	"github.com/google/uuid"
	"gorm.io/gorm"
)


func RegisterServices(ctx context.Context, Name, Email, Password string) (*models.AuthData, error){
	db := database.Db
	Email = strings.ToLower(strings.TrimSpace(Email))
	
	var existing models.AuthData
	if err := db.WithContext(ctx).Where("email = ?", Email).First(&existing).Error; err == nil {
		return nil, errors.New("user already exists")
	}

	registerUser := &models.AuthData{
		ID: uuid.New(),
		Name: Name,
		Email: Email,
		Password: Password,
	}
	if err := db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		if err := tx.Create(registerUser).Error; err != nil {
			return fmt.Errorf("failed to create user: %v", err)
		}
		return nil
	}); err != nil {
		return nil, fmt.Errorf("transaction failed: %v", err)
	}
	return registerUser, nil
}