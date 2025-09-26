package initialize

import (
	"QQZone/global"
	"context"
	"fmt"

	"github.com/redis/go-redis/v9"
)

var Ctx = context.Background()

func InitRedis(conf *Config) {
	global.RDB = redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%d", conf.Redis.Host, conf.Redis.Port),
		Password: conf.Redis.Password,
		DB:       conf.Redis.DB,
	})

	if err := global.RDB.Ping(context.Background()).Err(); err != nil {
		panic(fmt.Sprintf("Redis 连接失败: %v", err))
	}
}
