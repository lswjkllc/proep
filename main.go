package main

import (
	"fmt"

	mn "github.com/lswjkllc/proep/src"
)

func main() {
	config := mn.GetContainer("./config/config.yaml")
	fmt.Printf("系统配置: %v\n", config)
}
