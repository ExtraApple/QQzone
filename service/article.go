package service

import (
	"QQZone/global"
	"QQZone/model"
	"QQZone/utils"
	"errors"
	"log"
	"mime/multipart"
	"strings"
	"time"

	"gorm.io/gorm"
)

// CreateArticle 创建文章 + 上传媒体
func CreateArticle(userID uint, title, content string, files []*multipart.FileHeader) (*model.Article, error) {
	article := model.Article{
		Title:   title,
		Content: content,
		UserID:  userID,
	}

	// 先保存文章（拿到 article.ID）
	if err := global.DB.Create(&article).Error; err != nil {
		return nil, err
	}

	// 处理上传文件
	var medias []model.Media
	for _, fileHeader := range files {
		file, err := fileHeader.Open()
		if err != nil {
			return nil, err
		}
		defer file.Close()

		// 生成对象名称
		objectName := utils.GenerateObjectName(userID, fileHeader.Filename)

		// 上传文件到 MinIO
		url, err := utils.UploadFile("articles", objectName, file, fileHeader.Size)
		if err != nil {
			return nil, err
		}

		// 确定文件类型
		fileType := "image"
		if strings.HasPrefix(fileHeader.Header.Get("Content-Type"), "video/") {
			fileType = "video"
		}

		medias = append(medias, model.Media{
			ArticleID: article.ID,
			Type:      fileType,
			URL:       url,
			ObjectKey: objectName,
			Size:      fileHeader.Size,
			Duration:  0, // 视频时长可用 ffmpeg 获取，这里先写 0
		})
	}

	// 保存 Media
	if len(medias) > 0 {
		if err := global.DB.Create(&medias).Error; err != nil {
			return nil, err
		}
	}

	// 加载媒体到 Article
	article.Media = medias

	return &article, nil
}

// ListArticles 获取所有文章列表，带媒体，按时间倒序
func ListArticles() ([]model.Article, error) {
	var articles []model.Article
	if err := global.DB.Preload("Media").Order("created_at desc").Find(&articles).Error; err != nil {
		return nil, err
	}
	return articles, nil
}

// GetArticle 获取单篇文章
func GetArticle(articleID uint) (*model.Article, error) {
	var article model.Article
	if err := global.DB.Preload("Media").First(&article, articleID).Error; err != nil {
		return nil, err
	}
	return &article, nil
}

// DeleteArticle 删除文章及其媒体（只允许作者自己删）
func DeleteArticle(articleID uint, userID uint) error {
	var article model.Article
	if err := global.DB.Preload("Media").First(&article, articleID).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return err
		}
		return err
	}

	if article.UserID != userID {
		return errors.New("permission denied")
	}

	for _, media := range article.Media {
		if err := utils.DeleteFile("articles", media.ObjectKey); err != nil {
			log.Printf("Failed to delete file %s: %v", media.ObjectKey, err)
		}
	}

	if err := global.DB.Delete(&article).Error; err != nil {
		return err
	}
	return nil
}

// ArticleResponse 对外返回结构（隐藏 DB 层细节）
type ArticleResponse struct {
	ID        uint          `json:"id"`
	Title     string        `json:"title"`
	Content   string        `json:"content"`
	UserID    uint          `json:"user_id"`
	Media     []model.Media `json:"media"`
	CreatedAt time.Time     `json:"created_at"`
	UpdatedAt time.Time     `json:"updated_at"`
}

// ToArticleResponse 将 model.Article 转为对外响应结构
func ToArticleResponse(a *model.Article) ArticleResponse {
	return ArticleResponse{
		ID:        a.ID,
		Title:     a.Title,
		Content:   a.Content,
		UserID:    a.UserID,
		Media:     a.Media,
		CreatedAt: a.CreatedAt,
		UpdatedAt: a.UpdatedAt,
	}
}

// ToArticleListResponse 批量转换
func ToArticleListResponse(articles []model.Article) []ArticleResponse {
	res := make([]ArticleResponse, 0, len(articles))
	for i := range articles {
		res = append(res, ToArticleResponse(&articles[i]))
	}
	return res
}
