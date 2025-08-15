package handler

import (
	"net/http"

	"github.com/bayuTri-Code/Auth-Services/database"
	"github.com/bayuTri-Code/Auth-Services/internal/models"
	"github.com/bayuTri-Code/Auth-Services/internal/services"
	"github.com/bayuTri-Code/Auth-Services/internal/utils"
	"github.com/gin-gonic/gin"
)

type UserRequest struct {
	Name     string `json:"name" binding:"required"`
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func RegisterHandler(c *gin.Context) {
	var req UserRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		utils.ResponseError(c, http.StatusBadRequest, "Invalid Your Input :"+err.Error())
		return
	}
	Register, err := services.RegisterServices(c, req.Name, req.Email, req.Password)
	if err != nil {
		utils.ResponseError(c, http.StatusBadRequest, err.Error())
		return
	}

	utils.ResponseSuccess(c, http.StatusCreated, gin.H{
		"message": "create user successfully!",
		"user_id": Register.ID,
		"name":    Register.Email,
		"email":   Register.Password,
	})
}

func LoginHandler(c *gin.Context) {
	var input models.AuthData
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var user models.AuthData
	if err := database.Db.Where("email = ?", input.Email).First(&user).Error; err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid email or password"})
		return
	}

	input.Password = user.Password

	token, err := utils.GenerateJWT(user.ID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate token"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Succes to get token",
		"token":   token})
}
