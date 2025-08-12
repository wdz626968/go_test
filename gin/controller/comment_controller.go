package controller

import (
	"go_test/gin/controller/dto"
	"go_test/gin/model"
	"go_test/gin/service"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

var commentService = service.NewCommentService()

func GetComment(context *gin.Context) {
	param := context.Param("id")
	postId, _ := strconv.Atoi(param)
	comment := commentService.GetCommentById(uint(postId))
	context.JSON(http.StatusOK, comment)
}

func GetCommentsNyPostId(context *gin.Context) {
	param := context.Param("post_id")
	postId, _ := strconv.Atoi(param)
	comments := commentService.GetCommentsByPostId(uint(postId))
	context.JSON(http.StatusOK, comments)
}

func CreateComments(context *gin.Context) {
	var req dto.Comment
	if err := context.ShouldBindJSON(&req); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	comment := model.Comment{
		Model:   gorm.Model{},
		Content: req.Content,
		UserID:  req.UserID,
		PostID:  req.PostID,
	}
	comments, _ := commentService.CreateComments(comment)
	context.JSON(http.StatusOK, comments)
}

func UpdateComments(context *gin.Context) {

}

func DeleteComments(context *gin.Context) {

}
