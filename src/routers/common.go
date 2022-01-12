package routers

import (
	"net/http"

	"github.com/labstack/echo"
)

func AddHealthRouter(e *echo.Echo) {
	e.GET("/health", func(c echo.Context) error {
		return c.String(http.StatusOK, "ok")
	})
}
