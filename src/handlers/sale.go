package handlers

import (
	"context"
	"fmt"

	"github.com/go-redis/redis/v8"
	"github.com/labstack/echo"
	"go.uber.org/zap"

	sc "github.com/lswjkllc/proep/src"
	"github.com/lswjkllc/proep/src/logger"
	us "github.com/lswjkllc/proep/src/utils"
)

func flashSale(tx *redis.Tx) error {
	cxt := context.Background()

	// 先查询下当前watch监听的key的值
	v, err := tx.Get(cxt, "sales").Result()
	if err != nil && err != redis.Nil {
		return err
	}

	// 这里可以处理业务
	fmt.Println(v)

	// 如果key的值没有改变的话，Pipelined函数才会调用成功
	_, err = tx.Pipelined(cxt, func(pipe redis.Pipeliner) error {
		// 在这里给key设置最新值
		pipe.Set(cxt, "sales", "new value", 0)
		return nil
	})
	return err
}

func FlashSale(c echo.Context) error {
	flashCount := 10
	logger.Info(c, "", zap.Int("flashCount", flashCount))

	// container
	container := sc.GetContainer()

	// 秒杀
	err := container.GoodsUsecase.FlashSale(flashCount, "sales")
	if err != nil {
		logger.Error(c, err.Error())
		return us.ResponseJson(c, us.Fail, err.Error(), nil)
	}

	return us.ResponseJson(c, us.Success, "请求成功", nil)
}
