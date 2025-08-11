package main

import (
	"strconv"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type Product struct {
	Code  string
	Price uint
	gorm.Model
}

type User struct {
	gorm.Model
	Name       string
	CreditCard CreditCard
	Age        *int `gorm:"default:18"`
}

type CreditCard struct {
	gorm.Model
	Number string
	UserID uint
}

func main() {
	db, err := gorm.Open(sqlite.Open("identifier.sqlite"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	// Migrate the schema
	//db.AutoMigrate(&Product{})
	db.AutoMigrate(&User{})
	db.AutoMigrate(&CreditCard{})
	// Create
	////Insert 单个添加
	//p := Product{
	//	Code:  "123",
	//	Price: 1,
	//	Model: gorm.Model{},
	//}
	//db.Create(&p)
	////Insert 批量添加
	//products := []Product{
	//	{
	//		Code:  "333",
	//		Price: 2,
	//		Model: gorm.Model{},
	//	}, {
	//		Code:  "444",
	//		Price: 3,
	//		Model: gorm.Model{},
	//	},
	//}
	//tx := db.Create(&products)
	//fmt.Println(tx.RowsAffected)
	//fmt.Println(tx.Error)
	//Insert 指定字段
	//db.Debug().Select("code").Create(&products)
	//Insert 忽略某些字段
	//db.Debug().Omit("code").Create(&products)

	//分批次插入batch Insert
	//db.Debug().CreateInBatches(&products, 100)
	//map[string]interface{} []map[string]interface{}{} 创建记录(当使用map来创建时，钩子方法不会执行，关联不会被保存且不会回写主键)
	//db.Model(&products[0]).Create(map[string]interface{}{"code": "333", "price": 2})
	//db.Model(Product{}).Create([]map[string]interface{}{{"code": "2333", "price": 2}, {"code": "3444", "price": 3}})

	//关联创建
	// INSERT INTO `users` ...
	// INSERT INTO `credit_cards` ...
	user := User{
		Name: "John Doe",
		CreditCard: CreditCard{
			Model:  gorm.Model{},
			Number: "1234132312312312",
		},
	}
	db.Debug().Create(&user)
	//跳过关联更新
	db.Omit("CreditCard").Create(&user)
	db.Omit(clause.Associations).Create(&user)

	//Upsert
	db.Clauses(clause.OnConflict{
		Columns:   []clause.Column{{Name: "id"}},
		DoUpdates: clause.AssignmentColumns([]string{"name", "age"}),
	}).Create(&user)

	db.Clauses(clause.OnConflict{
		UpdateAll: true,
	}).Create(&user)

	//// Read
	//var product Product
	//db.First(&product, 1)                        // find product with integer primary key
	//db.First(&product, "code = ?", "D42")        // find product with code D42
	//db.Debug().Find(&product, "code = ?", "D42") // find product with code D42
	//
	//// Update - update product's price to 200
	//db.Model(&product).Update("Price", 200)
	//// Update - update multiple fields
	//db.Model(&product).Updates(Product{Price: 200, Code: "F42"}) // non-zero fields
	//db.Model(&product).Updates(map[string]interface{}{"Price": 200, "Code": "F42"})
	//
	//// Delete - delete product
	//db.Delete(&product, 1)
}

// BeforeCreate钩子函数
func (product *Product) BeforeCreate(db *gorm.DB) (err error) {
	product.Code += strconv.Itoa(int(product.Price))
	return
}
