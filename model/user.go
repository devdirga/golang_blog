package model

import "time"

type User struct {
	ID        int       `json:"id" gorm:"primaryKey"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	Password  string    `json:"password"`
	PostCount int       `json:"post_count" gorm:"not null;default:0"`
	Token     string    `json:"token"`
	CreatedAt time.Time `json:"created_at" gorm:"not null, default:now()"`
	UpdatedAt time.Time `json:"Updated_at" gorm:"not null, default:now()"`
}
