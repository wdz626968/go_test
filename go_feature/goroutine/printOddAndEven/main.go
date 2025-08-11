package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	// 等待两个goroutine完成
	wg.Add(2)

	go func() {
		defer wg.Done()
		printEven()
	}()

	go func() {
		defer wg.Done()
		printOdd()
	}()
	// 等待所有goroutine完成
	wg.Wait()
	fmt.Println("主线程")
}

func printOdd() {
	for i := 1; i < 10; i++ {
		if i%2 != 0 {
			fmt.Println(i, "is odd")
		}
	}
}

func printEven() {
	for i := 2; i <= 10; i++ {
		if i%2 == 0 {
			fmt.Println(i, "is even")
		}
	}
}
