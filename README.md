# service

### 代码分层

```text
- presentation              表现层
  - fasthttp      
  - gin
  - grpc
- application               应用层
- domain                    领域层
  - implement               领域聚合实现
    - v1                    版本
- dal                       数据访问层(db、redis等交互)
  - db                      数据持久化(db交互)
- common                    公共依赖(常量、变量、工具函数)
  - abstract                基础抽象(全局接口、基础接口)
  - components              基础组件
  - define                  常/变量定义
  - model                   简单对象
  - utils                   工具函数
- conf                      已发布配置(环境配置文件)
- scripts                   发布脚本(编译、安装)
```


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


