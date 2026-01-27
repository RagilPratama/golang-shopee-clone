package main

import (
	"golang-shopee/config"
	"golang-shopee/models"
	"log"
)

func main() {
	config.ConnectDatabase()

	config.DB.AutoMigrate(&models.Feature{})

	features := []models.Feature{
		{
			ID:       "cashback",
			Title:    "Cashback Saldo",
			Subtitle: "Rp200.003",
			Amount:   "Rp1.548",
			Icon:     "cash",
		},
		{
			ID:       "checkin",
			Title:    "Klaim 25RB!",
			Subtitle: "Cek-in!",
			Amount:   "",
			Icon:     "ticket-percent",
		},
		{
			ID:       "transfer",
			Title:    "Kirim Uang",
			Subtitle: "Gratis Admin",
			Amount:   "",
			Icon:     "bank-transfer",
		},
	}

	for _, feature := range features {
		if err := config.DB.Save(&feature).Error; err != nil {
			log.Printf("Failed to seed feature %s: %v", feature.ID, err)
		} else {
			log.Printf("Seeded feature: %s", feature.Title)
		}
	}
}

