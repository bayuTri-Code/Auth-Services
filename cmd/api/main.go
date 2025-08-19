package main

import (
	"fmt"

	
	"github.com/bayuTri-Code/Auth-Services/database"
	"github.com/bayuTri-Code/Auth-Services/internal/config"
	"github.com/bayuTri-Code/Auth-Services/internal/routes"
)

func main() {
	config.ConfigDb()
	database.PostgresConn()

	// routes
	r := routes.Routes()

	host := "0.0.0.0"
	port := "8080"

	fmt.Printf("server is running in http://%s:%s\n", host, port)
	r.Run(host + ":" + port)
}
