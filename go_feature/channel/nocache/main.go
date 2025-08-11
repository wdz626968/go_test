package main

import (
	"fmt"
	"sync"
)

func main() {
	group := sync.WaitGroup{}
	group.Add(2)
	channel := make(chan int)
	go func() {
		defer group.Done()
		sendOnly(channel)
	}()

	go func() {
		defer group.Done()
		//阻塞
		//time.Sleep(1 * time.Second)
		receiveOnly(channel)
	}()
	group.Wait()
	fmt.Println("通道执行逻辑结束，回到主线程")
}

func sendOnly(ch chan int) {
	for i := 1; i <= 10; i++ {
		ch <- i
		fmt.Println("发送channel信息:", i)
	}
}

func receiveOnly(ch chan int) {
	for i := 1; i <= 10; i++ {
		chInfo := <-ch
		fmt.Println("收到channel信息:", chInfo)
	}
}
