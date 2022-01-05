package main

import (
	"fmt"

	mn "github.com/lswjkllc/proep/src"
)

func main() {
	config := mn.GetContainer("./config/config.yaml")
	fmt.Printf("系统配置: %v\n\n", config)

	// 创建 user
	data := map[string]interface{}{
		"name": "123", "age": 20, "gender": "女", "email": "xxx@mu.top", "password": "xx-123"}
	user, _ := config.UserUsecase.CreateUser(data)
	fmt.Printf("Create User: %v\n", user)
}
