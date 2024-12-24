package model

// Entities (Blog Post)
type Post struct {
	ID      int    `json:"id" gorm:"primaryKey"`
	Title   string `json:"title"`
	Content string `json:"content"`
	Comment int    `json:"comment" gorm:"not null;default:0"`
	Author  int    `json:"author"`
}
