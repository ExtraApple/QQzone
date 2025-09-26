package middleware

import (
	"QQZone/global"
	"QQZone/utils"
	"context"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Authorization header missing"})
			c.Abort()
			return
		}

		parts := strings.SplitN(authHeader, " ", 2)
		if len(parts) != 2 || parts[0] != "Bearer" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "invalid Authorization header"})
			c.Abort()
			return
		}
		tknStr := parts[1]

		// 1. 校验 JWT
		claims, err := utils.ParseToken(tknStr)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "token invalid or expired"})
			c.Abort()
			return
		}

		// 2. Redis 黑名单检查
		if global.RDB != nil {
			ctx := context.Background()
			val, _ := global.RDB.Get(ctx, "blacklist:"+tknStr).Result()
			if val == "true" {
				c.JSON(http.StatusUnauthorized, gin.H{"error": "token revoked"})
				c.Abort()
				return
			}
		}

		// 3. 保存 claims
		c.Set("user", claims)

		// 额外保存 userID，方便 API 层直接使用
		c.Set("userID", claims.UserID)

		c.Next()
	}
}
