package db

import (
	"fmt"
	"time"

	coms "github.com/lswjkllc/proep/src/commons"
	ms "github.com/lswjkllc/proep/src/models"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func InitDB(migrate bool) *gorm.DB {
	// 获取 config
	config := coms.GetConfigByPath("../../config/config.yaml")
	// 获取 mysql dns
	mysqlDns := config.DataBase.MysqlData.GetDsn()
	// 获取 mysql 连接
	db, err := gorm.Open(mysql.Open(mysqlDns), &gorm.Config{})
	if err != nil {
		panic(fmt.Sprintf("创建数据库连接失败: %v", err))
	}
	// 自动迁移数据结构(table schema)
	if migrate {
		db.AutoMigrate(&ms.User{})
	}

	return db
}

func GetFirstRecord(db *gorm.DB) ms.User {
	// 获取最后一条记录
	var lastUser ms.User
	db.First(&lastUser)

	return lastUser
}

func GetLastRecord(db *gorm.DB) ms.User {
	// 获取最后一条记录
	var lastUser ms.User
	db.Last(&lastUser)

	return lastUser
}

func CreateUser(db *gorm.DB) ms.User {
	// 创造一个记录
	curTs := time.Now().UnixNano() / 1e6
	name := fmt.Sprintf("Name-%d", curTs)
	email := fmt.Sprintf("%d@yz.cn", curTs)
	passWord := fmt.Sprint(curTs)
	// 创建
	newUser := ms.User{Name: name, Gender: "男", Age: 18, Email: email, PassWord: passWord}
	db.Create(&newUser)
	fmt.Printf("新创建: %v\n", newUser)

	return newUser
}

func UpdateUser(db *gorm.DB) ms.User {
	// 获取最后一条记录
	lastUser := GetLastRecord(db)
	fmt.Printf("修改前: %v\n", lastUser)
	// 修改字段值
	name := fmt.Sprintf("%s-modify", lastUser.Name)
	email := fmt.Sprintf("modify-%s", lastUser.Email)
	passWord := fmt.Sprintf("%s-modify", lastUser.PassWord)
	// 修改
	db.Model(&lastUser).Updates(ms.User{Name: name, Age: lastUser.Age + 1, Email: email, PassWord: passWord})
	fmt.Printf("修改后: %v\n", lastUser)

	return lastUser
}

func DeleteUser(db *gorm.DB) ms.User {
	// 获取最后一条记录
	lastUser := GetLastRecord(db)
	fmt.Printf("删除前: %v\n", lastUser)
	// 删除记录
	db.Delete(&lastUser)
	fmt.Printf("删除后: %v\n", lastUser)
	return lastUser
}
