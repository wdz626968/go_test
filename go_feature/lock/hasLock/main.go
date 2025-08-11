package main

import (
	"fmt"
	"sync"
)

func main() {
	count := 1
	mutex := sync.Mutex{}
	wg := sync.WaitGroup{}
	wg.Add(10)
	for i := 0; i < 10; i++ {
		go func() {
			defer wg.Done()
			//必须进行引用传递，不能进行值传递
			add(&mutex, &count)
		}()
	}
	wg.Wait()
	fmt.Println(count)
}

func add(mutex *sync.Mutex, count *int) {
	for i := 0; i < 1000; i++ {
		mutex.Lock()
		*count++
		mutex.Unlock()
	}
}
