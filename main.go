package main

import (
	"fmt"

	mn "github.com/lswjkllc/proep/src"
)

func main() {
	config := mn.GetContainer()
	fmt.Printf("系统配置: %v\n", config)
}
