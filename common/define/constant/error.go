package constant

// 错误定义

// 通用错误
const (
	Success         = 10000
	ErrInner        = 990001
	ErrParam        = 990002
	ErrLimit        = 990003
	ErrNotAvailable = 990004
)

// 网关模块
const (
	ErrGatewayNotExists = 10001
	ErrGatewayTimeout   = 10002
)

// 用户模块
const (
	ErrPhoneExists = 20001
	ErrEmailExists = 20002
)
