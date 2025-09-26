package api

import (
	"QQZone/service"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// 文章创建
func ArticleCreate(c *gin.Context) {
	// 获取用户ID
	userID, exists := c.Get("userID")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
		return
	}
	uid := userID.(uint)

	// 解析表单字段
	title := c.PostForm("title")
	content := c.PostForm("content")
	if title == "" || content == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "title and content required"})
		return
	}

	// 获取上传文件
	form, err := c.MultipartForm()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid form data"})
		return
	}
	files := form.File["files"]

	// 调用 service
	article, err := service.CreateArticle(uid, title, content, files)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"msg":     "create article ok",
		"article": article,
	})
}

// 文章删除
func ArticleDelete(c *gin.Context) {
	// 获取用户ID
	userID, exists := c.Get("userID")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
		return
	}
	uid := userID.(uint)

	// 解析文章ID
	articleID, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid articleID"})
		return
	}

	// 调用 service 删除文章
	if err := service.DeleteArticle(uint(articleID), uid); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"msg": "delete article ok"})
}
