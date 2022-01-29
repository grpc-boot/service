package model

import "service/common/define/constant"

type Response struct {
	Code int32       `json:"code" yaml:"code" xml:"code"`
	Msg  string      `json:"msg" yaml:"msg" xml:"msg"`
	Data interface{} `json:"data" yaml:"data" xml:"data"`
}

func RespSuccess(data interface{}) Response {
	return Response{
		Code: constant.Success,
		Msg:  "ok",
		Data: data,
	}
}
