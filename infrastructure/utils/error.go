package utils

import "sync"

const (
	ErrPhoneExists = 10001
	ErrEmailExists = 10002
)

var errorMap = map[int32]string{
	ErrPhoneExists: "手机号已存在",
	ErrEmailExists: "邮箱已存在",
}

var errPool = sync.Pool{
	New: func() interface{} {
		return &Error{}
	},
}

// NewError 实例化错误对象
func NewError(code int32, msg string) *Error {
	e := errPool.Get().(*Error)
	e.code = code
	e.msg = msg
	return e
}

// NewErrorWithCode with code实例化错误对象
func NewErrorWithCode(code int32) *Error {
	msg, _ := errorMap[code]
	return NewError(code, msg)
}

// Error 通用错误对象
type Error struct {
	code int32
	msg  string
}

// Error 获取错误信息
func (e *Error) Error() string {
	return e.msg
}

// Code 获取错误码
func (e *Error) Code() int32 {
	return e.code
}

// Close 释放错误对象
func (e *Error) Close() {
	e.code = 0
	e.msg = ""
	errPool.Put(e)
}
