package global

import (
	"sync"

	"gorm.io/gorm"
)

// 统一响应结构
type Response struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

// 自定义错误类型
type AppError struct {
	Code    int
	Message string
	Err     error
}

func (e *AppError) Error() string {
	return e.Message
}

var (
	DB     *gorm.DB
	dbOnce sync.Once
)

// InitDB 使用sync.Once确保数据库连接只初始化一次
func InitDB(db *gorm.DB) {
	dbOnce.Do(func() {
		DB = db
	})
}
