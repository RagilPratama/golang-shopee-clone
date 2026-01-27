package controllers

import (
	"golang-shopee/config"
	"golang-shopee/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

// GET /menus
// Get all menus
// @Summary      Show all menus
// @Description  Get all menus
// @Tags         menus
// @Accept       json
// @Produce      json
// @Success      200  {object}  map[string]interface{}
// @Router       /menus [get]
func FindMenus(c *gin.Context) {
	var menus []models.Menu
	config.DB.Find(&menus)

	c.JSON(http.StatusOK, gin.H{"data": menus})
}

// POST /menus
// Create new menu
// @Summary      Create a new menu
// @Description  Create a new menu
// @Tags         menus
// @Accept       json
// @Produce      json
// @Param        menu  body      models.Menu  true  "Menu"
// @Success      200   {object}  models.Menu
// @Router       /menus [post]
func CreateMenu(c *gin.Context) {
	var input models.Menu
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	menu := models.Menu{
		ID:        input.ID,
		Title:     input.Title,
		Icon:      input.Icon,
		IconColor: input.IconColor,
		BgColor:   input.BgColor,
		Route:     input.Route,
	}

	config.DB.Create(&menu)

	c.JSON(http.StatusOK, gin.H{"data": menu})
}
