package service

import (
	"go_test/gin/global"
	"go_test/gin/model"
)

type PostService struct {
}

func NewPostService() *PostService {
	return &PostService{}
}

func (postService *PostService) GetPost(postId uint) model.Post {
	post := model.Post{}
	global.DB.Where("id = ?", postId).First(&post)
	return post
}

func (postService *PostService) GetAllPosts() []model.Post {
	posts := []model.Post{}
	global.DB.Preload("user").Find(&posts)
	return posts
}

func (postService *PostService) CreatePosts(post model.Post) model.Post {
	global.DB.Create(&post)
	return post
}

func (postService *PostService) UpdatePosts(post model.Post) (model.Post, error) {
	result := global.DB.Model(&post).Updates(post)
	if result.Error != nil {
		return model.Post{}, result.Error
	}
	return post, nil
}

func (postService *PostService) DeletePosts(postId uint) model.Post {
	post := model.Post{}
	global.DB.Where("id = ?", postId).Delete(&model.Post{}).First(&post)
	return post
}
