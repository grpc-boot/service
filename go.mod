module github.com/grpc-boot/service

go 1.16

require (
	google.golang.org/grpc v1.37.0
	google.golang.org/protobuf v1.25.0
)

replace google.golang.org/grpc => github.com/grpc/grpc-go v1.37.0
