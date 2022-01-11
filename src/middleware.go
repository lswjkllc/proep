package src

import (
	"time"

	"github.com/labstack/echo"
	"github.com/lswjkllc/proep/src/logger"
	"go.uber.org/zap"
)

// 请求结束中间件
func RequestEnd(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) (err error) {
		startTime := time.Now()
		defer func() {
			costTime := time.Since(startTime)
			traceId := c.Request().Header.Get("traceId")
			logger.Info(
				"RequestEnd",
				zap.String("traceId", traceId),                   // 请求 链路Id
				zap.String("uri", c.Path()),                      // 请求 Uri
				zap.String("host", c.Request().Host),             // 请求 Host
				zap.String("remoteIp", c.RealIP()),               // 请求 RealIP
				zap.String("userAgent", c.Request().UserAgent()), // 请求 用户代码
				zap.Int64("bytesIn", c.Request().ContentLength),  // 请求 数据长度
				zap.String("method", c.Request().Method),         // 请求 方式
				zap.Time("startTime", startTime),                 // 请求 开始时间
				zap.Duration("costTime", costTime),               // 请求/响应 耗时
				zap.Int("status", c.Response().Status),           // 响应 状态
				zap.Int64("bytesOut", c.Response().Size),         // 响应 数据状态
			)
		}()
		return next(c)
	}
}
