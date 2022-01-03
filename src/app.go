package main

import "fmt"

func main() {
	container := GetContainer()
	fmt.Printf("config type => %T\n", container)
	fmt.Printf("config value => %v\n", container)
}
