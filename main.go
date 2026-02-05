package main

import (
	"log"

	"golang-shopee/config"
	"golang-shopee/models"
	"golang-shopee/routes"

	"github.com/joho/godotenv"
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
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found, assuming production environment (Railway/Docker). Using system variables.")
	} else {
		log.Println(".env file loaded successfully")
	}

	config.ConnectDatabase()
	config.DB.AutoMigrate(&models.Shop{}, &models.Product{}, &models.Menu{}, &models.Feature{}, &models.User{})

	r := routes.SetupRouter()
	r.Run(":8080")
}
