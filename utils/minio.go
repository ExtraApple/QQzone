package utils

import (
	"QQZone/global"
	"context"
	"fmt"
	"io"
	"mime/multipart"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/minio/minio-go/v7"
)

// 上传文件
func UploadFile(bucketName, objectName string, file multipart.File, fileSize int64) (string, error) {
	//检查Minio是否初始化
	if global.MinioClient == nil {
		return "", fmt.Errorf("MinIO client not initialized")
	}
	//检查是否创建对应的Bucket
	exists, err := global.MinioClient.BucketExists(context.Background(), bucketName)
	if err != nil {
		return "", fmt.Errorf("chect bucket exists error:%v", err)
	}
	if !exists {
		err := global.MinioClient.MakeBucket(context.Background(), bucketName, minio.MakeBucketOptions{})
		if err != nil {
			return "", fmt.Errorf("create bucket error:%v", err)
		}
	}
	//设置对象元数据
	contentType := "application/octet-stream"
	ext := strings.ToLower(filepath.Ext(objectName)) //提取文件扩展名并转为小写
	switch ext {
	case ".jpg", ".jpeg":
		contentType = "image/jpeg"
	case ".png":
		contentType = "image/png"
	case ".gif":
		contentType = "image/gif"
	case ".mp4":
		contentType = "video/mp4"
	case ".mov":
		contentType = "video/quicktime"
	}
	//上传文件
	_, err = global.MinioClient.PutObject(context.Background(), bucketName, objectName, file, fileSize, minio.PutObjectOptions{
		ContentType: contentType,
	})
	if err != nil {
		return "", fmt.Errorf("upload file error: %v", err)
	}

	// 返回文件URL
	return fmt.Sprintf("%s/%s/%s", global.MinioClient.EndpointURL().String(), bucketName, objectName), nil
}

// 从MinIO删除文件
func DeleteFile(bucketName, objectName string) error {
	//检查Minio是否初始化
	if global.MinioClient == nil {
		return fmt.Errorf("MinIO client not initialized")
	}
	err := global.MinioClient.RemoveObject(context.Background(), bucketName, objectName, minio.RemoveObjectOptions{})
	if err != nil {
		return fmt.Errorf("delete file error:%v", err)
	}
	return nil
}

// GenerateObjectName 生成唯一的对象名称
func GenerateObjectName(userID uint, fileName string) string {
	ext := filepath.Ext(fileName)
	return fmt.Sprintf("articles/%d/%d_%s%s", userID, time.Now().UnixNano(), RandomString(8), ext)
}

// RandomString 生成随机字符串
func RandomString(length int) string {
	const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	b := make([]byte, length)
	for i := range b {
		b[i] = charset[time.Now().UnixNano()%int64(len(charset))]
	}
	return string(b)
}

// CreateTempFile 创建临时文件
func CreateTempFile(fileHeader *multipart.FileHeader) (*os.File, error) {
	file, err := fileHeader.Open()
	if err != nil {
		return nil, err
	}
	defer file.Close()

	// 创建临时文件
	tmpFile, err := os.CreateTemp("", "upload-*.tmp")
	if err != nil {
		return nil, err
	}
	// 复制文件内容
	if _, err := io.Copy(tmpFile, file); err != nil {
		tmpFile.Close()
		return nil, err
	}

	return tmpFile, nil
}
