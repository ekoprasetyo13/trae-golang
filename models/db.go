package models

import (
	"fmt"
	"log"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/plugin/opentelemetry/tracing"
)

var DB *gorm.DB

// InitDB initializes the SQLite database and performs migrations
func InitDB() {
	var err error
	DB, err = gorm.Open(sqlite.Open("sample.db"), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}

	if err := DB.Use(tracing.NewPlugin()); err != nil {
		log.Fatal("Failed to initialize database tracing:", err)
	}

	fmt.Println("Database connection established")

	// Migrate schemas
	err = DB.AutoMigrate(&User{}, &Product{})
	if err != nil {
		log.Fatal("Failed to migrate database:", err)
	}

	fmt.Println("Database migration completed")

	// Seed data if products are empty
	SeedData()
}

// SeedData inserts some sample data into the database
func SeedData() {
	var count int64
	DB.Model(&Product{}).Count(&count)
	if count == 0 {
		products := []Product{
			{Name: "Laptop", Description: "High performance laptop", Price: 1200.00, Category: "Electronics", Stock: 10},
			{Name: "Smartphone", Description: "Latest flagship phone", Price: 800.00, Category: "Electronics", Stock: 25},
			{Name: "Coffee Maker", Description: "Brews best coffee", Price: 50.00, Category: "Home", Stock: 50},
			{Name: "Headphones", Description: "Noise cancelling", Price: 150.00, Category: "Accessories", Stock: 30},
			{Name: "Running Shoes", Description: "Lightweight and comfortable", Price: 80.00, Category: "Sports", Stock: 15},
			{Name: "Desk Chair", Description: "Ergonomic chair", Price: 200.00, Category: "Furniture", Stock: 12},
			{Name: "Monitor", Description: "4K resolution", Price: 300.00, Category: "Electronics", Stock: 8},
			{Name: "Keyboard", Description: "Mechanical keyboard", Price: 100.00, Category: "Accessories", Stock: 20},
			{Name: "Backpack", Description: "Waterproof backpack", Price: 60.00, Category: "Fashion", Stock: 40},
			{Name: "Water Bottle", Description: "Stainless steel", Price: 20.00, Category: "Home", Stock: 100},
			{Name: "Gaming Console", Description: "Next-gen gaming", Price: 500.00, Category: "Electronics", Stock: 5},
			{Name: "Smart Watch", Description: "Fitness tracker", Price: 250.00, Category: "Electronics", Stock: 15},
			{Name: "Blender", Description: "Powerful motor", Price: 70.00, Category: "Home", Stock: 20},
			{Name: "Table Lamp", Description: "Modern design", Price: 40.00, Category: "Furniture", Stock: 35},
			{Name: "Yoga Mat", Description: "Non-slip surface", Price: 30.00, Category: "Sports", Stock: 50},
		}

		for _, p := range products {
			DB.Create(&p)
		}
		fmt.Println("Seed data inserted")
	}
}
