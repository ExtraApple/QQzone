package api

import (
	"QQZone/global"
	"QQZone/initialize"
	"QQZone/service"
	"QQZone/utils"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type ReqRegister struct {
	Username string `json:"username" binding:"required,min=3,max=64"`
	Password string `json:"password" binding:"required,min=6"`
}

type ReqLogin struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func Register(c *gin.Context) {
	var req ReqRegister
	if err := c.ShouldBind(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	user, err := service.RegisterUser(req.Username, req.Password)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"user": user,
	})
}

func Login(c *gin.Context) {
	var req ReqLogin
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user, err := service.Authenticate(req.Username, req.Password)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	token, err := utils.GenerateToken(user.ID, user.Username, user.Role)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "token generation failed"})
		return
	}

	// Redis 存储 token
	if global.RDB == nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "redis not initialized"})
		return
	}

	err = global.RDB.Set(initialize.Ctx, utils.RedisSessionKey(token), user.Username, 24*time.Hour).Err()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to store session in redis", "detail": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"token": token})
}

func AdminOnly(c *gin.Context) {
	// middleware 已经验证并放入 context
	u, _ := c.Get("user")
	c.JSON(http.StatusOK, gin.H{
		"msg":  "hello admin",
		"user": u,
	})
}
