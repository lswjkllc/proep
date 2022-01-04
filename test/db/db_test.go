package db

import (
	"fmt"
	"testing"

	"gorm.io/gorm"
)

// 	获取连接
var db *gorm.DB = InitDB(false)

func TestUser(t *testing.T) {
	// 测试创建
	fmt.Printf("====== 开始创建:\n")
	t.Log(CreateUser(db))
	// 测试修改
	fmt.Printf("\n====== 开始修改:\n")
	t.Log(UpdateUser(db))
	// 测试删除
	fmt.Printf("\n====== 开始删除:\n")
	t.Log(DeleteUser(db))
	// // 查看插入后的全部元素
	// var users []ms.User
	// db.Find(&users)
	// fmt.Printf("All record: %v\n\n", users)
}
