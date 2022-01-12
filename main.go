package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"

	sc "github.com/lswjkllc/proep/src"
	rs "github.com/lswjkllc/proep/src/routers"
)

func shutdown(c chan os.Signal, container *sc.Container) {
	for true {
		// 监听所有信号
		signal.Notify(c, syscall.SIGINT)
		<-c
		container.Close()

		os.Exit(0)
	}
}

func main() {
	// 信号通道
	c := make(chan os.Signal)

	// 初始化 echo 对象
	e := echo.New()

	// 获取 Container
	container := sc.GetContainer()
	fmt.Printf("系统配置: %+v\n", container.BaseConfig)
	go shutdown(c, container)

	// 注册计算时间中间件
	e.Use(sc.RequestEnd)
	// 注册错误恢复中间件
	e.Use(middleware.Recover())

	// 添加路由
	rs.AddUserRouter(e)

	// 开启服务
	e.Logger.Fatal(e.Start(container.BaseConfig.CommonBase.Addr))
}
