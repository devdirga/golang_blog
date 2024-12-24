package model

// Entities (Blog Post)
type Post struct {
	ID      uint   `json:"id" gorm:"primaryKey"`
	Title   string `json:"title"`
	Content string `json:"content"`
	Comment uint   `json:"comment" gorm:"not null;default:0"`
}
