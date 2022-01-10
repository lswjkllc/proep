package main

import (
	"fmt"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"

	sc "github.com/lswjkllc/proep/src"
	rs "github.com/lswjkllc/proep/src/routers"
)

func main() {
	// 获取 Container
	container := sc.GetContainer()
	fmt.Printf("系统配置: %v\n\n", container.BaseConfig)

	// 初始化 echo 对象
	e := echo.New()

	// 注册计算时间中间件
	e.Use(sc.RequestEnd)
	// 注册错误恢复中间件
	e.Use(middleware.Recover())

	// 添加路由
	rs.AddUserRouter(e)

	// 开启服务
	e.Logger.Fatal(e.Start(":1323"))
}
