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
				zap.String("traceId", traceId),
				zap.String("uri", c.Path()),
				zap.String("host", c.Request().Host),
				zap.String("remoteIp", c.RealIP()),
				zap.String("userAgent", c.Request().UserAgent()),
				zap.String("method", c.Request().Method),
				zap.Int("status", c.Response().Status),
				zap.Int64("bytesIn", c.Request().ContentLength),
				zap.Int64("bytesOut", c.Response().Size),
				zap.Time("startTime", startTime),
				zap.Duration("costTime", costTime),
			)
		}()
		return next(c)
	}
}
