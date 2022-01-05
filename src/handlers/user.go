package handlers

import (
	"github.com/labstack/echo"
	sc "github.com/lswjkllc/proep/src"
	ms "github.com/lswjkllc/proep/src/models"
	us "github.com/lswjkllc/proep/src/utils"
)

func CreateUser(c echo.Context) error {
	// 绑定参数
	user := new(ms.User)
	if err := c.Bind(user); err != nil {
		return us.ResponseJson(c, us.Fail, err.Error(), nil)
	}
	// 获取 Container
	container := sc.GetContainer("")
	// 创建 User
	err := container.UserUsecase.CreateUser(user)
	if err != nil {
		return us.ResponseJson(c, us.Fail, err.Error(), nil)
	}
	// 返回响应
	return us.ResponseJson(c, us.Success, "", *user)
}
