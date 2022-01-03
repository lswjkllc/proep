package main

import "fmt"

func main() {
	config := GetConfig()
	fmt.Printf("config type => %T\n", config)
	fmt.Printf("config value => %v\n", config)
}
