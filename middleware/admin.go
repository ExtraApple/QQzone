package middleware

import (
	"QQZone/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

// 只允许 admin 角色访问
func AdminOnly() gin.HandlerFunc {
	return func(c *gin.Context) {
		u, exists := c.Get("user")
		if !exists {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "user not found"})
			c.Abort()
			return
		}

		claims, ok := u.(*utils.CustomClaims)
		if !ok {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "claims type error"})
			c.Abort()
			return
		}

		if claims.Role != "admin" {
			c.JSON(http.StatusForbidden, gin.H{"error": "admin only"})
			c.Abort()
			return
		}

		c.Next()
	}
}
