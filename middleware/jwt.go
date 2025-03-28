package middleware

import (
	"Progect/utils"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

func JWTAuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenHeader := c.Request.Header.Get("Authorization")
		if tokenHeader == "" || !strings.HasPrefix(tokenHeader, "Bearer ") {
			c.JSON(http.StatusUnauthorized, gin.H{
				"err": "不存在身份信息",
			})
			c.Abort()
			//表示终止不执行接下来的组建了
		}
		token := strings.TrimPrefix(tokenHeader, "Bearer ")
		claims, err := utils.ParseToken(token)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{
				"err": "TOKEN 时间已经失效",
			})
			c.Abort()
			return
		}
		c.Set("username", claims.Username)
		c.Next()
	}
}
