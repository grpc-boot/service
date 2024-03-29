package constant

import "os"

// logger 相关

// 日志
const (
	LoggerFlag = os.O_RDWR | os.O_CREATE | os.O_APPEND
	LoggerMode = 0666
)

// zap日志字段
const (
	ZapPath     = `Path`
	ZapError    = `Error`
	ZapInfo     = `Info`
	ZapIp       = `Ip`
	ZapTraceId  = `TraceId`
	ZapEndpoint = `Endpoint`
)
