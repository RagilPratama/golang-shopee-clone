package controllers

import (
	"golang-shopee/config"
	"golang-shopee/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func FindFeatures(c *gin.Context) {
	var features []models.Feature
	config.DB.Find(&features)

	c.JSON(http.StatusOK, gin.H{"data": features})
}

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

