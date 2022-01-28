package define

import "os"

// 容器内置配置
const (
	Db         = `db`
	Redis      = `redis`
	Etcd       = `etcd`
	EtcdConf   = `etcdConf`
	EtcdNaming = `etcdNaming`
)

// 容器用户自定义配置
const (
	UserDb    = `user-db`
	UserRedis = `user-redis`
)

// 默认配置文件
const (
	JsonConfig = `conf/app.json`
	YamlConfig = `conf/app.yml`
)

// 错误定义
const (
	ErrPhoneExists = 10001
	ErrEmailExists = 10002
)

// 日志
const (
	LoggerFlag = os.O_RDWR | os.O_CREATE | os.O_APPEND
	LoggerMode = 0666
)
