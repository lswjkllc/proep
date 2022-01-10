package routers

import (
	"github.com/labstack/echo"

	hs "github.com/lswjkllc/proep/src/handlers"
)

func AddUserRouter(e *echo.Echo) {
	e.POST("/user/create", hs.CreateUser)

	e.GET("/html", hs.GetHtml)
}
