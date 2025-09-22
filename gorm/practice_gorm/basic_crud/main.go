package main

import (
	"fmt"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type Student struct {
	ID    uint
	Name  string
	Age   int
	Grade string
	gorm.Model
}

func main() {
	db, err := initDB()
	if err != nil {
		panic(err)
	}

	err = db.AutoMigrate(&Student{})
	if err != nil {
		panic(err)
	}
	//编写SQL语句向 students 表中插入一条新记录，学生姓名为 "张三"，年龄为 20，年级为 "三年级"。
	student := Student{Name: "张三", Age: 10, Grade: "三年级"}
	db.Debug().Create(&student)

	//编写SQL语句查询 students 表中所有年龄大于 18 岁的学生信息。
	db.Debug().First(&Student{}, "age > 18")

	result := Student{}

	db.Debug().Select("*").Where("age>18").Find(&result)
	fmt.Println(result)

	//编写SQL语句将 students 表中姓名为 "张三" 的学生年级更新为 "四年级"。
	db.Debug().Model(&Student{}).Where("name=?", "张三").Update("grade", "四年级")

	//编写SQL语句删除 students 表中年龄小于 15 岁的学生记录。
	//db.Debug().Where("age < ?", 15).Delete(&Student{})
	db.Debug().Delete(&Student{}, "age < ?", 15)

	db.Debug().Select("*").Where("id=?", 10).Find(&result)
	fmt.Println(result)
}

func initDB() (*gorm.DB, error) {
	db, err := gorm.Open(sqlite.Open("identifier.sqlite"), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	return db, err
}
