package global

import (
	"sync"

	"gorm.io/gorm"
)

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
