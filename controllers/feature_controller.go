package controllers

import (
	"golang-shopee/config"
	"golang-shopee/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

// GET /features
// Get all features
// @Summary      Show all features
// @Description  Get all features
// @Tags         features
// @Accept       json
// @Produce      json
// @Success      200  {object}  map[string]interface{}
// @Router       /features [get]
func FindFeatures(c *gin.Context) {
	var features []models.Feature
	config.DB.Find(&features)

	c.JSON(http.StatusOK, gin.H{"data": features})
}

// POST /features
// Create new feature
// @Summary      Create a new feature
// @Description  Create a new feature
// @Tags         features
// @Accept       json
// @Produce      json
// @Param        feature  body      models.Feature  true  "Feature"
// @Success      200      {object}  models.Feature
// @Router       /features [post]
func CreateFeature(c *gin.Context) {
	var input models.Feature
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	feature := models.Feature{
		ID:       input.ID,
		Title:    input.Title,
		Subtitle: input.Subtitle,
		Amount:   input.Amount,
		Icon:     input.Icon,
	}

	config.DB.Create(&feature)

	c.JSON(http.StatusOK, gin.H{"data": feature})
}
