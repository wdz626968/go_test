package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

func main() {
	var count int32
	var wg sync.WaitGroup
	wg.Add(10)
	for i := 0; i < 10; i++ {
		go func() {
			defer wg.Done()
			add(&count)
		}()
	}
	wg.Wait()
	fmt.Println("count:", count)
}

func add(count *int32) {
	for i := 1; i <= 1000; i++ {
		atomic.AddInt32(count, 1)
	}

}
