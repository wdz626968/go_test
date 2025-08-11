package controller

import (
	"go_test/gin/controller/dto"
	"go_test/gin/model"
	"go_test/gin/repository"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

var userRepository = repository.NewUserRepository()

func GetUser(context *gin.Context) {
	id := context.Param("id")
	atoi, _ := strconv.Atoi(id)

	user := userRepository.GetUser(uint(atoi))

	context.JSON(http.StatusCreated, user)

}

func CreateUser(context *gin.Context) {
	var req dto.User
	if err := context.ShouldBindJSON(&req); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	userInfo := model.User{
		Name:  req.Name,
		Email: req.Email,
		Age:   req.Age,
	}
	user := userRepository.CreateUser(userInfo)

	context.JSON(http.StatusCreated, user)

}
