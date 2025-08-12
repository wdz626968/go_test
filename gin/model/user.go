package model

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Username string `gorm:"unique;not null;default:hhh"`
	Password string `gorm:"not null"`
	Email    string `gorm:"unique;not null"`
	Name     string `gorm:"not null"`
	Age      uint
	Role     string `gorm:"not null;default:'user'"`
}

type Post struct {
	gorm.Model
	Title   string `gorm:"not null"`
	Content string `gorm:"not null"`
	UserID  uint
	User    User
}

type Comment struct {
	gorm.Model
	Content string `gorm:"not null"`
	UserID  uint
	User    User
	PostID  uint
	Post    Post
}
