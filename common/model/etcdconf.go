package model

import "github.com/grpc-boot/betcd"

// EtcdConf etcd配置
type EtcdConf struct {
	Name   string             `json:"name" yaml:"name"`
	Option betcd.ConfigOption `json:"option" yaml:"option"`
}
