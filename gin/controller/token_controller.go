package controller

import (
	"context"
	"go_test/gin/service"
	"net/http"

	"github.com/blocto/solana-go-sdk/client"
	"github.com/blocto/solana-go-sdk/common"
	"github.com/blocto/solana-go-sdk/types"
	"github.com/gin-gonic/gin"
)

// FUarP2p5EnxD66vVDL4PWRoWMzA56ZVHG24hpEDFShEz
var feePayer, _ = types.AccountFromBase58("4TMFNY9ntAn3CHzguSAvDNLPRoQTaK3sWbQQXdDXaE6KWRBLufGL6PJdsD2koiEe3gGmMdRK3aAw7sikGNksHJrN")

// 9aE476sH92Vz7DMPyq5WLPkrKWivxeuTKEFKd2sZZcde
var alice, _ = types.AccountFromBase58("4voSPg3tYuWbKzimpQK9EbXHmuyy5fUrtXvpLDMLkmY6TRncaTHAKGD8jUg3maB5Jbrd9CkQg4qjJMyN6sQvnEF2")

var mintPubkey = common.PublicKeyFromString("F6tecPzBMF47yJ2EN6j2aGtE68yR5jehXcZYVZa6ZETo")

var tokenService = service.NewTokenService()
var stockService = service.NewStockService()

/**
 * 获取账户代币信息列表
 */
func GetTokenAccountsByOwner(ctx *gin.Context) {
	cli := client.NewClient("https://devnet.helius-rpc.com/?api-key=8a9947bf-2456-4824-b675-98bf7750e9ac")

	//mint := ctx.Param("mint")
	accounts, err := cli.GetTokenAccountsByOwnerByMint(context.Background(), alice.PublicKey.ToBase58(), mintPubkey.ToBase58())
	if err != nil {
		ctx.JSON(500, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(200, accounts)
	ctx.JSON(http.StatusCreated, accounts)
}

/*
 * 获取市场总览信息
 */
func GetMarketOverview(ctx *gin.Context) {
	//todo 对DB数据做聚合操作
	//ctx.JSON(http.StatusOK, marketService.GetMarketOverview())
}

/*
 * 获取股票列表信息
 */
func GetStockList(ctx *gin.Context) {
	//todo 分页查询DB中的股票与代币信息与价格等信息
	//todo 是否是实时去拿信息？ 通过代币信息调用聚合预言机合约拿到实时的价格、涨跌幅等构造前端所需要的数据结构
	//ctx.JSON(http.StatusOK, stockService.GetStockList())
}

/**
 * 买入
 */
func BuyStock(ctx *gin.Context) {
	//todo 调用token_service发行对应的代币
	tokenService.MintToken(ctx)
	//todo 调用第三方券商API用指定账号购买美股，如果在休市，则按照（一定规则）从cryptoStock的资产池中发股票
	stockService.BuyStock(ctx.Param("code"), ctx.Param("num"))
	//todo 监听脸上交易成功后，把流动性注入到池子中，要更新当前DB中的股票池信息
	stockService.UpdateStockPool(ctx)
	//ctx.JSON(http.StatusOK, stockService.BuyStock(ctx.Param("code"), ctx.Param("num")))
}

/**
 * 卖出
 */
func SellStock(ctx *gin.Context) {
	//todo 调用token_service销毁对应的代币
	tokenService.BurnToken(ctx)
	//todo 调用第三方券商API用指定账号卖出美股，如果在休市，则按照（一定规则）向cryptoStock的资产池中增加股票
	stockService.SellStock(ctx.Param("code"), ctx.Param("num"))
	//todo 监听脸上交易成功后，把流动性销毁，并且更新当前DB中的股票池信息
	//ctx.JSON(http.StatusOK, stockService.SellStock(ctx.Param()))
}

/**
 * 初始化股票池
 */
func InitStockPool(ctx *gin.Context) {
	//todo 调用token_service发行对应的代币
	tokenService.MintToken(ctx)
	//todo 调用第三方券商API用指定账号卖出美股，如果在休市，则按照（一定规则）向cryptoStock的资产池中增加股票
	stockService.BuyStock(ctx.Param("code"), ctx.Param("num"))
	//todo 股票、代币绑定信息入库
}
