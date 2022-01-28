package model

import (
	"github.com/grpc-boot/betcd"
)

// EtcdConf etcd配置
type EtcdConf struct {
	Name      string            `json:"name" yaml:"name"`
	Prefix    []string          `json:"prefix" yaml:"prefix"`
	KeyOption []betcd.KeyOption `json:"keyOption" yaml:"keyOption"`
}
