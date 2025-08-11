package main

import (
	"fmt"
	"sync"
)

// 分十个协程更新1w条数据
func main() {
	wg := sync.WaitGroup{}
	wg.Add(10)
	data := []int{}
	for i := 0; i < 10000; i++ {
		data = append(data, i)
	}
	for i := 0; i < 10; i++ {
		go func() {
			defer wg.Done()
			dataCleanJob(data[100*i : (i+1)*100])
		}()
	}
	wg.Wait()
	fmt.Println("回到主线程，并退出")

}

func dataCleanJob(dataJob []int) {
	for i := range dataJob {
		updateData(dataJob[i])
	}
}

func updateData(dataJob int) {
	fmt.Println("更新数据", dataJob)
}
