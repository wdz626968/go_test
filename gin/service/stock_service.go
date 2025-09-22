package service

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"go_test/gin/controller/dto"
	"io"
	"log"
	"net/http"
	"strconv"
	"time"
)

type StockService struct {
}

func NewStockService() *StockService {
	return &StockService{}
}

/**
 * 购买股票
 */
func (s StockService) BuyStock(param string, param2 string) {
	//todo 调用第三方券商API用指定账号购买美股，如果在休市，则按照（一定规则）从cryptoStock的资产池中发股票

}

/**
 * 卖出股票
 */
func (s StockService) SellStock(param string, param2 string) {
	//todo 调用第三方券商API用指定账号卖出美股，如果在休市，则按照（一定规则）向cryptoStock的资产池中增加股票
}

/**
 * 初始化股票池
 */
func (s StockService) InitStockPool(ctx *gin.Context) {

}

func (s StockService) UpdateStockPool(ctx *gin.Context) {

}

func (s StockService) GetStockList() []dto.StockFromITick {
	url := "https://api.itick.org/symbol/list?type=stock&region=US"

	req, _ := http.NewRequest("GET", url, nil)

	req.Header.Add("accept", "application/json")
	req.Header.Add("token", "c98991b1e51849a79e2e4e5d73c788a89d6231c63f6345fe9ae35c00f7b86884")

	res, _ := http.DefaultClient.Do(req)

	defer res.Body.Close()
	body, _ := io.ReadAll(res.Body)
	stocks, err := ConvertStocks(body)
	if err != nil {
		log.Fatal(err)
	}
	return stocks
}

/**
 * 获取股票具体信息
 */
func (s StockService) GetStockInfos(codes string) (dto.StockData, error) {
	url := "https://api.itick.org/stock/info?type=stock&region=US&code=" + codes

	req, _ := http.NewRequest("GET", url, nil)

	req.Header.Add("accept", "application/json")
	req.Header.Add("token", "c98991b1e51849a79e2e4e5d73c788a89d6231c63f6345fe9ae35c00f7b86884")

	res, _ := http.DefaultClient.Do(req)

	defer res.Body.Close()
	body, _ := io.ReadAll(res.Body)

	fmt.Println(string(body))
	var response dto.StockInfoResponse
	err := json.Unmarshal(body, &response)
	var stockData dto.StockData
	if err != nil {
		return stockData, err
	}
	return response.Data, nil
}
func ConvertStocks(jsonData []byte) ([]dto.StockFromITick, error) {
	var response dto.StockListResponse
	if err := json.Unmarshal(jsonData, &response); err != nil {
		return nil, fmt.Errorf("解析JSON失败: %v", err)
	}

	if response.Code != 0 {
		return nil, fmt.Errorf("API返回错误: %s", response.Msg)
	}

	return response.Data, nil
}

func (s StockService) CalculateAvgGain(kType string, limit string, codes string) (map[string]dto.AvgInfo, error) {
	client := &http.Client{Timeout: 10 * time.Second}
	req, _ := http.NewRequest("GET", "https://api.itick.org/stock/klines", nil)

	q := req.URL.Query()
	q.Add("codes", codes)
	q.Add("region", "US")
	q.Add("kType", kType)
	q.Add("limit", limit)
	req.URL.RawQuery = q.Encode()

	req.Header.Add("accept", "application/json")
	req.Header.Add("token", "c98991b1e51849a79e2e4e5d73c788a89d6231c63f6345fe9ae35c00f7b86884")
	log.Println("req:", req)

	resp, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("请求失败: %v", err)
	}
	defer resp.Body.Close()
	fmt.Println("resp:", resp)
	var result dto.ITickResponse
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, fmt.Errorf("数据解析错误: %v", err)
	}

	var avgInfo = make(map[string]dto.AvgInfo)
	var totalGain = 0.0
	var totalGainPercent = 0.0
	var totalVolume = 0.0
	validDays := 0
	for key, tick := range result.Data {
		atoi, err := strconv.Atoi(limit)
		if len(tick) < atoi || err != nil {
			return nil, fmt.Errorf("数据不足8个交易日")
		}
		if len(tick) == 1 {
			totalVolume += tick[len(tick)-1].Volume
			info := dto.AvgInfo{
				StockChange:        totalGain,
				StockChangePercent: totalGainPercent,
				StockChangeVolume:  totalVolume,
				Price:              tick[len(tick)-1].Close,
			}
			avgInfo[key] = info
		} else {
			for i := 1; i < len(tick); i++ {
				prevClose := tick[i-1].Close
				currentClose := tick[i].Close
				if prevClose == 0 {
					continue // 跳过除零情况
				}
				dailyGain := currentClose - prevClose
				totalGain += dailyGain
				dailyGainPercent := (currentClose - prevClose) / prevClose * 100
				totalGainPercent += dailyGainPercent
				totalVolume += tick[i-1].Volume
				validDays++
			}
			info := dto.AvgInfo{
				StockChange:        totalGain / float64(validDays),
				StockChangePercent: totalGainPercent / float64(validDays),
				StockChangeVolume:  totalVolume / float64(validDays),
				Price:              tick[len(tick)-1].Close,
			}
			avgInfo[key] = info

		}
	}

	//if validDays == 0 {
	//	return nil, fmt.Errorf("无有效数据可计算")
	//}

	log.Println("totalGain", totalGain)
	log.Println("validDays", validDays)

	return avgInfo, nil
}

func (s StockService) CalculateMarketOverview(info map[string]dto.AvgInfo) dto.Overview {
	totalChange24h := 0.0
	totalMarketCap := 0.0
	totalStockNum := len(info)
	totalVolume24h := 0.0
	for _, value := range info {
		totalChange24h += value.StockChange
		//totalMarketCap += value.StockNumValue
		totalVolume24h += value.StockChangeVolume
	}
	return dto.Overview{
		TotalChange24h: totalChange24h,
		TotalMarketCap: totalMarketCap,
		TotalStockNum:  totalStockNum,
		TotalVolume24h: totalVolume24h,
	}
}
