package model

// App 应用配置
type App struct {
	Name      string `json:"name" yaml:"name"`
	Addr      string `json:"addr" yaml:"addr"`
	AutoRegis bool   `json:"autoRegis" yaml:"autoRegis"`
	Location  string `json:"location" yaml:"location"`
}

// Build 初始化
func (a *App) Build() {
	if a.Location == "" {
		a.Location = "Asia/Shanghai"
	}
}
