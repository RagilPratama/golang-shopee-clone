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
	// Optimasi: Gunakan Native Query untuk performa maksimal dan kontrol penuh
	query := `
		SELECT 
			p.id, p.title, p.price, p.rating, p.rating_count, p.sold, p.image, p.image_url,
			p.is_mall, p.is_ori, p.is_trending, p.is_favorite, p.coin, p.status, p.diskon, p.category,
			p.description, p.shipping, p.promo, p.created_at, p.kota, p.durasi, p.shop_id,
			s.id, s.name, s.rating, s.product_count, s.chat_percentage, s.location
		FROM products p
		LEFT JOIN shops s ON p.shop_id = s.id
	`

	rows, err := config.DB.Raw(query).Rows()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
			"data":  []models.Product{},
		})
		return
	}
	defer rows.Close()

	for rows.Next() {
		var p models.Product
		var s models.Shop

		// Scan setiap kolom secara manual ke dalam struct
		// Pastikan urutan scan sesuai dengan urutan SELECT
		if err := rows.Scan(
			&p.ID, &p.Title, &p.Price, &p.Rating, &p.RatingCount, &p.Sold, &p.Image, &p.ImageUrl,
			&p.IsMall, &p.IsOri, &p.IsTrending, &p.IsFavorite, &p.Coin, &p.Status, &p.Diskon, &p.Category,
			&p.Description, &p.Shipping, &p.Promo, &p.CreatedAt, &p.Kota, &p.Durasi, &p.ShopID,
			&s.ID, &s.Name, &s.Rating, &s.ProductCount, &s.ChatPercentage, &s.Location,
		); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to scan product: " + err.Error()})
			return
		}

		p.Shop = s
		products = append(products, p)
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
