module github.com/grpc-boot/service

go 1.16

require (
	go.etcd.io/etcd/api/v3 v3.5.0-alpha.0
	go.etcd.io/etcd/client/v3 v3.5.0-alpha.0
	google.golang.org/grpc v1.37.0
	google.golang.org/protobuf v1.25.0
)

replace google.golang.org/grpc => github.com/grpc/grpc-go v1.37.0
