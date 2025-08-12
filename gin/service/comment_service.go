package service

import (
	"go_test/gin/global"
	"go_test/gin/model"
)

type CommentService struct {
}

func NewCommentService() *CommentService {
	return &CommentService{}
}

func (commentService *CommentService) CreateComments(comment model.Comment) (model.Comment, error) {
	global.DB.Create(&comment)
	return comment, nil
}

func (commentService *CommentService) GetCommentsByPostId(postId uint) []model.Comment {
	commentsRes := []model.Comment{}

	global.DB.Debug().Preload("User").Preload("Post").Where("post_id = ?", postId).Find(&commentsRes)
	return commentsRes
}

func (commentService *CommentService) GetCommentById(id uint) []model.Comment {
	commentsRes := []model.Comment{}

	global.DB.Debug().Preload("User").Preload("Post").Where("id = ?", id).First(&commentsRes)
	return commentsRes
}
