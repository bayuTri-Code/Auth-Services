package handler

import (

	"github.com/gin-gonic/gin"
)

func RegisterHandler(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "Register ok",
	})
}

func LoginHandler(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "Login ok",
	})
}
