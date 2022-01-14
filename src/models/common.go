package models

import (
	"fmt"
	"time"

	"github.com/go-redis/redis/v8"
	coms "github.com/lswjkllc/proep/src/commons"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
)

func InitDB(mysqlConfig *coms.MysqlDataEntity, debug bool) *gorm.DB {
	// 获取 mysql dns
	mysqlDns := mysqlConfig.GetDsn()
	// 获取 gorm Config
	gormConfig := gorm.Config{NamingStrategy: schema.NamingStrategy{SingularTable: true}}
	if debug {
		// debug 模式下, 打开 sql 日志
		gormConfig.Logger = logger.Default.LogMode(logger.Info)
	}
	// 获取 mysql 连接
	db, err := gorm.Open(mysql.Open(mysqlDns), &gormConfig)
	if err != nil {
		panic(fmt.Sprintf("创建数据库连接失败: %v", err))
	}
	// 自动迁移数据结构(table schema)
	if debug {
		// debug 模式下, 迁移数据结构
		db.AutoMigrate(&User{})
	}

	return db
}

func InitCache(config *coms.RedisDataEntity) *redis.Client {
	client := redis.NewClient(&redis.Options{
		Addr:         config.Addr,
		Password:     config.Password,
		DB:           config.Name,
		PoolSize:     config.PoolSize,
		MinIdleConns: config.MinIdle,
		PoolTimeout:  time.Duration(config.Timeout),
	})
	return client
}
