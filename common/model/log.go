package model

// Log 日志配置
type Log struct {
	Level     int8   `json:"level" yaml:"level"`
	ErrorPath string `json:"errorPath" yaml:"errorPath"`
	InfoPath  string `json:"infoPath"  yaml:"infoPath"`
}
