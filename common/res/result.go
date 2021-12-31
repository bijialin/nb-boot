package res

import (
	"github.com/bijialin/nb-boot/common/constant"
)

type Result struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

func Ok() *Result {
	return &Result{Code: 1,
		Msg: constant.API_OK_MSG}
}
func Error() *Result {
	return &Result{Code: -1,
		Msg: "请求失败"}
}

func Build(code int, msg string, data interface{}) *Result {
	if code == 0 {
		code = constant.API_OK_CODE
	}

	if msg == "" {
		msg = constant.API_OK_MSG
	}

	return &Result{
		Code: code,
		Msg:  msg,
		Data: data,
	}
}
