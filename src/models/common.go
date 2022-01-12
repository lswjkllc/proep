package models

import (
	"fmt"

	coms "github.com/lswjkllc/proep/src/commons"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

func InitDB(mysqlConfig *coms.MysqlDataEntity, migrate bool) *gorm.DB {
	// 获取 mysql dns
	mysqlDns := mysqlConfig.GetDsn()
	// 获取 mysql 连接
	db, err := gorm.Open(
		mysql.Open(mysqlDns),
		&gorm.Config{
			NamingStrategy: schema.NamingStrategy{SingularTable: true}})
	if err != nil {
		panic(fmt.Sprintf("创建数据库连接失败: %v", err))
	}
	// 自动迁移数据结构(table schema)
	if migrate {
		db.AutoMigrate(&User{})
	}

	return db
}

func InitCache(config *coms.RedisDataEntity) string {
	return ""
}
