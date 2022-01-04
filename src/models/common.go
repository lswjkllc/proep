package models

import (
	"fmt"

	coms "github.com/lswjkllc/proep/src/commons"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func InitDB(config *coms.ConfigInfo, migrate bool) *gorm.DB {
	// 获取 mysql dns
	mysqlDns := config.DataBase.MysqlData.GetDns()
	// 获取 mysql 连接
	db, err := gorm.Open(mysql.Open(mysqlDns), &gorm.Config{})
	if err != nil {
		panic(fmt.Sprintf("创建数据库连接失败: %v", err))
	}
	// 自动迁移数据结构(table schema)
	if migrate {
		db.AutoMigrate(&User{})
	}

	return db
}
