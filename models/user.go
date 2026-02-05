package models

import "time"

type User struct {
	ID        uint      `json:"id" gorm:"primaryKey"`
	Name      string    `json:"name"`
	Email     string    `json:"email" gorm:"unique"`
	AvatarURL string    `json:"avatar_url"`
	GithubID  string    `json:"github_id" gorm:"unique"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
