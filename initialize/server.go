package initialize

import (
	"fmt"
	"os"

	"gopkg.in/yaml.v3"
)

type Config struct {
	Server struct {
		Port int `yaml:"port"`
	} `yaml:"server"`

	Mysql struct {
		Host     string `yaml:"host"`
		Port     int    `yaml:"port"`
		User     string `yaml:"user"`
		Password string `yaml:"password"`
		DB       string `yaml:"db"`
	} `yaml:"mysql"`
	Redis struct {
		Host     string `yaml:"host"`
		Port     int    `yaml:"port"`
		Password string `yaml:"password"`
		DB       int    `ymal:"db"`
	} `yaml:"redis"`
	Minio struct {
		Endpoint        string `yaml:"endpoint"`
		AccessKeyID     string `yaml:"accessKeyID"`
		SecretAccessKey string `yaml:"secretAccessKey"`
		UseSSL          bool   `yaml:"useSSl"`
		BucketName      string `yaml:"bucketName"`
		Location        string `yaml:"location"`
	} `yaml:"minio"`
}

func InitConfig() *Config {
	data, err := os.ReadFile("config.yaml")
	if err != nil {
		panic(fmt.Sprintf("读取配置文件失败: %v", err))
	}
	var conf Config
	if err := yaml.Unmarshal(data, &conf); err != nil {
		panic(fmt.Sprintf("解析配置文件失败: %v", err))
	}
	return &conf
}
