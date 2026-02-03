package main

import (
	"fmt"
	"golang-shopee/config"
	"log"
	"math/rand"
	"time"

	"github.com/google/uuid"
	"github.com/lib/pq"
)

type Product struct {
	ID          string         `json:"id" gorm:"primaryKey"`
	Title       string         `json:"title"`
	Price       string         `json:"price"`
	Rating      float64        `json:"rating"`
	RatingCount int            `json:"ratingCount"`
	Sold        int            `json:"sold"`
	Image       string         `json:"image"`
	ImageUrl    pq.StringArray `json:"imgUrl" gorm:"type:text[]"`

	IsMall     bool   `json:"isMall"`
	IsOri      bool   `json:"isOri"`
	IsTrending bool   `json:"isTrending"`
	IsFavorite bool   `json:"isFavorite"`
	Coin       string `json:"coin"`
	Status     string `json:"status"`
	Diskon     string `json:"diskon"`
	Category   string `json:"category"`

	Description string `json:"description"`
	Shipping    string `json:"shipping"`
	Promo       string `json:"promo"`
	CreatedAt   string `json:"createdAt"`
	Kota        string `json:"kota"`
	Durasi      string `json:"durasi"`

	ShopID uint `json:"ShopID"`
}

func main() {
	config.ConnectDatabase()

	// Clear existing products
	fmt.Println("Clearing existing products...")
	if err := config.DB.Exec("DELETE FROM products").Error; err != nil {
		log.Fatalf("Failed to clear products: %v", err)
	}

	// Seed data
	fmt.Println("Seeding products...")
	rand.Seed(time.Now().UnixNano())

	categories := []string{"Handphone & Aksesoris", "Komputer & Laptop", "Pakaian Pria", "Pakaian Wanita", "Elektronik"}
	statuses := []string{"active", "inactive"}
	shippings := []string{"Reguler (Cashless)", "Hemat", "Kargo"}
	promos := []string{"Cashback 100rb", "Gratis Ongkir", "Diskon 50%"}
	kotas := []string{"Jakarta Pusat", "Jakarta Barat", "Bandung", "Surabaya", "Medan"}
	durasis := []string{"Estimasi 2-3 Hari", "Estimasi 1-2 Hari", "Estimasi 3-5 Hari"}

	imageUrls := []string{
		"https://images.unsplash.com/photo-1610945415295-d9bbf067e59c?auto=format&fit=crop&w=500&q=60",
		"https://images.unsplash.com/photo-1678911820864-e2c567c655d2?auto=format&fit=crop&w=500&q=60",
		"https://images.unsplash.com/photo-1598327770170-5ca08b33c5d7?auto=format&fit=crop&w=500&q=60",
		"https://images.unsplash.com/photo-1517336714731-489689fd1ca8?auto=format&fit=crop&w=500&q=60",
		"https://images.unsplash.com/photo-1523275335684-37898b6baf30?auto=format&fit=crop&w=500&q=60",
	}

	for i := 0; i < 50; i++ {
		product := Product{
			ID:          uuid.New().String(),
			Title:       fmt.Sprintf("Product Dummy %d", i+1),
			Price:       fmt.Sprintf("%d000", rand.Intn(100)+50),
			Rating:      float64(rand.Intn(20)+30) / 10.0, // 3.0 - 5.0
			RatingCount: rand.Intn(2000),
			Sold:        rand.Intn(10000),
			Image:       imageUrls[rand.Intn(len(imageUrls))],
			ImageUrl: pq.StringArray{
				imageUrls[rand.Intn(len(imageUrls))],
				imageUrls[rand.Intn(len(imageUrls))],
				imageUrls[rand.Intn(len(imageUrls))],
			},
			IsMall:      rand.Intn(2) == 1,
			IsOri:       rand.Intn(2) == 1,
			IsTrending:  rand.Intn(2) == 1,
			IsFavorite:  rand.Intn(2) == 1,
			Coin:        fmt.Sprintf("%d", rand.Intn(100)*10),
			Status:      statuses[rand.Intn(len(statuses))],
			Diskon:      fmt.Sprintf("%d%%", rand.Intn(50)+5),
			Category:    categories[rand.Intn(len(categories))],
			Description: fmt.Sprintf("Deskripsi untuk Product Dummy %d. Barang berkualitas tinggi dan original.", i+1),
			Shipping:    shippings[rand.Intn(len(shippings))],
			Promo:       promos[rand.Intn(len(promos))],
			CreatedAt:   time.Now().Format(time.RFC3339),
			Kota:        kotas[rand.Intn(len(kotas))],
			Durasi:      durasis[rand.Intn(len(durasis))],
			ShopID:      uint(rand.Intn(5) + 1), // 1-5
		}

		if err := config.DB.Create(&product).Error; err != nil {
			log.Printf("Failed to create product %d: %v", i+1, err)
		}
	}

	fmt.Println("Seeding completed successfully!")
}
