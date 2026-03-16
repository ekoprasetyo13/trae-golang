package main

import (
	"sample-api/controllers"
	"sample-api/middleware"
	"sample-api/models"

	_ "sample-api/docs"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title Sample Go API with SQLite
// @version 1.0
// @description This is a sample API server for a Go/SQLite project.
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host localhost:8080
// @BasePath /api/v1

// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization

func main() {
	// Initialize Database
	models.InitDB()

	r := gin.Default()

	// Swagger documentation
	r.GET("/docs-api/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	v1 := r.Group("/api/v1")
	{
		// Auth routes
		auth := v1.Group("/auth")
		{
			auth.POST("/register", controllers.Register)
			auth.POST("/login", controllers.Login)
		}

		// Public Product routes (10 APIs)
		public := v1.Group("/products")
		{
			public.GET("", controllers.GetProducts)
			public.GET("/:id", controllers.GetProductByID)
			public.GET("/category/:category", controllers.GetProductsByCategory)
			public.GET("/search", controllers.SearchProducts)
			public.GET("/expensive", controllers.GetExpensiveProducts)
			public.GET("/cheap", controllers.GetCheapProducts)
			public.GET("/out-of-stock", controllers.GetOutOfStockProducts)
			public.GET("/categories", controllers.GetCategories)
			public.GET("/stats", controllers.GetProductStats)
			public.GET("/top-rated", controllers.GetTopRatedProducts)
		}

		// Error routes
		errors := v1.Group("/errors")
		{
			errors.GET("/400", controllers.Error400)
			errors.GET("/401", controllers.Error401)
			errors.GET("/403", controllers.Error403)
			errors.GET("/404", controllers.Error404)
			errors.GET("/405", controllers.Error405)
			errors.GET("/408", controllers.Error408)
			errors.GET("/500", controllers.Error500)
			errors.GET("/503", controllers.Error503)
			errors.GET("/504", controllers.Error504)
		}

		// Protected routes (5 APIs)
		protected := v1.Group("")
		protected.Use(middleware.AuthMiddleware())
		{
			protected.GET("/user/profile", controllers.GetUserProfile)
			protected.GET("/user/stats", controllers.GetUserStats)
			protected.POST("/products", controllers.CreateProduct)
			protected.PUT("/products/:id", controllers.UpdateProduct)
			protected.DELETE("/products/:id", controllers.DeleteProduct)
		}
	}

	r.Run(":8080")
}
