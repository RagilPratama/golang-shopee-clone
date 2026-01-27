package models

type Menu struct {
	ID        string `json:"id" gorm:"primaryKey"`
	Title     string `json:"title"`
	Icon      string `json:"icon"`
	IconColor string `json:"iconColor"`
	BgColor   string `json:"bgColor"`
	Route     string `json:"route"`
}
