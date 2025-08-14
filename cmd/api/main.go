package main

import (
	"fmt"

	"github.com/bayuTri-Code/Auth-Services/database"
	"github.com/bayuTri-Code/Auth-Services/internal/config"
	"github.com/bayuTri-Code/Auth-Services/internal/routes"
)



func main(){
	config.ConfigDb()
	database.PostgresConn()

	//routes
	r := routes.Routes()
	
	port := ":8080";

	fmt.Printf("server is running in http://localhost:%s", port)
	r.Run(port)
}