package dto

import "gorm.io/gorm"

type User struct {
	ID       uint   `json:"id"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	Age      uint   `json:"age"`
	Username string `json:"username"`
	Password string `json:"password"`
	// admin 或 user
	Role string `gorm:"not null;default:'user'" json:"role"`
	gorm.Model
}
