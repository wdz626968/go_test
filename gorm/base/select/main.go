package main

import (
	"fmt"
	"go_test/gorm/constant"
	"time"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

func main() {
	db := InitDB()
	user := User{ID: 1, Name: "John Doe"}
	//// 获取第一条记录（主键升序）
	//// SELECT * FROM users ORDER BY id LIMIT 1;
	//db.First(&user)
	//// 获取一条记录，没有指定排序字段
	//// SELECT * FROM users LIMIT 1;
	//db.Take(&user)
	//// 获取最后一条记录（主键降序）
	//// SELECT * FROM users ORDER BY id DESC LIMIT 1;
	//db.Last(&user)
	//db.Limit(1).Find(&user)
	//result := map[string]interface{}{}
	//db.Debug().Model(&User{}).First(&result)
	//fmt.Println(result)
	//
	////根据主键查询
	//db.First(&user, 10)
	//// SELECT * FROM users WHERE id = 10;
	//
	//db.First(&user, "10")
	//// SELECT * FROM users WHERE id = 10;
	//
	//db.Find(&user, []int{1, 2, 3})
	//// SELECT * FROM users WHERE id IN (1,2,3);
	//db.First(&user, "id = ?", "1b74413f-f3b8-409f-ac47-e8c062e3472a")
	//// SELECT * FROM users WHERE id = "1b74413f-f3b8-409f-ac47-e8c062e3472a";
	//
	//var user1 = User{ID: 10}
	//db.First(&user1)
	//// SELECT * FROM users WHERE id = 10;
	//
	//var result1 User
	//db.Model(User{ID: 10}).First(&result1)
	// SELECT * FROM users WHERE id = 10;
	//string 条件
	tx := db.Debug().Where("id = ?", user.ID).Find(&user)
	fmt.Println("tx:", *tx)
	db.Where("name <> ?", "jinzhu").Find(&user)
	// SELECT * FROM users WHERE name <> 'jinzhu';
	db.Where("name IN ?", []string{"jinzhu", "jinzhu 2"}).Find(&user)
	// SELECT * FROM users WHERE name IN ('jinzhu','jinzhu 2');

	// LIKE
	db.Where("name LIKE ?", "%jin%").Find(&user)
	// SELECT * FROM users WHERE name LIKE '%jin%';

	//AND
	db.Where("name = ? AND age >= ?", "jinzhu", "22").Find(&user)
	// SELECT * FROM users WHERE name = 'jinzhu' AND age >= 22;
	lastWeek := time.Now().AddDate(0, 0, -2)
	today := time.Now()
	// Time
	db.Where("updated_at > ?", lastWeek).Find(&user)
	// SELECT * FROM users WHERE updated_at > '2000-01-01 00:00:00';

	// BETWEEN
	db.Debug().Where("created_at BETWEEN ? AND ?", lastWeek, today).Find(&user)
	// SELECT * FROM users WHERE created_at BETWEEN '2000-01-01 00:00:00' AND '2000-01-08 00:00:00';

	//Struct & Map 条件
	//struct
	db.Debug().Where(&User{Name: "John Doe"}).Find(&user)

	//Map
	var users []User
	db.Debug().Where(map[string]interface{}{"name": "John Doe", "age": 18}).Find(&users)
	fmt.Println(users)

	// Slice of primary keys
	db.Where([]int64{3, 4, 5}).Find(&users)
	fmt.Println(users)
	// SELECT * FROM users WHERE id IN (20, 21, 22);

	//Struct零值不会当作条件
	db.Where(&User{Name: "jinzhu", Age: 0}).Find(&users)
	// SELECT * FROM users WHERE name = "jinzhu";
	//Map中的零值会当作条件
	db.Where(map[string]interface{}{"Name": "jinzhu", "Age": 0}).Find(&users)
	// SELECT * FROM users WHERE name = "jinzhu" AND age = 0;

	//内联条件
	// Get by primary key if it were a non-integer type
	db.First(&user, "id = ?", "string_primary_key")
	// SELECT * FROM users WHERE id = 'string_primary_key';

	// Plain SQL
	db.Find(&user, "name = ?", "jinzhu")
	// SELECT * FROM users WHERE name = "jinzhu";

	db.Find(&users, "name <> ? AND age > ?", "jinzhu", 20)
	// SELECT * FROM users WHERE name <> "jinzhu" AND age > 20;

	// Struct
	db.Find(&users, User{Age: 20})
	// SELECT * FROM users WHERE age = 20;

	// Map
	db.Find(&users, map[string]interface{}{"age": 20})
	// SELECT * FROM users WHERE age = 20;

	//Not条件
	db.Not("name = ?", "jinzhu").First(&user)
	// SELECT * FROM users WHERE NOT name = "jinzhu" ORDER BY id LIMIT 1;

	// Not In
	db.Not(map[string]interface{}{"name": []string{"jinzhu", "jinzhu 2"}}).Find(&users)
	// SELECT * FROM users WHERE name NOT IN ("jinzhu", "jinzhu 2");

	// Struct
	db.Not(User{Name: "jinzhu", Age: 18}).First(&user)
	// SELECT * FROM users WHERE name <> "jinzhu" AND age <> 18 ORDER BY id LIMIT 1;

	// Not In slice of primary keys
	db.Not([]int64{1, 2, 3}).First(&user)
	// SELECT * FROM users WHERE id NOT IN (1,2,3) ORDER BY id LIMIT 1;

	//or条件
	db.Where("role = ?", "admin").Or("role = ?", "super_admin").Find(&users)
	// SELECT * FROM users WHERE role = 'admin' OR role = 'super_admin';

	// Struct
	db.Where("name = 'jinzhu'").Or(User{Name: "jinzhu 2", Age: 18}).Find(&users)
	// SELECT * FROM users WHERE name = 'jinzhu' OR (name = 'jinzhu 2' AND age = 18);

	// Map
	db.Where("name = 'jinzhu'").Or(map[string]interface{}{"name": "jinzhu 2", "age": 18}).Find(&users)
	// SELECT * FROM users WHERE name = 'jinzhu' OR (name = 'jinzhu 2' AND age = 18);

	//选择特定字段
	db.Select("name", "age").Find(&users)
	// SELECT name, age FROM users;

	db.Select([]string{"name", "age"}).Find(&users)
	// SELECT name, age FROM users;

	db.Table("users").Select("COALESCE(age,?)", 42).Rows()
	// SELECT COALESCE(age,'42') FROM users;

	//排序
	db.Order("age desc,name").Find(&users)
	db.Order("age desc").Order("name").Find(&users)
	db.Clauses(clause.OrderBy{
		Expression: clause.Expr{SQL: "FIELD(id,?)", Vars: []interface{}{[]int{1, 2, 3}}, WithoutParentheses: true},
	}).Find(&User{})

	//limit offset
	db.Limit(3).Find(&users)
	// SELECT * FROM users LIMIT 3;

	users1 := User{}
	users2 := User{}
	// Cancel limit condition with -1
	db.Limit(10).Find(&users1).Limit(-1).Find(&users2)
	// SELECT * FROM users LIMIT 10; (users1)
	// SELECT * FROM users; (users2)

	db.Offset(3).Find(&users)
	// SELECT * FROM users OFFSET 3;

	db.Limit(10).Offset(5).Find(&users)
	// SELECT * FROM users OFFSET 5 LIMIT 10;

	// Cancel offset condition with -1
	db.Offset(10).Find(&users1).Offset(-1).Find(&users2)
	// SELECT * FROM users OFFSET 10; (users1)
	// SELECT * FROM users; (users2)

	//Group by & having
	type Result struct {
		Date  time.Time
		Total int
	}
	result := Result{}
	db.Model(&User{}).Select("name, sum(age) as total").Where("name LIKE ?", "group%").Group("name").First(&result)
	// SELECT name, sum(age) as total FROM `users` WHERE name LIKE "group%" GROUP BY `name` LIMIT 1

	db.Model(&User{}).Select("name, sum(age) as total").Group("name").Having("name = ?", "group").Find(&result)
	// SELECT name, sum(age) as total FROM `users` GROUP BY `name` HAVING name = "group"

	db.Table("orders").Select("date(created_at) as date, sum(amount) as total").Group("date(created_at)").Rows()

	db.Table("orders").Select("date(created_at) as date, sum(amount) as total").Group("date(created_at)").Having("sum(amount) > ?", 100).Rows()

	db.Table("orders").Select("date(created_at) as date, sum(amount) as total").Group("date(created_at)").Having("sum(amount) > ?", 100).Scan(&result)

	//distinct
	db.Distinct("name", "age").Order("name, age desc").Find(User{})

	type result1 struct {
		Name  string
		Email string
	}
	//join
	db.Model(&User{}).Select("users.name, emails.email").Joins("left join emails on emails.user_id = users.id").Scan(result1{})
	// SELECT users.name, emails.email FROM `users` left join emails on emails.user_id = users.id
	db.Table("users").Select("users.name, emails.email").Joins("left join emails on emails.user_id = users.id").Scan(&result)

	db.Joins("JOIN emails ON emails.user_id = users.id AND emails.email = ?", "jinzhu@example.org").Joins("JOIN credit_cards ON credit_cards.user_id = users.id").Where("credit_cards.number = ?", "411111111111").Find(&user)

	db.Joins("Company").Find(&users)
	// SELECT `users`.`id`,`users`.`name`,`users`.`age`,`Company`.`id` AS `Company__id`,`Company`.`name` AS `Company__name` FROM `users` LEFT JOIN `companies` AS `Company` ON `users`.`company_id` = `Company`.`id`;

	// inner join
	db.InnerJoins("Company").Find(&users)
	// SELECT `users`.`id`,`users`.`name`,`users`.`age`,`Company`.`id` AS `Company__id`,`Company`.`name` AS `Company__name` FROM `users` INNER JOIN `companies` AS `Company` ON `users`.`company_id` = `Company`.`id`;

	// Raw SQL
	db.Raw("SELECT name, age FROM users WHERE name = ?", "Antonio").Scan(&result)
}

type User struct {
	ID    uint   `gorm:"primaryKey"`
	Name  string `gorm:"column:name"`
	Email string `gorm:"column:email"`
	Age   int8   `gorm:"column:age"`
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
