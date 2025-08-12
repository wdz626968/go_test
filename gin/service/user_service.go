package service

import (
	"go_test/gin/global"
	"go_test/gin/model"
)

type UserService struct {
}

func NewUserService() *UserService {
	return &UserService{}
}
func (userService *UserService) ListUsers() []model.User {
	users := []model.User{}
	global.DB.Debug().Limit(10).Find(&users)
	return users
}

func (userService *UserService) CreateUser(user model.User) model.User {
	tx := global.DB.Debug().Create(&user)
	if tx.Error != nil {
	}

	return user
}

func (userService *UserService) GetUser(id uint) model.User {
	var user model.User = model.User{}
	global.DB.Debug().Where(id).First(&user)
	return user
}

func (userService *UserService) DeleteUser(id uint) {
	global.DB.Debug().Where(id).Delete(&model.User{})
}

func (userService *UserService) UpdateUser(user model.User) {
	global.DB.Debug().Model(&model.User{}).Where("id = ?", user.ID).Updates(user)
}
