package main

import (
	"go_test/gin/config"
	"go_test/gin/controller"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	router.GET("/welcome", func(c *gin.Context) {
		firstName := c.DefaultQuery("firstname", "Guest")
		c.String(http.StatusOK, "Hello %s", firstName)
	})
	// RESTful 路由示例
	v1 := router.Group("/api/v1")
	{
		users := v1.Group("/users")
		{
			users.GET("/:id", controller.GetUser)
			users.POST("", controller.CreateUser)
			//users.GET("/:id", getUser)
		}
	}
	config.InitConfig()
	router.Run()
}

func getUser(context *gin.Context) {

}
