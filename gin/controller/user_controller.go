package controller

import (
	"go_test/gin/controller/dto"
	"go_test/gin/global"
	"go_test/gin/model"
	"go_test/gin/service"
	"go_test/gin/utils"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

var userService = service.NewUserService()

func GetUser(context *gin.Context) {
	id := context.Param("id")
	atoi, _ := strconv.Atoi(id)

	user := userService.GetUser(uint(atoi))

	context.JSON(http.StatusCreated, user)

}

func CreateUser(context *gin.Context) {
	var req dto.User
	if err := context.ShouldBindJSON(&req); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	userInfo := model.User{
		Name:     req.Name,
		Email:    req.Email,
		Age:      req.Age,
		Username: req.Username,
		Password: req.Password,
	}
	user := userService.CreateUser(userInfo)

	context.JSON(http.StatusCreated, user)

}

func Register(c *gin.Context) {
	var user dto.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	// 加密密码
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to hash password"})
		return
	}
	user.Password = string(hashedPassword)

	if err := global.DB.Create(&user).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create user"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "User registered successfully"})
}

func Login(c *gin.Context) {
	var user dto.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var storedUser model.User
	if err := global.DB.Where("username = ?", user.Username).First(&storedUser).Error; err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid username or password"})
		return
	}

	// 验证密码
	if err := bcrypt.CompareHashAndPassword([]byte(storedUser.Password), []byte(user.Password)); err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid username or password"})
		return
	}

	// 生成 JWT
	token, err := utils.GenerateJWT(storedUser.Username, storedUser.Role, storedUser.ID)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate token"})
		return
	}
	c.JSON(http.StatusOK, token)
	// 剩下的逻辑...
}
