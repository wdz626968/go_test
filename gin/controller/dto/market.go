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
