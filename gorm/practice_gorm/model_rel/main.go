package main

import (
	"fmt"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func main() {
	db, err := initDB()
	if err != nil {
		panic(err)
	}
	db.AutoMigrate(&User{})
	db.AutoMigrate(&Post{})
	db.AutoMigrate(&Comment{})

	//分步创建
	//user := User{Username: "dzhwu1", Email: "xxx1x"}
	//db.Create(&user)
	//
	//post := Post{Title: "hhh", Content: "113seqwdqd", UserID: user.ID}
	//db.Create(&post)
	//
	//comment := Comment{Content: "测试", PostID: post.ID, UserID: user.ID}
	//db.Create(&comment)

	//直接关联创建
	//db.Create(&User{
	//	Model:    gorm.Model{},
	//	Username: "dzhwu",
	//	Email:    "xxxx",
	//	Posts: []Post{
	//		{
	//			Model:   gorm.Model{},
	//			Title:   "hhh",
	//			Content: "113seqwdqd",
	//			Comments: []Comment{
	//				{
	//					Model:   gorm.Model{},
	//					Content: "1111",
	//				},
	//			},
	//		},
	//	},
	//})
	user := User{}

	db.Preload("Posts").Preload("Posts.Comments").Where("id = ?", 2).First(&user)
	fmt.Println(user)
	db.Preload("Comments").Where("user_id", 2).First(&Post{})
	post := Post{}
	db.Preload("Comments").Select("posts.*, count(*) AS comment_count").Joins("left join  comments on comments.post_id = posts.id ").Group("posts.id").Order("comment_count desc").First(&post)
	fmt.Println(post)
}

type User struct {
	gorm.Model
	Username string `gorm:"unique;not null"`
	Email    string `gorm:"unique;not null"`
	Posts    []Post
}

type Post struct {
	gorm.Model
	Title    string `gorm:"not null"`
	Content  string `gorm:"type:text"`
	UserID   uint   `gorm:"index"`
	Comments []Comment
}

type Comment struct {
	gorm.Model
	Content string `gorm:"type:text;not null"`
	PostID  uint   `gorm:"index"`
	UserID  uint   `gorm:"index"`
}

func initDB() (*gorm.DB, error) {
	db, err := gorm.Open(sqlite.Open("identifier.sqlite"), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	return db, err
}
