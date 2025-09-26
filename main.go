package main

import (
	"QQZone/initialize"
	"QQZone/router"
	"fmt"
	"log"
)

func main() {
	conf := initialize.InitConfig()
	fmt.Println("配置加载成功 ✅")
	//初始化MySQL
	initialize.InitMysql(conf)
	fmt.Println("数据库初始化 ✅")
	//初始化Redis
	initialize.InitRedis(conf)
	fmt.Println("Redis初始化 ✅")
	//初始化Minio
	initialize.InitMinio(conf)
	fmt.Println("Minio初始化 ✅")
	r := router.Router()
	if err := r.Run(":8080"); err != nil {
		log.Fatal(err)
	}
}
