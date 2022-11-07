package main

import (
	"go-api-pet/database"
	_ "go-api-pet/docs"
	"go-api-pet/routes"
	"os"
)

// @title           Pet API
// @version         1.0
// @description    	This is to manage registered pets in the application

// @contact.name   Janaina Pedrina

// @host      localhost:3000
// @BasePath  /
func main() {
	os.Setenv("env", "DEV")
	database.Connection()
	routes.HandleRequests()
}
