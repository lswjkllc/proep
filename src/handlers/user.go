package handlers

import (
	"fmt"
	"strconv"

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
	// 创建 User
	err := sc.GetContainer().UserUsecase.CreateUser(user)
	if err != nil {
		return us.ResponseJson(c, us.Fail, err.Error(), nil)
	}

	return us.ResponseJson(c, us.Success, "", *user)
}

func GetUser(c echo.Context) error {
	// 获取 path 参数
	sid := c.Param("id")
	// str to int
	id, err := strconv.Atoi(sid)
	if err != nil {
		return us.ResponseJson(c, us.Fail, err.Error(), nil)
	}
	// 获取 User
	user, err := sc.GetContainer().UserUsecase.GetUserById(id)
	if err != nil {
		return us.ResponseJson(c, us.Fail, err.Error(), nil)
	}

	return us.ResponseJson(c, us.Success, "", user)
}

func UpdateUser(c echo.Context) error {
	// 获取 path 参数
	sid := c.Param("id")
	// str to int
	id, err := strconv.Atoi(sid)
	if err != nil {
		return us.ResponseJson(c, us.Fail, err.Error(), nil)
	}
	// 获取 Container
	container := sc.GetContainer()
	// 获取 User
	user, err := container.UserUsecase.GetUserById(id)
	fmt.Printf("old user: %+v\n", user)
	if err != nil {
		return us.ResponseJson(c, us.Fail, err.Error(), nil)
	}
	// 绑定参数
	if err := c.Bind(&user); err != nil {
		return us.ResponseJson(c, us.Fail, err.Error(), nil)
	}
	// 保存 user
	container.UserUsecase.Save(&user)
	fmt.Printf("new user: %+v\n", user)

	return us.ResponseJson(c, us.Success, "", user)
}

func DeleteUser(c echo.Context) error {
	// 获取 path 参数
	sid := c.Param("id")
	// str to int
	id, err := strconv.Atoi(sid)
	if err != nil {
		return us.ResponseJson(c, us.Fail, err.Error(), nil)
	}
	// 删除
	err = sc.GetContainer().UserUsecase.DeleteUserById(id, false)
	if err != nil {
		return us.ResponseJson(c, us.Fail, err.Error(), nil)
	}
	return us.ResponseJson(c, us.Success, "", nil)
}
