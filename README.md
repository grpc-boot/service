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

### 示例表

```sql
CREATE TABLE `gateway` (
  `id` int unsigned NOT NULL AUTO_INCREMENT COMMENT '主键',
  `name` varchar(32) CHARACTER SET utf8 COLLATE utf8_general_ci DEFAULT '' COMMENT '名称',
  `path` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci DEFAULT '' COMMENT '路径',
  `second_limit` int unsigned DEFAULT '5000' COMMENT '每秒请求数',
  PRIMARY KEY (`id`),
  UNIQUE KEY `path` (`path`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;
```

```sql
CREATE TABLE `user` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `nickname` varchar(32) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '昵称',
  `passwd` char(32) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL DEFAULT '' COMMENT '登录密码',
  `phone` bigint unsigned DEFAULT NULL COMMENT '手机号',
  `email` varchar(64) CHARACTER SET utf8 COLLATE utf8_general_ci DEFAULT '' COMMENT '邮箱',
  `last_login_at` bigint DEFAULT NULL COMMENT '上次登录时间',
  `wx_uuid` varchar(64) CHARACTER SET utf8 COLLATE utf8_general_ci DEFAULT NULL COMMENT '微信uuid',
  `wx_pub_openid` varchar(64) CHARACTER SET utf8 COLLATE utf8_general_ci DEFAULT NULL COMMENT '微信公众号openid',
  `wx_sapp_openid` varchar(64) CHARACTER SET utf8 COLLATE utf8_general_ci DEFAULT NULL COMMENT '微信小程序openid',
  `created_at` bigint unsigned NOT NULL COMMENT '创建时间',
  `updated_at` bigint unsigned NOT NULL DEFAULT '0' COMMENT '更新时间',
  `is_on` tinyint unsigned NOT NULL DEFAULT '1' COMMENT '是否启用',
  PRIMARY KEY (`id`),
  UNIQUE KEY `phone` (`phone`),
  UNIQUE KEY `wx_uuid` (`wx_uuid`),
  UNIQUE KEY `wx_pub_openid` (`wx_pub_openid`),
  UNIQUE KEY `wx_sapp_openid` (`wx_sapp_openid`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;
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


