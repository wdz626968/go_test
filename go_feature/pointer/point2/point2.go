package main

import "fmt"

func main() {
	var slicePtr = []int{1, 2, 3}
	fmt.Println("切片修改之前：", slicePtr)
	doubleSlice(&slicePtr)
	fmt.Println("切片修改之后：", slicePtr)
}

func doubleSlice(slice *[]int) {
	for i := range *slice {
		(*slice)[i] *= 2
	}
}
