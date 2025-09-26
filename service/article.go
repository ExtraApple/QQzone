package service

import (
	"QQZone/global"
	"QQZone/model"
	"QQZone/utils"
	"errors"
	"log"
	"mime/multipart"
	"strings"

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

// DeleteArticle 删除文章及其媒体（只允许作者自己删）
func DeleteArticle(articleID uint, userID uint) error {
	var article model.Article
	// 预加载 Media
	if err := global.DB.Preload("Media").First(&article, articleID).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return err
		}
		return err
	}

	// 权限校验：只能删除自己的文章
	if article.UserID != userID {
		return errors.New("permission denied")
	}

	// 删除 MinIO 文件
	for _, media := range article.Media {
		if err := utils.DeleteFile("articles", media.ObjectKey); err != nil {
			log.Printf("⚠️ Failed to delete file %s: %v", media.ObjectKey, err)
		}
	}

	// 删除文章（级联删除 Media）
	if err := global.DB.Delete(&article).Error; err != nil {
		return err
	}

	return nil
}
