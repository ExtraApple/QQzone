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
	initialize.InitMysql(conf)
	fmt.Println("数据库初始化 ✅")
	initialize.InitRedis(conf)
	fmt.Println("Redis初始化 ✅")
	initialize.InitMinio(conf)
	fmt.Println("Minio初始化 ✅")
	r := router.Router()
	addr := fmt.Sprintf("0.0.0.0:%d", conf.Server.Port)
	fmt.Printf("服务启动于 http://localhost:%d\n", conf.Server.Port)
	if err := r.Run(addr); err != nil {
		log.Fatal(err)
	}
}
