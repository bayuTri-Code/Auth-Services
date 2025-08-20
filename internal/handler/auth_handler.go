package handler

import (
	"net/http"

	"github.com/bayuTri-Code/Auth-Services/internal/models"
	"github.com/bayuTri-Code/Auth-Services/internal/services"
	"github.com/bayuTri-Code/Auth-Services/internal/utils"
	"github.com/gin-gonic/gin"
)

func RegisterHandler(c *gin.Context) {
	var req models.RegisterRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		utils.ResponseError(c, http.StatusBadRequest, "Invalid Input: "+err.Error())
		return
	}

	user, err := services.RegisterServices(c, req.Name, req.Email, req.Password)
	if err != nil {
		utils.ResponseError(c, http.StatusBadRequest, err.Error())
		return
	}

	
	resp := models.RegisterResponse{
		ID:    user.ID,
		Name:  user.Name,
		Email: user.Email,
	}

	c.JSON(http.StatusCreated, gin.H{
		"success": true,
		"message": "User created successfully",
		"data":    resp,
	})
}

func LoginHandler(c *gin.Context) {
	var req models.LoginRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		utils.ResponseError(c, http.StatusBadRequest, "Invalid Input: "+err.Error())
		return
	}
	
	user, err := services.GetUserByEmail(req.Email)
	if err != nil {
		utils.ResponseError(c, http.StatusUnauthorized, "Email not found")
		return
	}

	
	token, err := services.LoginServices(req.Email, req.Password)
	if err != nil {
		utils.ResponseError(c, http.StatusUnauthorized, err.Error())
		return
	}

	
	resp := models.LoginResponse{
		User: models.RegisterResponse{
			ID:    user.ID,
			Name:  user.Name,
			Email: user.Email,
		},
		
	}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"message": "Login success",
		"data":    resp,
		"Token": token,
	})
}
