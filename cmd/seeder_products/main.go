package main

import (
	"fmt"
	"golang-shopee/config"
	"golang-shopee/models"
)

func main() {
	config.ConnectDatabase()
	config.DB.AutoMigrate(&models.Shop{}, &models.Product{})

	products := []models.Product{
		{
			ID:          "p1",
			Title:       "iPhone 14 Pro Max 256GB",
			Price:       "Rp19.999.000",
			Rating:      5,
			RatingCount: 1820,
			Sold:        1245,
			Image:       "https://images.unsplash.com/photo-1511707171634-5f897ff02aa9?auto=format&fit=crop&w=1200&q=80",
			IsMall:      true,
			IsOri:       true,
			IsFavorite:  false,
			Coin:        "Rp200.000",
			Status:      "Tersedia",
			Diskon:      "Y",
			Category:    "Elektronik",
			Description: "iPhone 14 Pro Max 256GB dengan kamera Pro 48MP, layar Super Retina XDR 6.7 inci, dan performa tinggi berkat chip A16 Bionic.",
			Shipping:    "Gratis Ongkir",
			Promo:       "Cashback & Cicilan 0%",
			CreatedAt:   "2025-12-01T00:00:00+07:00",
			Kota:        "KAB BEKASI",
			Durasi:      "2-3 Hari",
			ShopID:      5,
		},
		{
			ID:          "p2",
			Title:       "iPhone 14 128GB",
			Price:       "Rp13.999.000",
			Rating:      4.9,
			RatingCount: 950,
			Sold:        820,
			Image:       "https://images.unsplash.com/photo-1512499617640-c2f999018b72?auto=format&fit=crop&w=1200&q=80",
			IsMall:      true,
			IsOri:       true,
			IsFavorite:  true,
			Coin:        "Rp150.000",
			Status:      "Tersedia",
			Diskon:      "Y",
			Category:    "Elektronik",
			Description: "iPhone 14 128GB dengan performa tinggi dan kamera ganda yang tajam.",
			Shipping:    "Gratis Ongkir",
			Promo:       "Voucher toko & cashback",
			CreatedAt:   "2025-11-20T00:00:00+07:00",
			Kota:        "JAKARTA PUSAT",
			Durasi:      "1-2 Hari",
			ShopID:      1,
		},
		{
			ID:          "p3",
			Title:       "MacBook Air M2 13\" 256GB",
			Price:       "Rp18.999.000",
			Rating:      4.9,
			RatingCount: 540,
			Sold:        430,
			Image:       "https://images.unsplash.com/photo-1517336714731-489689fd1ca8?auto=format&fit=crop&w=1200&q=80",
			IsMall:      true,
			IsOri:       true,
			IsFavorite:  false,
			Coin:        "Rp250.000",
			Status:      "Tersedia",
			Diskon:      "N",
			Category:    "Elektronik",
			Description: "MacBook Air M2 13 inci dengan desain tipis, baterai tahan lama, dan performa kencang.",
			Shipping:    "Gratis Ongkir",
			Promo:       "Cicilan 0% hingga 12 bulan",
			CreatedAt:   "2025-10-15T00:00:00+07:00",
			Kota:        "BANDUNG",
			Durasi:      "2-4 Hari",
			ShopID:      2,
		},
		{
			ID:          "p4",
			Title:       "Innisfree Green Tea Seed Serum 80ml",
			Price:       "Rp350.000",
			Rating:      4.9,
			RatingCount: 3210,
			Sold:        2800,
			Image:       "https://images.unsplash.com/photo-1612815154858-60aa4c59eaa2?auto=format&fit=crop&w=1200&q=80",
			IsMall:      true,
			IsOri:       true,
			IsFavorite:  true,
			Coin:        "Rp10.000",
			Status:      "Tersedia",
			Diskon:      "Y",
			Category:    "Kecantikan",
			Description: "Serum pelembap dengan ekstrak green tea untuk kulit lembap dan segar.",
			Shipping:    "Gratis Ongkir Xtra",
			Promo:       "Bundling hemat & cashback",
			CreatedAt:   "2025-09-10T00:00:00+07:00",
			Kota:        "SURABAYA",
			Durasi:      "2-3 Hari",
			ShopID:      3,
		},
		{
			ID:          "p5",
			Title:       "AirPods Pro 2nd Gen",
			Price:       "Rp3.999.000",
			Rating:      4.8,
			RatingCount: 1650,
			Sold:        1400,
			Image:       "https://images.unsplash.com/photo-1588423771073-b8903fbb85b5?auto=format&fit=crop&w=1200&q=80",
			IsMall:      true,
			IsOri:       true,
			IsFavorite:  false,
			Coin:        "Rp50.000",
			Status:      "Tersedia",
			Diskon:      "Y",
			Category:    "Elektronik",
			Description: "AirPods Pro generasi kedua dengan Active Noise Cancellation dan Adaptive Transparency.",
			Shipping:    "Gratis Ongkir",
			Promo:       "Cashback & voucher toko",
			CreatedAt:   "2025-12-05T00:00:00+07:00",
			Kota:        "JAKARTA PUSAT",
			Durasi:      "1-2 Hari",
			ShopID:      1,
		},
	}

	for _, p := range products {
		if err := config.DB.Save(&p).Error; err != nil {
			fmt.Println("failed save product:", p.ID, err)
		}
	}

	fmt.Println("Seeded products successfully")
}

