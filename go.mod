module service

go 1.16

require (
	github.com/grpc-boot/base v1.0.20
	github.com/grpc-boot/betcd v0.0.0-20220114060045-aaa8b0110ff0
	github.com/grpc-boot/gateway v1.0.0
	github.com/grpc-boot/gedis v1.0.1
	github.com/grpc-boot/orm v1.0.1
)

replace google.golang.org/grpc => github.com/grpc/grpc-go v1.37.0
