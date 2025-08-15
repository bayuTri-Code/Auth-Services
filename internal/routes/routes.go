package routes

import (
	"log"

	"github.com/bayuTri-Code/Auth-Services/internal/handler"
	"github.com/gin-gonic/gin"
)

func Routes() *gin.Engine {
	gin.SetMode(gin.ReleaseMode)

	r := gin.New()

	r.Use(gin.Recovery())

	err := r.SetTrustedProxies([]string{"127.0.0.1"})
	if err != nil {
		log.Panicf("Failed to set trusted proxies: %v", err)
	}

	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "success api test",
		})
	})

	service := r.Group("/auth")
	{
		service.POST("/register", handler.RegisterHandler)
		service.POST("/login", handler.LoginHandler)
	}
	
	return r
}
