module service

go 1.16

require (
	github.com/gin-gonic/gin v1.7.7
	github.com/go-ozzo/ozzo-routing v2.1.4+incompatible // indirect
	github.com/golang/gddo v0.0.0-20210115222349-20d68f94ee1f // indirect
	github.com/grpc-boot/base v1.1.2
	github.com/grpc-boot/betcd v0.0.3
	github.com/grpc-boot/gateway v1.0.5
	github.com/grpc-boot/gedis v1.0.1
	github.com/grpc-boot/orm v1.1.0
	github.com/qiangxue/fasthttp-routing v0.0.0-20160225050629-6ccdc2a18d87
	github.com/valyala/fasthttp v1.33.0
	go.etcd.io/etcd/client/v3 v3.5.1
	go.uber.org/zap v1.20.0
	google.golang.org/grpc v1.44.0 // indirect
	google.golang.org/grpc/cmd/protoc-gen-go-grpc v1.2.0 // indirect
	google.golang.org/protobuf v1.27.1
)

replace google.golang.org/grpc => github.com/grpc/grpc-go v1.37.0
