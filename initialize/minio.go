package initialize

import (
	"QQZone/global"
	"log"

	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
)

func InitMinio(conf *Config) {
	var err error
	global.MinioClient, err = minio.New(conf.Minio.Endpoint, &minio.Options{
		Creds:  credentials.NewStaticV4(conf.Minio.AccessKeyID, conf.Minio.SecretAccessKey, ""),
		Secure: conf.Minio.UseSSL,
	})
	if err != nil {
		log.Fatalf("InitMinio error:%v", err)
	}
	log.Println("MinIO 连接成功")
}
