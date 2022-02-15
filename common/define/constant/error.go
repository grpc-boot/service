package constant

import "google.golang.org/grpc/codes"

// 错误定义

// 通用错误
const (
	Success            = int32(0)
	ErrInner           = int32(codes.Internal)
	ErrParam           = int32(codes.InvalidArgument)
	ErrLimit           = int32(codes.PermissionDenied)
	ErrNotAvailable    = int32(codes.Unavailable)
	ErrUnauthenticated = int32(codes.Unauthenticated)
)

// 网关模块
const (
	ErrGatewayNotExists int32 = 10001
	ErrGatewayTimeout   int32 = 10002
)

// 用户模块
const (
	ErrPhoneExists int32 = 20001
	ErrEmailExists int32 = 20002
)
