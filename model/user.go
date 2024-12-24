package model

type User struct {
	ID   uint   `json:"id" gorm:"primaryKey"`
	Name string `json:"name"`
	Post int    `json:"post" gorm:"not null;default:0"`
}
