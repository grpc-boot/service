package model

// App 应用配置
type App struct {
	Name      string `json:"name" yaml:"name"`
	Addr      string `json:"addr" yaml:"addr"`
	AutoRegis bool   `json:"autoRegis" yaml:"autoRegis"`
}
