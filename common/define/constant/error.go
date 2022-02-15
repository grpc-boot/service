package constant

import "google.golang.org/grpc/codes"

// 错误定义

// 通用错误
const (
	Success            = 0
	ErrInner           = uint32(codes.Internal)
	ErrParam           = uint32(codes.InvalidArgument)
	ErrLimit           = uint32(codes.PermissionDenied)
	ErrNotAvailable    = uint32(codes.Unavailable)
	ErrUnauthenticated = uint32(codes.Unauthenticated)
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
