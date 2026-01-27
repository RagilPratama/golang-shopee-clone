package routes

import (
	"golang-shopee/controllers"
	_ "golang-shopee/docs"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()

	r.GET("/products", controllers.FindProducts)
	r.POST("/products", controllers.CreateProduct)
	r.GET("/products/:id", controllers.FindProduct)
	r.PUT("/products/:id", controllers.UpdateProduct)
	r.DELETE("/products/:id", controllers.DeleteProduct)

	r.GET("/menus", controllers.FindMenus)
	r.POST("/menus", controllers.CreateMenu)

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	return r
}
