package model

import (
	"github.com/grpc-boot/gedis"
	"github.com/grpc-boot/orm"
	clientv3 "go.etcd.io/etcd/client/v3"
)

// Conf 配置
type Conf struct {
	App   App             `json:"app" yaml:"app"`
	Log   Log             `json:"log" yaml:"log"`
	Db    orm.GroupOption `json:"db" yaml:"db"`
	Redis gedis.Option    `json:"redis" yaml:"redis"`
	Etcd  clientv3.Config `json:"etcd" yaml:"etcd"`
}
