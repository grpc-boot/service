package model

// EtcdNaming etcdNaming配置
type EtcdNaming struct {
	Name    string `json:"name" yaml:"name"`
	Service string `json:"service" yaml:"service"`
}
