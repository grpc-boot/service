package components

import (
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"sync"

	"service/common/define/constant"
	"service/common/model"

	"github.com/grpc-boot/base"
)

var errorMap = map[uint32]string{
	constant.Success:            "操作成功",
	constant.ErrInner:           "出现了点小状况，请稍后再试",
	constant.ErrParam:           "参数错误",
	constant.ErrLimit:           "请求过于频繁，请稍候重试",
	constant.ErrNotAvailable:    "服务暂时不可用，请稍候重试",
	constant.ErrUnauthenticated: "未认证请求",

	constant.ErrGatewayNotExists: "网关不存在",
	constant.ErrGatewayTimeout:   "请求超时",

	constant.ErrPhoneExists: "手机号已存在",
	constant.ErrEmailExists: "邮箱已存在",
}

var errPool = sync.Pool{
	New: func() interface{} {
		return &Error{}
	},
}

// NewError 实例化错误对象
func NewError(code uint32, msg string) *Error {
	e := errPool.Get().(*Error)
	e.code = code
	e.msg = msg
	return e
}

// NewErrorWithCode with code实例化错误对象
func NewErrorWithCode(code uint32) *Error {
	msg, _ := errorMap[code]
	return NewError(code, msg)
}

// Code2Response 状态码转换为response
func Code2Response(code uint32) model.Response {
	msg, _ := errorMap[code]
	return model.Response{
		Code: code,
		Msg:  msg,
		Data: base.SetValue,
	}
}

// Code2Error 状态码转为status
func Code2Error(code uint32) error {
	msg, _ := errorMap[code]
	return status.Errorf(codes.Code(code), msg)
}

// Error 通用错误对象
type Error struct {
	code uint32
	msg  string
}

// Error 获取错误信息
func (e *Error) Error() string {
	return e.msg
}

// Code 获取错误码
func (e *Error) Code() uint32 {
	return e.code
}

// ToResponse 转换为response
func (e *Error) ToResponse() model.Response {
	return model.Response{
		Code: e.code,
		Msg:  e.msg,
		Data: base.SetValue,
	}
}

// Close 释放错误对象
func (e *Error) Close() {
	e.code = 0
	e.msg = ""
	errPool.Put(e)
}
