package utils

import (
	"net/http"

	"github.com/labstack/echo"
)

type ResponseCode int

const (
	Fail    ResponseCode = 1
	Success ResponseCode = 0
)

func ResponseJson(c echo.Context, code ResponseCode, message string, data interface{}) error {

	result := make(map[string]interface{})
	result["code"] = code
	result["message"] = message
	result["data"] = data

	return c.JSON(http.StatusOK, result)
}
