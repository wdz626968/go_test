package main

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"log"
)

func main() {
	db1, err := initDB()
	if err != nil {
	}
	db1.AutoMigrate(&Employee{})
	employees := []Employee{{
		Name:       "dzhwu",
		Department: "技术部",
		Salary:     12,
	}, {
		Name:       "mia",
		Department: "技术部",
		Salary:     10,
	},
	}
	db1.Create(&employees)
	db, err := sqlx.Connect("sqlite3", "identifier.sqlite")
	if err != nil {
	}

	// 查询技术部员工
	techEmployees, err := queryTechDeptEmployees(db)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("技术部员工: %+v\n", techEmployees)

	maxSalaryEmployee, err := queryMaxSalary(db)
	if err != nil {
	}
	fmt.Println(maxSalaryEmployee)

	//查询查询价格大于 50 元的书籍，并将结果映射到 Book 结构体切片中，确保类型安全。
	db1.AutoMigrate(&Book{})
	books := []Book{{
		Title:  "hhh",
		Author: "dzhwu",
		Price:  100,
	}, {
		Title:  "hhh2",
		Author: "mia",
		Price:  120,
	},
	}
	db1.Create(&books)
	price, err := queryBooksByPrice(db)
	if err != nil {
	}
	fmt.Println(price)
}

// Employee 结构体映射数据库表
type Employee struct {
	ID         int    `db:"id"`
	Name       string `db:"name"`
	Department string `db:"department"`
	Salary     int    `db:"salary"`
}

type Book struct {
	ID     int    `db:"id"`
	Title  string `db:"title"`
	Author string `db:"author"`
	Price  int    `db:"price"`
}

func initDB() (*gorm.DB, error) {
	db, err := gorm.Open(sqlite.Open("identifier.sqlite"), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	return db, err
}

// 查询技术部员工
func queryTechDeptEmployees(db *sqlx.DB) ([]Employee, error) {
	var employees []Employee
	query := `SELECT id, name, department, salary FROM employees WHERE department = $1`
	err := db.Select(&employees, query, "技术部")
	return employees, err
}

func queryMaxSalary(db *sqlx.DB) (Employee, error) {
	var employee Employee
	query := `select * from employees where salary = (select max(salary) FROM employees) limit 1 `
	err := db.Get(&employee, query)
	return employee, err
}

func queryBooksByPrice(db *sqlx.DB) ([]Book, error) {
	var books []Book
	query := `SELECT * FROM books WHERE price > $1`
	err := db.Select(&books, query, 50)
	return books, err
}
