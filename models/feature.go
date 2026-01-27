package models

type Feature struct {
	ID       string `json:"id" gorm:"primaryKey"`
	Title    string `json:"title"`
	Subtitle string `json:"subtitle"`
	Amount   string `json:"amount"`
	Icon     string `json:"icon"`
}

