package define

// 容器配置
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
