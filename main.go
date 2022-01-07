package main

import (
	"fmt"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"

	sc "github.com/lswjkllc/proep/src"
	// coms "github.com/lswjkllc/proep/src/commons"
	hs "github.com/lswjkllc/proep/src/handlers"
)

func AddUserRouter(e *echo.Echo) {
	e.POST("/user/create", hs.CreateUser)

	e.GET("/html", hs.GetHtml)
}

func main() {
	// 获取 Container
	container := sc.GetContainer()
	fmt.Printf("系统配置: %v\n\n", container.BaseConfig)

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
