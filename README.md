# service

> 1.下载protoc

[https://github.com/protocolbuffers/protobuf/releases/](https://github.com/protocolbuffers/protobuf/releases/)

> 2.安装`google.golang.org/protobuf`

```bash
go mod edit -replace=google.golang.org/grpc=github.com/grpc/grpc-go@latest
go mod tidy
```

> 3.生成代码
```bash
protoc --go_out=plugins=grpc:. ./proto/*.proto
```


