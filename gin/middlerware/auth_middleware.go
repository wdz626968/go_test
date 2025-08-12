package middlerware

import (
	"errors"
	global "go_test/gin/global"
	"go_test/gin/utils"
	"log"
	"net/http"
	"runtime/debug"
	"strings"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

// parseTokenFromRequest 从请求中解析JWT token
func parseTokenFromRequest(c *gin.Context) (*utils.UserClaims, error) {
	tokenString := c.GetHeader("Authorization")
	if tokenString == "" {
		return nil, errors.New("未提供Token")
	}

	if strings.HasPrefix(tokenString, "Bearer ") {
		tokenString = strings.TrimPrefix(tokenString, "Bearer ")
	}

	return utils.ParseJWT(tokenString)
}

func authMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 1. 基础JWT认证
		userClaims, err := parseTokenFromRequest(c)
		if err != nil {
			if err.Error() == "未提供Token" {
				c.JSON(http.StatusUnauthorized, gin.H{"error": "未提供Token"})
			} else {
				c.JSON(http.StatusUnauthorized, gin.H{"error": "Token无效或已过期"})
			}
			c.Abort()
			return
		}

		// 2. 验证JWT中必须包含用户ID
		if userClaims.UserID == 0 {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Token格式无效，请重新登录"})
			c.Abort()
			return
		}
		c.Next()
	}
}

func rbacAuthMiddleware(roles []string) gin.HandlerFunc {
	return func(c *gin.Context) {
		userClaims, _ := parseTokenFromRequest(c)
		for i := range roles {
			if userClaims.Role == roles[i] {
				c.Next()
			}
		}
		c.JSON(http.StatusUnauthorized, gin.H{"error": "角色权限不足，请重新登录"})
		c.Abort()
		return
	}
}

// AuthMiddleware 基础JWT认证中间件
func JWTAuthMiddleware() gin.HandlerFunc {
	return authMiddleware()
}

// RBACAuthMiddleware 角色认证中间件
func RBACAuthMiddleware(roles []string) gin.HandlerFunc {
	return rbacAuthMiddleware(roles)
}

func CORSMiddleware() gin.HandlerFunc {
	return cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS", "PATCH"},
		AllowHeaders:     []string{"*"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	})
}

// 统一错误处理中间件
func ErrorHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				log.Printf("Panic: %v\n%s", err, debug.Stack())
				c.AbortWithStatusJSON(http.StatusInternalServerError,
					global.Response{
						Code:    500,
						Message: "Internal Server Error",
						Data:    "Unexpected server error occurred",
					})
			}
		}()

		c.Next()

		if len(c.Errors) > 0 {
			err := c.Errors.Last()
			switch e := err.Err.(type) {
			case *global.AppError:
				c.AbortWithStatusJSON(e.Code, global.Response{
					Code:    e.Code,
					Message: e.Message,
				})
			default:
				c.AbortWithStatusJSON(http.StatusBadRequest,
					global.Response{
						Code:    400,
						Message: "Bad Request",
						Data:    err.Error(),
					})
			}
		}
	}
}

func RequestLogger() gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()
		path := c.Request.URL.Path

		c.Next()

		latency := time.Since(start)
		status := c.Writer.Status()
		clientIP := c.ClientIP()
		method := c.Request.Method

		log.Printf("[GIN] %v | %3d | %13v | %15s | %-7s %s",
			start.Format("2025/08/12 - 15:04:05"),
			status,
			latency,
			clientIP,
			method,
			path,
		)
	}
}
