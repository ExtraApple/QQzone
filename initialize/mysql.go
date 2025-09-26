package initialize

import (
	"QQZone/global"
	"QQZone/model"
	"fmt"
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func InitMysql(conf *Config) {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		conf.Mysql.User,
		conf.Mysql.Password,
		conf.Mysql.Host,
		conf.Mysql.Port,
		conf.Mysql.DB,
	)

	var err error
	global.DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{}) // ⚡ 改这里
	if err != nil {
		log.Fatalf("数据库连接失败: %v", err)
	}

	// 自动迁移表数据
	if err := global.DB.AutoMigrate(&model.User{}, &model.Article{}, &model.Media{}, &model.Comment{}, &model.UserFriend{}); err != nil {
		log.Fatalf("数据表迁移失败:%v", err)
	}
}
