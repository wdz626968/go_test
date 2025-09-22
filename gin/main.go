package main

import (
	"go_test/gin/config"
	"go_test/gin/controller"
	"go_test/gin/global"
	"go_test/gin/middlerware"
	"go_test/gin/model"
	"net/http"

	"github.com/gin-gonic/gin"
)

// [客户端请求]
//
//	↓
//
// [Logger中间件(RequestLogger)] → 记录请求开始时间
//
//	↓
//
// [CORS中间件(CORSMiddleware)] → 处理跨域请求
//
//	↓
//
// [JWT鉴权(JWTAuthMiddleware)] → 验证访问令牌
//
//	↓
//
// [RBAC鉴权(RBACAuthMiddleware)] → 校验用户权限
//
//	↓
//
// [业务处理] → 核心业务逻辑
//
//	↓
//
// [Logger中间件(RequestLogger)] ← 记录响应耗时
func main() {
	router := gin.Default()
	router.Use(middlerware.ErrorHandler())
	router.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	router.GET("/welcome", func(c *gin.Context) {
		firstName := c.DefaultQuery("firstname", "Guest")
		c.String(http.StatusOK, "Hello %s", firstName)
	})
	// 添加CORS中间件，允许所有跨域请求
	router.Use(middlerware.CORSMiddleware())

	v1 := router.Group("/api/v1")
	{
		auth := v1.Group("/auth")
		{
			auth.POST("", controller.Register)
			auth.POST("/login", controller.Login)
			//users.GET("/:id", getUser)
		}

		users := v1.Group("/users")
		{
			users.GET("/:id([0-9]+)", controller.GetUser)
			//users.GET("/:id", getUser)
		}
		posts := v1.Group("/posts", middlerware.JWTAuthMiddleware(), middlerware.RBACAuthMiddleware([]string{"user", "admin"}))
		{
			posts.GET("/:id([0-9]+)", controller.GetPost)
			posts.GET("/all", controller.GetAllPosts)
			posts.POST("/", controller.CreatePosts)
			posts.PUT("/", controller.UpdatePosts)
			posts.DELETE("/", controller.DeletePosts)
		}

		comments := v1.Group("/comments", middlerware.JWTAuthMiddleware(), middlerware.RBACAuthMiddleware([]string{"user", "admin"}))
		{
			comments.GET("/id/:id([0-9]+)", controller.GetComment)
			comments.GET("/post_id/:post_id([0-9]+)", controller.GetCommentsNyPostId)
			comments.POST("/", controller.CreateComments)
			//comments.PUT("/", controller.UpdateComments)
			//comments.DELETE("/", controller.DeleteComments)
		}
		tokens := v1.Group("/tokens")
		{
			tokens.GET("/:mint", controller.GetTokenAccountsByOwner)
			//users.GET("/:id", getUser)
		}
		stocks := v1.Group("/stocks")
		{
			stocks.GET("/list", controller.GetStockList)
			stocks.GET("/info", controller.GetStockInfo)
			stocks.GET("/overview", controller.GetMarketOverview)
			//users.GET("/:id", getUser)
		}
	}
	router.Use(middlerware.RequestLogger())

	config.InitConfig()
	global.DB.AutoMigrate(&model.Comment{}, &model.Post{}, &model.User{})
	router.Run()
}
