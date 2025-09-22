package dto

/**
 * @Description: 市场总览
 */
type marketOverview struct {
	//总市值
	TotalMarketCap float64 `json:"total_market_cap"`
	//24h涨幅
	TotalChange24h float64 `json:"total_change_24h"`
	//24h成交量
	TotalVolume24h float64 `json:"total_volume_24h"`
	//可投资股票种类数量
	TotalStockNum int `json:"total_stock_num"`
}

/**
 * @Description: 股票信息
 */
type Stock struct {
	//股票ID
	ID uint `json:"id"`
	//股票名称
	Name string `json:"name"`
	//股票代码
	Code string `json:"code"`
	//股票缩略图
	Image string `json:"image"`
	//股票对应的代币地址
	TokenAddress string `json:"token_address"`
	//股票简述
	Description string `json:"description"`
	//当前市场价格
	StockValue float64 `json:"stock_value"`
	//当前涨幅
	StockChange float64 `json:"market_change"`
	//当前涨幅百分比
	StockChangePercent float64 `json:"stock_change_percent"`
	//股票市场成交量
	StockVolume float64 `json:"stock_volume"`
	//总市值
	StockCap float64 `json:"stock_cap"`
	//24h涨幅
	StockChange24h float64 `json:"market_change_24h"`
	//7日涨幅
	StockChange7d float64 `json:"market_change_7d"`
	//币股池总共持有数量
	StockNum int `json:"stock_num"`
	//币股池持有总价
	StockNumValue float64 `json:"stock_num_value"`
}
type StockFromITick struct {
	Symbol     string `json:"c"`
	Name       string `json:"n"`
	Type       string `json:"t"`
	Exchange   string `json:"e"`
	Sector     string `json:"s"`
	Identifier string `json:"l"`
}

type StockListResponse struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
	//data改为泛型
	Data []StockFromITick `json:"data"`
}
type StockInfoResponse struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
	//data改为泛型
	Data StockData `json:"data"`
}

type StockData struct {
	Symbol         string  `json:"c"`   // 股票代码
	Name           string  `json:"n"`   // 公司名称
	Type           string  `json:"t"`   // 类型
	Exchange       string  `json:"e"`   // 交易所
	Sector         string  `json:"s"`   // 行业
	Industry       string  `json:"i"`   // 子行业
	LegalName      string  `json:"l"`   // 法律名称
	Currency       string  `json:"r"`   // 货币
	BusinessDesc   string  `json:"bd"`  // 业务描述
	WebsiteURL     string  `json:"wu"`  // 官网URL
	MarketCap      float64 `json:"mcb"` // 市值(以基础货币计)
	TotalSharesOut float64 `json:"tso"` // 流通股总数
	PE             float64 `json:"pet"` // 市盈率
	FinancialCurr  string  `json:"fcc"` // 财务报告货币

}

type ITickResponse struct {
	Code int               `json:"code"`
	Data map[string][]Tick `json:"data"`
}

type Tick struct {
	Close  float64 `json:"c"`
	Time   int64   `json:"t"`
	High   float64 `json:"h"`
	Low    float64 `json:"l"`
	Open   float64 `json:"o"`
	Volume float64 `json:"v"`
}

type StockInfo struct {
	Symbol               string  `json:"c"`                       // 股票代码
	Name                 string  `json:"n"`                       // 公司名称
	BusinessDesc         string  `json:"bd"`                      // 业务描述
	Price                float64 `json:"price"`                   //当前价格
	MarketCap            float64 `json:"mcb"`                     // 市值(以基础货币计)
	StockChange7d        float64 `json:"stock_change_7d"`         // 7日涨跌幅
	StockChangePercent7d float64 `json:"stock_change_percent_7d"` //7日涨跌幅百分比
	StockChange7dVolume  float64 `json:"stock_change_7d_volume"`  //7日均成交量
	StockNum             int     `json:"stock_num"`               //币股池持有该股票的数量
	StockNumValue        float64 `json:"stock_num_value"`         //币股池持有该股票的数量价值
}

type AvgInfo struct {
	StockChange        float64 `json:"stock_change"`         // 涨跌幅
	StockChangePercent float64 `json:"stock_change_percent"` //涨跌幅百分比
	StockChangeVolume  float64 `json:"stock_change_volume"`  //均成交量
	Price              float64 `json:"price"`                //当前价格
}

type Overview struct {
	TotalChange24h float64 `json:"total_change_24h"`
	TotalMarketCap float64 `json:"total_market_cap"`
	TotalStockNum  int     `json:"total_stock_num"`
	TotalVolume24h float64 `json:"total_volume_24h"`
}
