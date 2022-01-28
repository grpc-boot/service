package constant

import "os"

// 日志
const (
	LoggerFlag = os.O_RDWR | os.O_CREATE | os.O_APPEND
	LoggerMode = 0666
)

// zap日志字段
const (
	FieldError    = `Error`
	FieldIp       = `Ip`
	FieldTraceId  = `TraceId`
	FieldEndpoint = `Endpoint`
)
