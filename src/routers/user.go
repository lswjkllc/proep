package routers

import (
	"github.com/labstack/echo"

	hs "github.com/lswjkllc/proep/src/handlers"
)

func AddUserRouter(e *echo.Echo) {
	e.POST("/user", hs.CreateUser)
	e.GET("/user/:id", hs.GetUser)
	e.PUT("/user/:id", hs.UpdateUser)

	e.GET("/html", hs.GetHtml)
}
