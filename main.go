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

// @host            golang-shopee-clone-production.up.railway.app
// @schemes         https
// @BasePath        /

func main() {
	config.ConnectDatabase()
	config.DB.AutoMigrate(&models.Product{}, &models.Menu{}, &models.Feature{})

	r := routes.SetupRouter()
	r.Run(":8080")
}
