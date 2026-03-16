package controllers

import (
	"net/http"
	"sample-api/models"
	"sample-api/utils"

	"github.com/gin-gonic/gin"
)

type ProductInput struct {
	Name        string  `json:"name" binding:"required"`
	Description string  `json:"description"`
	Price       float64 `json:"price" binding:"required"`
	Category    string  `json:"category" binding:"required"`
	Stock       int     `json:"stock" binding:"required"`
}

// GetUserProfile godoc
// @Summary Get current user profile
// @Description Get profile of the authenticated user
// @Tags protected
// @Security BearerAuth
// @Produce json
// @Success 200 {object} utils.Response
// @Router /user/profile [get]
func GetUserProfile(c *gin.Context) {
	userID, _ := c.Get("user_id")
	var user models.User
	if err := models.DB.First(&user, userID).Error; err != nil {
		utils.Error(c, http.StatusNotFound, "User not found")
		return
	}
	utils.Success(c, user, "User profile retrieved")
}

// CreateProduct godoc
// @Summary Create a new product
// @Description Add a new product to the catalog
// @Tags protected
// @Security BearerAuth
// @Accept json
// @Produce json
// @Param input body ProductInput true "Product data"
// @Success 201 {object} utils.Response
// @Router /products [post]
func CreateProduct(c *gin.Context) {
	var input ProductInput
	if err := c.ShouldBindJSON(&input); err != nil {
		utils.Error(c, http.StatusBadRequest, err.Error())
		return
	}

	product := models.Product{
		Name:        input.Name,
		Description: input.Description,
		Price:       input.Price,
		Category:    input.Category,
		Stock:       input.Stock,
	}

	if err := models.DB.Create(&product).Error; err != nil {
		utils.Error(c, http.StatusInternalServerError, "Failed to create product")
		return
	}

	utils.Success(c, product, "Product created successfully")
}

// UpdateProduct godoc
// @Summary Update an existing product
// @Description Update product details by ID
// @Tags protected
// @Security BearerAuth
// @Accept json
// @Produce json
// @Param id path int true "Product ID"
// @Param input body ProductInput true "Updated product data"
// @Success 200 {object} utils.Response
// @Router /products/{id} [put]
func UpdateProduct(c *gin.Context) {
	id := c.Param("id")
	var product models.Product
	if err := models.DB.First(&product, id).Error; err != nil {
		utils.Error(c, http.StatusNotFound, "Product not found")
		return
	}

	var input ProductInput
	if err := c.ShouldBindJSON(&input); err != nil {
		utils.Error(c, http.StatusBadRequest, err.Error())
		return
	}

	models.DB.Model(&product).Updates(models.Product{
		Name:        input.Name,
		Description: input.Description,
		Price:       input.Price,
		Category:    input.Category,
		Stock:       input.Stock,
	})

	utils.Success(c, product, "Product updated successfully")
}

// DeleteProduct godoc
// @Summary Delete a product
// @Description Remove a product from the catalog by ID
// @Tags protected
// @Security BearerAuth
// @Produce json
// @Param id path int true "Product ID"
// @Success 200 {object} utils.Response
// @Router /products/{id} [delete]
func DeleteProduct(c *gin.Context) {
	id := c.Param("id")
	var product models.Product
	if err := models.DB.First(&product, id).Error; err != nil {
		utils.Error(c, http.StatusNotFound, "Product not found")
		return
	}

	models.DB.Delete(&product)
	utils.Success(c, nil, "Product deleted successfully")
}

// GetUserStats godoc
// @Summary Get authenticated user stats
// @Description Get some statistics for the logged-in user
// @Tags protected
// @Security BearerAuth
// @Produce json
// @Success 200 {object} utils.Response
// @Router /user/stats [get]
func GetUserStats(c *gin.Context) {
	username, _ := c.Get("username")
	stats := gin.H{
		"welcome_message": "Hello " + username.(string),
		"active_session":  true,
		"last_login":      "Just now",
	}
	utils.Success(c, stats, "User stats retrieved")
}
