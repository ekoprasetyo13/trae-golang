package controllers

import (
	"net/http"
	"sample-api/models"
	"sample-api/utils"

	"github.com/gin-gonic/gin"
)

// GetProducts godoc
// @Summary Get all products
// @Description Get a list of all sample products
// @Tags public
// @Accept json
// @Produce json
// @Success 200 {object} utils.Response
// @Router /products [get]
func GetProducts(c *gin.Context) {
	var products []models.Product
	models.DB.Find(&products)
	utils.Success(c, products, "Products retrieved successfully")
}

// GetProductByID godoc
// @Summary Get product by ID
// @Description Get a single product by its ID
// @Tags public
// @Accept json
// @Produce json
// @Param id path int true "Product ID"
// @Success 200 {object} utils.Response
// @Failure 404 {object} utils.Response
// @Router /products/{id} [get]
func GetProductByID(c *gin.Context) {
	id := c.Param("id")
	var product models.Product
	if err := models.DB.First(&product, id).Error; err != nil {
		utils.Error(c, http.StatusNotFound, "Product not found")
		return
	}
	utils.Success(c, product, "Product retrieved successfully")
}

// GetProductsByCategory godoc
// @Summary Get products by category
// @Description Get a list of products in a specific category
// @Tags public
// @Accept json
// @Produce json
// @Param category path string true "Category name"
// @Success 200 {object} utils.Response
// @Router /products/category/{category} [get]
func GetProductsByCategory(c *gin.Context) {
	category := c.Param("category")
	var products []models.Product
	models.DB.Where("category = ?", category).Find(&products)
	utils.Success(c, products, "Products in category retrieved")
}

// SearchProducts godoc
// @Summary Search products
// @Description Search products by name
// @Tags public
// @Accept json
// @Produce json
// @Param q query string true "Search query"
// @Success 200 {object} utils.Response
// @Router /products/search [get]
func SearchProducts(c *gin.Context) {
	query := c.Query("q")
	var products []models.Product
	models.DB.Where("name LIKE ?", "%"+query+"%").Find(&products)
	utils.Success(c, products, "Search results retrieved")
}

// GetExpensiveProducts godoc
// @Summary Get expensive products
// @Description Get products with price greater than 500
// @Tags public
// @Accept json
// @Produce json
// @Success 200 {object} utils.Response
// @Router /products/expensive [get]
func GetExpensiveProducts(c *gin.Context) {
	var products []models.Product
	models.DB.Where("price > ?", 500).Find(&products)
	utils.Success(c, products, "Expensive products retrieved")
}

// GetCheapProducts godoc
// @Summary Get cheap products
// @Description Get products with price less than 100
// @Tags public
// @Accept json
// @Produce json
// @Success 200 {object} utils.Response
// @Router /products/cheap [get]
func GetCheapProducts(c *gin.Context) {
	var products []models.Product
	models.DB.Where("price < ?", 100).Find(&products)
	utils.Success(c, products, "Cheap products retrieved")
}

// GetOutOfStockProducts godoc
// @Summary Get out of stock products
// @Description Get products with zero stock
// @Tags public
// @Accept json
// @Produce json
// @Success 200 {object} utils.Response
// @Router /products/out-of-stock [get]
func GetOutOfStockProducts(c *gin.Context) {
	var products []models.Product
	models.DB.Where("stock = ?", 0).Find(&products)
	utils.Success(c, products, "Out of stock products retrieved")
}

// GetCategories godoc
// @Summary Get all categories
// @Description Get a list of unique product categories
// @Tags public
// @Accept json
// @Produce json
// @Success 200 {object} utils.Response
// @Router /products/categories [get]
func GetCategories(c *gin.Context) {
	var categories []string
	models.DB.Model(&models.Product{}).Distinct("category").Pluck("category", &categories)
	utils.Success(c, categories, "Categories retrieved successfully")
}

// GetProductStats godoc
// @Summary Get product statistics
// @Description Get basic statistics of products
// @Tags public
// @Accept json
// @Produce json
// @Success 200 {object} utils.Response
// @Router /products/stats [get]
func GetProductStats(c *gin.Context) {
	var count int64
	var avgPrice float64
	models.DB.Model(&models.Product{}).Count(&count)
	models.DB.Model(&models.Product{}).Select("AVG(price)").Scan(&avgPrice)

	stats := gin.H{
		"total_products": count,
		"average_price":  avgPrice,
	}
	utils.Success(c, stats, "Product statistics retrieved")
}

// GetTopRatedProducts godoc
// @Summary Get top rated products
// @Description Mock endpoint for top rated products
// @Tags public
// @Accept json
// @Produce json
// @Success 200 {object} utils.Response
// @Router /products/top-rated [get]
func GetTopRatedProducts(c *gin.Context) {
	// Mocking top rated products
	var products []models.Product
	models.DB.Limit(3).Order("price desc").Find(&products)
	utils.Success(c, products, "Top rated products retrieved (mocked)")
}
