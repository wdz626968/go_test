package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"
)

const (
	apiKey     = "rtJCR35B93aRr6gb6TvXqDmv_jc3DiNE"
	baseURL    = "https://api.polygon.io/v2/aggs/ticker"
	timeout    = 10 * time.Second
	companies  = "AAPL,TSLA,GOOGL,MSFT,AMZN,NVDA"
	fromDate   = "2025-09-10"
	toDate     = "2025-09-22"
	multiplier = 1
	timespan   = "day"
)

type StockResponse struct {
	Ticker       string `json:"ticker"`
	QueryCount   int    `json:"queryCount"`
	ResultsCount int    `json:"resultsCount"`
	Adjusted     bool   `json:"adjusted"`
	Results      []struct {
		Volume    float64 `json:"v"`
		Open      float64 `json:"o"`
		Close     float64 `json:"c"`
		High      float64 `json:"h"`
		Low       float64 `json:"l"`
		Timestamp int64   `json:"t"`
		ItemCount int     `json:"n"`
	} `json:"results"`
}

type StockSummary struct {
	Ticker         string  `json:"ticker"`
	AvgGain        float64 `json:"avg_gain"`
	AvgGainPercent float64 `json:"avg_gain_percent"`
	AvgVolume      float64 `json:"avg_volume"`
	CurrentPrice   float64 `json:"current_price"`
}

func main() {
	client := &http.Client{Timeout: timeout}
	stockCodes := []string{"AAPL", "TSLA", "GOOGL", "MSFT", "AMZN", "NVDA"}

	var summaries []StockSummary
	for _, ticker := range stockCodes {
		req, err := http.NewRequest("GET", fmt.Sprintf("%s/%s/range/%d/%s/%s/%s",
			baseURL, ticker, multiplier, timespan, fromDate, toDate), nil)
		if err != nil {
			log.Fatalf("创建请求失败: %v", err)
		}

		q := req.URL.Query()
		q.Add("apiKey", apiKey)
		req.URL.RawQuery = q.Encode()

		resp, err := client.Do(req)
		if err != nil {
			log.Printf("获取%s数据失败: %v", ticker, err)
			continue
		}
		defer resp.Body.Close()

		if resp.StatusCode != http.StatusOK {
			log.Printf("%s返回状态码: %d", ticker, resp.StatusCode)
			continue
		}

		var result StockResponse
		if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
			log.Printf("解析%s数据失败: %v", ticker, err)
			continue
		}

		if len(result.Results) < 7 {
			log.Printf("%s数据不足7天", ticker)
			continue
		}

		summary := calculate7DayStats(ticker, result.Results)
		summaries = append(summaries, summary)
	}

	output, err := json.MarshalIndent(summaries, "", "  ")
	if err != nil {
		log.Fatalf("JSON格式化失败: %v", err)
	}
	fmt.Println(string(output))
}

func calculate7DayStats(ticker string, data []struct {
	Volume    float64 `json:"v"`
	Open      float64 `json:"o"`
	Close     float64 `json:"c"`
	High      float64 `json:"h"`
	Low       float64 `json:"l"`
	Timestamp int64   `json:"t"`
	ItemCount int     `json:"n"`
}) StockSummary {
	var totalGain, totalGainPercent, totalVolume float64
	days := len(data)

	for i := 1; i < days; i++ {
		prevClose := data[i-1].Close
		currentClose := data[i].Close
		gain := currentClose - prevClose
		gainPercent := (gain / prevClose) * 100

		totalGain += gain
		totalGainPercent += gainPercent
		totalVolume += data[i].Volume
	}

	return StockSummary{
		Ticker:         ticker,
		AvgGain:        totalGain / float64(days-1),
		AvgGainPercent: totalGainPercent / float64(days-1),
		AvgVolume:      totalVolume / float64(days-1),
		CurrentPrice:   data[days-1].Close,
	}
}
