package model

import "time"

type Post struct {
	ID        int       `json:"id" gorm:"primaryKey"`
	Title     string    `json:"title"`
	Content   string    `json:"content"`
	Comment   int       `json:"comment" gorm:"not null;default:0"`
	Author    int       `json:"author"`
	CreatedAt time.Time `json:"created_at" gorm:"not null, default:now()"`
	UpdatedAt time.Time `json:"Updated_at" gorm:"not null, default:now()"`
}
