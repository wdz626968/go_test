package main

import (
	"errors"
	"fmt"
	"go_test/gorm/constant"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

// 更新
func main() {
	db := InitDB()
	user := &User{}
	db.First(&user)

	user.Name = "jinzhu 2"
	user.Age = 100
	db.Debug().Save(&user)

	//没有主键id不会update
	db.Save(&User{Name: "jinzhu", Age: 100})
	// INSERT INTO `users` (`name`,`age`,`birthday`,`update_at`) VALUES ("jinzhu",100,"0000-00-00 00:00:00","0000-00-00 00:00:00")

	db.Save(&User{ID: 1, Name: "jinzhu", Age: 100})
	// UPDATE `users` SET `name`="jinzhu",`age`=100,`birthday`="0000-00-00 00:00:00",`update_at`="0000-00-00 00:00:00" WHERE `id` = 1

	///更新单个列
	// 根据条件更新
	db.Debug().Model(&User{}).Where("active = ?", true).Update("name", "hello")
	// UPDATE users SET name='hello', updated_at='2013-11-17 21:34:10' WHERE active=true;

	// User 的 ID 是 `111`
	db.Model(&user).Update("name", "hello")
	// UPDATE users SET name='hello', updated_at='2013-11-17 21:34:10' WHERE id=111;

	// 根据条件和 model 的值进行更新
	db.Debug().Model(&user).Where("active = ?", true).Update("name", "hello")
	// UPDATE users SET name='hello', updated_at='2013-11-17 21:34:10' WHERE id=111 AND active=true;

	///更新多列
	// 根据 `struct` 更新属性，只会更新非零值的字段
	db.Model(&user).Updates(User{Name: "hello", Age: 18, Active: false})
	// UPDATE users SET name='hello', age=18, updated_at = '2013-11-17 21:34:10' WHERE id = 111;

	// 根据 `map` 更新属性
	db.Model(&user).Updates(map[string]interface{}{"name": "hello", "age": 18, "active": false})
	// UPDATE users SET name='hello', age=18, active=false, updated_at='2013-11-17 21:34:10' WHERE id=111;

	///更新选定字段
	tx := db.Debug().Model(&user).Select("*").Omit("id").Updates(&User{Name: "hello", Age: 124, Active: true})
	err := tx.Error
	if err != nil {
		{
			fmt.Println(err)
		}
	}
	db.Debug().Model(&user).Omit("name").Updates(&User{Name: "hello", Age: 18, Active: true})
	db.Debug().Model(&user).Select("*").Omit("name").Updates(&User{Name: "hello", Age: 18, Active: true})
}

type User struct {
	ID     uint   `gorm:"primaryKey"`
	Name   string `gorm:"column:name"`
	Email  string `gorm:"column:email"`
	Age    int8   `gorm:"column:age"`
	Active bool   `gorm:"column:active"`
}

// 前置钩子函数
func (u *User) BeforeUpdate(tx *gorm.DB) (err error) {
	if u.Age > 100 {
		return errors.New("Age is greater than 100")
	}
	return
}

// InitDB 初始化数据库
func InitDB() *gorm.DB {
	db := ConnectDB()
	err := db.AutoMigrate(&User{})
	if err != nil {
		panic(err)
	}
	return db
}

// ConnectDB 连接数据库
func ConnectDB() *gorm.DB {
	db, err := gorm.Open(sqlite.Open(constant.DBPATH), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	return db
}
