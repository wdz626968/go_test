package main

import "fmt"

func main() {
	employee := Employee{
		Person: Person{
			Name: "dzhwu",
			Age:  18,
		},
		EmployeeId: 1,
	}
	employee.PrintInfo()
}

type Person struct {
	Name string
	Age  uint
}

type Employee struct {
	Person
	EmployeeId int
}

func (e *Employee) PrintInfo() {
	fmt.Println("employee:", *e)
}
