package repository

import (
	"go_test/gin/global"
	"go_test/gin/model"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func main() {
	gorm.Open(sqlite.Open("identifier.sqlite"), &gorm.Config{})
}

type UserRepository struct {
}

func NewUserRepository() *UserRepository {
	return &UserRepository{}
}
func (userService *UserRepository) ListUsers() []model.User {
	users := []model.User{}
	global.DB.Debug().Limit(10).Find(&users)
	return users
}

func (userService *UserRepository) CreateUser(user model.User) model.User {
	tx := global.DB.Debug().Create(&user)
	if tx.Error != nil {
	}

	return user
}

func (userService *UserRepository) GetUser(id uint) model.User {
	var user model.User = model.User{}
	global.DB.Debug().Where(id).First(&user)
	return user
}

func (userService *UserRepository) DeleteUser(id uint) {
	global.DB.Debug().Where(id).Delete(&model.User{})
}

func (userService *UserRepository) UpdateUser(user model.User) {
	global.DB.Debug().Model(&model.User{}).Where("id = ?", user.ID).Updates(user)
}
