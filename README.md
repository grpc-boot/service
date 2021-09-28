# service

> 1.下载protoc

[https://github.com/protocolbuffers/protobuf/releases/](https://github.com/protocolbuffers/protobuf/releases/)

> 2.安装`google.golang.org/protobuf`

```bash
go install google.golang.org/protobuf/cmd/protoc-gen-go
go install google.golang.org/grpc/cmd/protoc-gen-go-grpc

go mod edit -replace=google.golang.org/grpc=github.com/grpc/grpc-go@latest
go mod tidy
```

> 3.生成代码
```bash
protoc --go_out=. --go_opt=paths=import --go-grpc_out=. --go-grpc_opt=paths=import proto/*.proto
```


