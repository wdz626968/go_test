package main

import "fmt"

func main() {
	var testInt = new(int)
	*testInt = 12
	fmt.Println("引用传递修改之前：", *testInt)
	testReference(testInt)
	fmt.Println("引用传递修改之后：", *testInt)
	fmt.Println("值传递修改之前：", *testInt)
	testValue(*testInt)
	fmt.Println("值传递修改之后：", *testInt)
}

func testReference(intptr *int) *int {
	*intptr += 10
	return intptr
}

func testValue(testInf int) int {
	testInf += 10
	return testInf
}
