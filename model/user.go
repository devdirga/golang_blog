package model

type User struct {
	ID       int    `json:"id" gorm:"primaryKey"`
	Name     string `json:"name"`
	Username string `json:"username"`
	Password string `json:"password"`
	Post     int    `json:"post" gorm:"not null;default:0"`
	Token    string `json:"token"`
}
