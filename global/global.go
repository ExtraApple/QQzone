package global

import (
	"github.com/minio/minio-go/v7"
	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
)

var (
	DB          *gorm.DB
	RDB         *redis.Client
	MinioClient *minio.Client
)
