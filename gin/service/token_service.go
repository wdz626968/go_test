package service

import "github.com/gin-gonic/gin"

type TokenService struct {
}

func NewTokenService() *TokenService {
	return &TokenService{}
}

/**
 * 发行代币
 */
func (s TokenService) MintToken(ctx *gin.Context) {
	//todo 调用token_service发行对应的代币
	//todo 调用合约转账USDT/USDC到合约账号
	//todo 调用合约向指定账号转账代币
}

func (s TokenService) BurnToken(ctx *gin.Context) {
	//todo 调用token_service销毁对应的代币
	//todo 向用户调用合约转账USDT/USDC
	//todo 调用合约向指定账号将代币转出
}

/**
 * 查询CryptoStock发布的代币列表
 */
func GetTokenList(ctx *gin.Context) {

}
