package main

import (
	"golang-shopee/config"
	"golang-shopee/models"
	"golang-shopee/routes"
)

// @title           Golang Shopee API
// @version         1.0
// @description     This is a sample server for a Shopee-like API.
// @termsOfService  http://swagger.io/terms/

// @contact.name    API Support
// @contact.url     http://www.swagger.io/support
// @contact.email   support@swagger.io

// @license.name    Apache 2.0
// @license.url     http://www.apache.org/licenses/LICENSE-2.0.html

// @host            localhost:8080
// @BasePath        /

func main() {
	// Connect to database
	config.ConnectDatabase()

	// Auto Migrate
	config.DB.AutoMigrate(&models.Product{})

	// Setup Router
	r := routes.SetupRouter()

	// Run server
	r.Run(":8080")
}
