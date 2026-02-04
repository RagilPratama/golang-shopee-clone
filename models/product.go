package models

import "github.com/lib/pq"

type Shop struct {
	ID             uint    `json:"-" gorm:"primaryKey"`
	Name           string  `json:"name"`
	Rating         float64 `json:"rating"`
	ProductCount   int     `json:"productCount"`
	ChatPercentage int     `json:"chatPercentage"`
	Location       string  `json:"location"`
}

type Product struct {
	ID          string         `json:"id" gorm:"primaryKey"`
	Title       string         `json:"title"`
	Price       string         `json:"price"`
	Rating      float64        `json:"rating"`
	RatingCount int            `json:"ratingCount"`
	Sold        int            `json:"sold"`
	Image       string         `json:"image"`
	ImageUrl    pq.StringArray `json:"ImageUrl" gorm:"type:text[]" swaggertype:"array,string"`

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

	Shop   Shop `json:"shopDetail" gorm:"foreignKey:ShopID"`
	ShopID uint `json:"ShopID" gorm:"index"`
}
