package main

import (
	"net/http"

	"github.com/bayuTri-Code/Auth-Services/database"
	"github.com/bayuTri-Code/Auth-Services/internal/config"
	"github.com/gin-gonic/gin"
)

//testing
func homeTesting(c *gin.Context){
	c.JSON(http.StatusOK, gin.H{
		"message": "OK",
	})
}

func main(){
	config.ConfigDb()
	database.PostgresConn()

	//testing
	router := gin.Default()
	router.GET("/", homeTesting)
	if err := router.Run(); err != nil {
		return
	}
	
}