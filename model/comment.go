package model

import "time"

type Comment struct {
	ID        int       `json:"id" gorm:"primaryKey"`
	Content   string    `json:"content"`
	PostID    int       `json:"post_id" gorm:"not null"`
	UserID    int       `json:"user_id" gorm:"not null"`
	IsActive  bool      `json:"is_active" gorm:"not null"`
	CreatedAt time.Time `json:"created_at" gorm:"not null, default:now()"`
	UpdatedAt time.Time `json:"Updated_at" gorm:"not null, default:now()"`
}
