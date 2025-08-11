package main

import (
	"fmt"
	"sync"
)

// 有缓冲允许发送和接收操作异步执行，只要缓冲区未满/未空就不会阻塞
// 无缓冲要求发送和接收操作必须同时准备好，否则会阻塞，本质是同步通信
// 无缓冲通道：事件通知、严格同步
// 有缓冲通道：任务队列、流量控制
func main() {
	group := sync.WaitGroup{}
	group.Add(2)
	channel := make(chan int, 10)
	go func() {
		defer group.Done()
		sendOnly(channel)
	}()

	go func() {
		defer group.Done()
		//非阻塞
		//time.Sleep(1 * time.Second)
		receiveOnly(channel)
	}()
	group.Wait()
	fmt.Println("通道执行逻辑结束，回到主线程")
}

func sendOnly(ch chan int) {
	for i := 1; i <= 100; i++ {
		ch <- i
		fmt.Println("发送channel信息:", i)
	}
}

func receiveOnly(ch chan int) {
	for i := 1; i <= 100; i++ {
		chInfo := <-ch
		fmt.Println("收到channel信息:", chInfo)
	}
}
