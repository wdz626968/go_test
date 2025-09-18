package service

import "github.com/gin-gonic/gin"

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
