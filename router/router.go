package router

import (
	"QQZone/api"
	"QQZone/middleware"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func Router() *gin.Engine {
	r := gin.Default()

	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:5173", "http://127.0.0.1:5173"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	user := r.Group("/user")
	{
		user.POST("/register", api.Register)
		user.POST("/login", api.Login)
		user.DELETE("/logout", middleware.AuthMiddleware(), api.Logout)
		user.POST("/friends/add/:id", middleware.AuthMiddleware(), api.ADDFriend)
		user.DELETE("/friends/delete/:id", middleware.AuthMiddleware(), api.Deletefriend)
		user.GET("/friends", middleware.AuthMiddleware(), api.ListFriend)
		user.GET("/admin", middleware.AuthMiddleware(), middleware.AdminOnly(), api.AdminOnly)
	}
	articles := r.Group("/articles")
	{
		articles.GET("", api.ArticleList)
		articles.GET("/:id", api.ArticleDetail)
		articles.POST("/create", middleware.AuthMiddleware(), api.ArticleCreate)
		articles.DELETE("/:id", middleware.AuthMiddleware(), api.ArticleDelete)
		articles.POST("/:id/comments", middleware.AuthMiddleware(), api.CommentCreate)
		articles.GET("/:id/comments", api.GetComments)
		articles.DELETE("/comments/:id", middleware.AuthMiddleware(), api.CommentDelete)
	}

	return r
}
