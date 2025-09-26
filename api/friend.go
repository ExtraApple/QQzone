package api

import (
	"QQZone/service"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// 添加好友
func ADDFriend(c *gin.Context) {
	// 从认证中间件获取当前用户ID
	userIDValue, exists := c.Get("userID")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
		return
	}

	// 将userID转换为uint类型
	userID, ok := userIDValue.(uint)
	if !ok {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "invalid user id type"})
		return
	}

	// 从URL参数获取好友ID
	friendIDStr := c.Param("id")
	friendID, err := strconv.ParseUint(friendIDStr, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid friend id"})
		return
	}

	// 调用服务层添加好友
	err = service.AddFriend(userID, uint(friendID))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "friend added successfully"})
}

// 删除好友
func Deletefriend(c *gin.Context) {
	// 从认证中间件获取当前用户ID
	userIDValue, exists := c.Get("userID")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
		return
	}

	// 将userID转换为uint类型
	userID, ok := userIDValue.(uint)
	if !ok {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "invalid user id type"})
		return
	}

	// 从URL参数获取好友ID
	friendIDStr := c.Param("id")
	friendID, err := strconv.ParseUint(friendIDStr, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid friend id"})
		return
	}

	// 调用服务层删除好友
	err = service.DeleteFriend(userID, uint(friendID))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "friend deleted successfully"})
}

// 获取好友列表
func ListFriend(c *gin.Context) {
	// 从认证中间件获取当前用户ID
	userIDValue, exists := c.Get("userID")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
		return
	}

	// 将userID转换为uint类型
	userID, ok := userIDValue.(uint)
	if !ok {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "invalid user id type"})
		return
	}

	friends, err := service.ListFriend(userID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"friends": friends,
	})
}
