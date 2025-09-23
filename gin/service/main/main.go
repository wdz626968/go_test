package main

import (
	"encoding/json"
	"github.com/gorilla/websocket"
	"log"
	"net/http"
)

/**
 * 订阅iTick美股信息
 */
func SubscribeITick() {
	url := "wss://api-free.itick.org/stock" // 替换为实际iTick WebSocket端点
	//添加request Header
	headers := http.Header{
		"token": []string{"c98991b1e51849a79e2e4e5d73c788a89d6231c63f6345fe9ae35c00f7b86884"},
	}

	conn, _, err := websocket.DefaultDialer.Dial(url, headers)
	if err != nil {
		log.Fatal("连接失败:", err)
	}
	defer conn.Close()
	// 发送订阅请求
	subMsg := SubscriptionMsg{
		Action:  "subscribe",
		Symbols: "AAPL$US,TSLA$US",
		Types:   "depth,quote",
	}
	msgBytes, _ := json.Marshal(subMsg)
	if err := conn.WriteMessage(websocket.TextMessage, msgBytes); err != nil {
		log.Fatal("订阅请求发送失败:", err)
	}
	// 接收行情数据
	go func() {
		for {
			_, message, err := conn.ReadMessage()
			if err != nil {
				log.Println("读取错误:", err)
				return
			}
			log.Println("接收到行情数据:", string(message))
			var tick TickData
			if err := json.Unmarshal(message, &tick); err != nil {
				log.Println("解析错误:", err)
				continue
			}
			log.Printf("行情更新: %s %.2f (量:%d)\n",
				tick.Symbol, tick.Price, tick.Volume)
		}
	}()
	select {}
}
func main() {
	SubscribeITick()
}

type SubscriptionMsg struct {
	Action  string `json:"ac"`
	Symbols string `json:"params"`
	Types   string `json:"types"`
}

type TickData struct {
	Symbol    string  `json:"symbol"`
	Price     float64 `json:"price"`
	Volume    int64   `json:"volume"`
	Timestamp int64   `json:"timestamp"`
}
