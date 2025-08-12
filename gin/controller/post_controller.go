package controller

import (
	"go_test/gin/controller/dto"
	"go_test/gin/model"
	"go_test/gin/service"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

var postService = service.NewPostService()

func GetPost(context *gin.Context) {
	id := context.Param("id")
	atoi, _ := strconv.Atoi(id)

	post := postService.GetPost(uint(atoi))

	context.JSON(http.StatusOK, post)
}

func GetAllPosts(context *gin.Context) {
	post := postService.GetAllPosts()

	context.JSON(http.StatusOK, post)
}

func CreatePosts(context *gin.Context) {
	var req dto.Post
	if err := context.ShouldBindJSON(&req); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	post := model.Post{
		Model:   gorm.Model{},
		Title:   req.Title,
		Content: req.Content,
		UserID:  req.UserID,
	}
	postService.CreatePosts(post)
	context.JSON(http.StatusCreated, post)
}

func UpdatePosts(context *gin.Context) {
	var req dto.Post
	if err := context.ShouldBindJSON(&req); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	post := model.Post{
		Model: gorm.Model{
			ID:        req.ID,
			CreatedAt: time.Time{},
			UpdatedAt: time.Time{},
			DeletedAt: gorm.DeletedAt{},
		},
		Title:   req.Title,
		Content: req.Content,
		UserID:  req.UserID,
	}
	posts, _ := postService.UpdatePosts(post)

	context.JSON(http.StatusCreated, posts)
}

func DeletePosts(context *gin.Context) {
	id := context.Param("id")
	atoi, _ := strconv.Atoi(id)

	post := postService.DeletePosts(uint(atoi))

	context.JSON(http.StatusOK, post)
}
