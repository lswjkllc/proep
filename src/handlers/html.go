package handlers

import (
	"github.com/labstack/echo"
	us "github.com/lswjkllc/proep/src/utils"
)

type Html struct {
	Code    uint        `json:"code" yaml:"code"`
	Message string      `json:"message" yaml:"message"`
	Data    interface{} `json:"data" yaml:"data"`
}

type ProfessionSearchData struct {
	Professions []interface{} `json:"professions" yaml:"professions"`
	Total       int32         `json:"total" yaml:"total"`
	Limit       uint          `json:"limit" yaml:"limit"`
	Offset      uint          `json:"offset" yaml:"offset"`
}

func GetHtml(c echo.Context) error {
	html := Html{}
	// 组织请求
	req := us.Requests{
		Url:    "http://127.0.0.1:8001/profession/search",
		Method: us.POST,
		Body: map[string]interface{}{
			"origin": "b9afeacd09885ee3bea033ffb86563ae", "ownerCode": "241"},
		Headers: nil,
		Json:    true,
	}
	// 获取结果
	dt, err := req.Do(&html)
	c.Logger().Printf("cost time: %s", dt)
	if err != nil {
		return us.ResponseJson(c, us.Fail, err.Error(), nil)
	}
	// 错误判断
	if html.Code != 0 {
		return us.ResponseJson(c, us.Fail, html.Message, nil)
	}
	// 返回响应
	return us.ResponseJson(c, us.Success, "", html.Data)
}
