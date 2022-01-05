package main

import (
	"fmt"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	sc "github.com/lswjkllc/proep/src"
	hs "github.com/lswjkllc/proep/src/handlers"
)

func AddUserRouter(e *echo.Echo) {
	e.POST("/user/create", hs.CreateUser)
}

func main() {
	config := sc.GetContainer("./config/config.yaml")
	fmt.Printf("系统配置: %v\n\n", config)

	e := echo.New()
	// 注册中间件
	e.Use(middleware.Recover())
	e.Use(middleware.Logger())
	// e.Use(login)
	// 添加路由
	AddUserRouter(e)
	// 添加监听器
	// 开启服务
	e.Logger.Fatal(e.Start(":1323"))
}
