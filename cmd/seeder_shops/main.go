package main

import (
	"fmt"
	"golang-shopee/config"
	"golang-shopee/models"

	"gorm.io/gorm/clause"
)

func main() {
	config.ConnectDatabase()
	config.DB.AutoMigrate(&models.Shop{})

	shops := []models.Shop{
		{ID: 1, Name: "Apple Official Store", Rating: 4.9, ProductCount: 125, ChatPercentage: 98, Location: "Jakarta Pusat"},
		{ID: 2, Name: "Apple Authorized Store", Rating: 5.0, ProductCount: 87, ChatPercentage: 95, Location: "Bandung"},
		{ID: 3, Name: "Innisfree Official", Rating: 4.8, ProductCount: 210, ChatPercentage: 99, Location: "Surabaya"},
		{ID: 5, Name: "Apple Official Store", Rating: 4.9, ProductCount: 125, ChatPercentage: 98, Location: "Jakarta Pusat"},
	}

	for _, s := range shops {
		err := config.DB.Clauses(clause.OnConflict{
			Columns:   []clause.Column{{Name: "id"}},
			DoUpdates: clause.AssignmentColumns([]string{"name", "rating", "product_count", "chat_percentage", "location"}),
		}).Create(&s).Error
		if err != nil {
			fmt.Println("failed upsert shop:", s.ID, err)
		}
	}

	fmt.Println("Seeded shops successfully")
}
