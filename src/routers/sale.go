package routers

import (
	"github.com/labstack/echo"

	hs "github.com/lswjkllc/proep/src/handlers"
)

func AddSaleRouter(e *echo.Echo) {
	e.GET("/sale/flash", hs.FlashSale)
}
