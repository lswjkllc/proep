package utils

import (
	"time"

	"github.com/guonaihong/gout"
	"github.com/guonaihong/gout/dataflow"
)

type ReqMethod string

const (
	POST   ReqMethod = "post"
	GET    ReqMethod = "get"
	PUT    ReqMethod = "put"
	DELETE ReqMethod = "delete"
)

// 请求错误类
type ReqError struct {
	value string
}

func (e *ReqError) Error() string {
	return e.value
}

// 请求类
type Requests struct {
	Url     string                 `json:"url" yaml:"url"`
	Method  ReqMethod              `json:"method" yaml:"method"`
	Body    map[string]interface{} `json:"body" yaml:"body"`
	Headers map[string]string      `json:"headers" yaml:"headers"`
	Form    bool                   `json:"form" yaml:"form"`
	Json    bool                   `json:"json" yaml:"yaml"`
	Query   bool                   `json:"query" yaml:"query"`
	TraceId string                 `json:"tracdId" yaml:"traceId"`
}

func (req Requests) Do(resp interface{}) (time.Duration, error) {
	// 选择方法
	dataFlow := req.getDataFlowByMethod()
	if dataFlow == nil {
		return 0, &ReqError{"invalid request method"}
	}
	// 设置请求数据
	req.setData(dataFlow)
	// 设置头
	dataFlow.SetHeader(req.Headers)
	// 绑定数据
	dataFlow.BindJSON(resp)
	// 执行
	st := time.Now()
	err := dataFlow.Do()
	dt := time.Since(st)

	return dt, err
}

func (req Requests) setData(dataFlow *dataflow.DataFlow) {
	if req.Method == GET {
		dataFlow.SetQuery(req.Body)
	} else {
		if req.Json {
			dataFlow.SetJSON(req.Body)
		} else if req.Form {
			dataFlow.SetForm(req.Body)
		}
	}
}

func (req Requests) getDataFlowByMethod() *dataflow.DataFlow {
	// 定义 dataFlow
	var dataFlow *dataflow.DataFlow
	switch req.Method {
	case GET:
		dataFlow = gout.GET(req.Url)
	case POST:
		dataFlow = gout.POST(req.Url)
	case PUT:
		dataFlow = gout.PUT(req.Url)
	case DELETE:
		dataFlow = gout.DELETE(req.Url)
	}
	return dataFlow
}

func RequestUrl(url string, body map[string]interface{}) {

}
