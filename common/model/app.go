package model

// App 应用配置
type App struct {
	Name string `json:"name" yaml:"name"`
	Port uint16 `json:"port" yaml:"port"`
}
