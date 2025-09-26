package router

import (
	"QQZone/api"
	"QQZone/middleware"

	"github.com/gin-gonic/gin"
)

func Router() *gin.Engine {
	r := gin.Default()
	user := r.Group("/user")
	{
		user.POST("/register", api.Register)
		user.POST("/login", api.Login)
		user.POST("/friends/add/:id", middleware.AuthMiddleware(), api.ADDFriend)
		user.DELETE("/friends/delete/:id", middleware.AuthMiddleware(), api.Deletefriend)
		user.GET("/friends", middleware.AuthMiddleware(), api.ListFriend)
		// 受保护路由
		user.GET("/admin", middleware.AuthMiddleware(), middleware.AdminOnly(), api.AdminOnly)
	}
	articles := r.Group("/articles")
	{
		articles.POST("/create", middleware.AuthMiddleware(), api.ArticleCreate)
		articles.DELETE(":id", middleware.AuthMiddleware(), api.ArticleDelete)
		articles.POST(":id/comments", middleware.AuthMiddleware(), api.CommentCreate)
		articles.GET(":id/comments", api.GetComments)
		articles.DELETE("comments/:id", middleware.AuthMiddleware(), api.CommentDelete)
	}

	return r
}
