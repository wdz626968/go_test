package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"sync"
	"time"

	"github.com/gorilla/websocket"
)

const (
	apiKey      = "rtJCR35B93aRr6gb6TvXqDmv_jc3DiNE"
	wsURL       = "wss://socket.polygon.io/stocks"
	restBaseURL = "https://api.polygon.io/v2/reference/financials"
)

type StockData struct {
	Ticker        string  `json:"ticker"`
	CurrentPrice  float64 `json:"currentPrice"`
	PriceChange   float64 `json:"priceChange"`
	ChangePercent float64 `json:"changePercent"`
	Volume        int64   `json:"volume"`
	MarketCap     int64   `json:"marketCap"`
	LastUpdated   string  `json:"lastUpdated"`
}

func main() {
	if apiKey == "" {
		log.Fatal("请设置POLYGON_API_KEY环境变量")
	}

	tickers := []string{"AAPL", "TSLA", "GOOGL", "MSFT", "AMZN", "NVDA"}
	results := make(chan StockData, len(tickers))
	var wg sync.WaitGroup

	// 获取市值数据
	marketCaps := fetchMarketCaps(tickers)

	// 连接WebSocket获取实时数据
	wg.Add(1)
	go func() {
		defer wg.Done()
		connectWebSocket(tickers, marketCaps, results)
	}()

	go func() {
		wg.Wait()
		close(results)
	}()

	var stockData []StockData
	for data := range results {
		stockData = append(stockData, data)
	}

	output, _ := json.MarshalIndent(stockData, "", "  ")
	fmt.Println(string(output))
}

func fetchMarketCaps(tickers []string) map[string]int64 {
	client := &http.Client{Timeout: 10 * time.Second}
	caps := make(map[string]int64)
	var mu sync.Mutex
	var wg sync.WaitGroup

	for _, ticker := range tickers {
		wg.Add(1)
		go func(t string) {
			defer wg.Done()
			url := fmt.Sprintf("%s/%s?apiKey=%s", restBaseURL, t, apiKey)
			resp, err := client.Get(url)
			if err != nil {
				return
			}
			defer resp.Body.Close()

			var result struct {
				Results []struct {
					MarketCap int64 `json:"marketCap"`
				} `json:"results"`
			}

			if err := json.NewDecoder(resp.Body).Decode(&result); err != nil || len(result.Results) == 0 {
				return
			}

			mu.Lock()
			caps[t] = result.Results[0].MarketCap
			mu.Unlock()
		}(ticker)
	}

	wg.Wait()
	return caps
}

func connectWebSocket(tickers []string, marketCaps map[string]int64, results chan<- StockData) {
	dialer := websocket.DefaultDialer

	conn, _, err := dialer.Dial(wsURL, nil)
	if err != nil {
		log.Fatal("WebSocket连接失败:", err)
	}
	defer conn.Close()

	// 认证
	authMsg := fmt.Sprintf(`{"action":"auth","params":"%s"}`, apiKey)
	if err := conn.WriteMessage(websocket.TextMessage, []byte(authMsg)); err != nil {
		log.Fatal("认证失败:", err)
	}

	// 订阅股票
	subscribeMsg := `{"action":"subscribe","params":"A.*"}`
	if err := conn.WriteMessage(websocket.TextMessage, []byte(subscribeMsg)); err != nil {
		log.Fatal("订阅失败:", err)
	}

	tickerMap := make(map[string]bool)
	for _, t := range tickers {
		tickerMap[t] = true
	}

	received := make(map[string]bool)
	timeout := time.After(30 * time.Second)

	for {
		select {
		case <-timeout:
			return
		default:
			_, message, err := conn.ReadMessage()
			if err != nil {
				log.Println("读取消息错误:", err)
				return
			}

			var msg struct {
				Ev     string  `json:"ev"`
				Ticker string  `json:"T"`
				Price  float64 `json:"p"`
				Change float64 `json:"c"`
				Pct    float64 `json:"P"`
				Volume int64   `json:"v"`
				Time   int64   `json:"t"`
			}

			if err := json.Unmarshal(message, &msg); err != nil {
				continue
			}

			if msg.Ev == "A" && tickerMap[msg.Ticker[2:]] {
				ticker := msg.Ticker[2:]
				if !received[ticker] {
					results <- StockData{
						Ticker:        ticker,
						CurrentPrice:  msg.Price,
						PriceChange:   msg.Change,
						ChangePercent: msg.Pct,
						Volume:        msg.Volume,
						MarketCap:     marketCaps[ticker],
						LastUpdated:   time.Unix(msg.Time/1000, 0).Format(time.RFC3339),
					}
					received[ticker] = true
				}
			}

			if len(received) == len(tickers) {
				return
			}
		}
	}
}
