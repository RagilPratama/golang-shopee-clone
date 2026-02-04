package controllers

import (
	"net/http"

	"golang-shopee/config"
	"golang-shopee/models"

	"github.com/gin-gonic/gin"
)

// GET /products
// Get all products
// @Summary      Show all products
// @Description  Get all products
// @Tags         products
// @Accept       json
// @Produce      json
// @Success      200  {object}  models.ProductsResponse
// @Router       /products [get]
func FindProducts(c *gin.Context) {
	var products []models.Product
	// Optimasi: Gunakan Joins daripada Preload untuk mengurangi Round Trip Time (RTT) ke database
	// Joins akan melakukan 1 query (LEFT JOIN) sedangkan Preload melakukan 2 query
	tx := config.DB.Joins("Shop").Find(&products)
	if tx.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": tx.Error.Error(),
			"data":  []models.Product{},
		})
		return
	}
	if products == nil {
		products = []models.Product{}
	}

	c.JSON(http.StatusOK, gin.H{"data": products})
}

// POST /products
// Create new product
// @Summary      Create a product
// @Description  Create a new product
// @Tags         products
// @Accept       json
// @Produce      json
// @Param        product  body      models.Product  true  "Product JSON"
// @Success      200      {object}  models.ProductResponse
// @Failure      400      {object}  map[string]interface{}
// @Router       /products [post]
func CreateProduct(c *gin.Context) {
	var input models.Product
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	product := input

	config.DB.Create(&product)

	c.JSON(http.StatusOK, gin.H{"data": product})
}

// GET /products/:id
// Get single product
// @Summary      Show a product
// @Description  Get product by ID
// @Tags         products
// @Accept       json
// @Produce      json
// @Param        id   path      string  true  "Product ID"
// @Success      200  {object}  models.ProductResponse
// @Failure      400  {object}  map[string]interface{}
// @Router       /products/{id} [get]
func FindProduct(c *gin.Context) {
	var product models.Product

	if err := config.DB.Preload("Shop").Where("id = ?", c.Param("id")).First(&product).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Product not found!"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": product})
}

// PUT /products/:id
// Update a product
// @Summary      Update a product
// @Description  Update product by ID
// @Tags         products
// @Accept       json
// @Produce      json
// @Param        id       path      string          true  "Product ID"
// @Param        product  body      models.Product  true  "Product JSON"
// @Success      200      {object}  map[string]interface{}
// @Failure      400      {object}  map[string]interface{}
// @Router       /products/{id} [put]
func UpdateProduct(c *gin.Context) {
	var product models.Product
	if err := config.DB.Where("id = ?", c.Param("id")).First(&product).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Product not found!"})
		return
	}

	var input models.Product
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	config.DB.Model(&product).Updates(input)

	c.JSON(http.StatusOK, gin.H{"data": product})
}

// DELETE /products/:id
// Delete a product
// @Summary      Delete a product
// @Description  Delete product by ID
// @Tags         products
// @Accept       json
// @Produce      json
// @Param        id   path      string  true  "Product ID"
// @Success      200  {object}  map[string]interface{}
// @Failure      400  {object}  map[string]interface{}
// @Router       /products/{id} [delete]
func DeleteProduct(c *gin.Context) {
	var product models.Product
	if err := config.DB.Where("id = ?", c.Param("id")).First(&product).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Product not found!"})
		return
	}

	config.DB.Delete(&product)

	c.JSON(http.StatusOK, gin.H{"data": true})
}
