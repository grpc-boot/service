package conf

import (
	"github.com/grpc-boot/gedis"
	"github.com/grpc-boot/orm"
	clientv3 "go.etcd.io/etcd/client/v3"
)

// Conf 配置
type Conf struct {
	App   App                 `json:"app" yaml:"app"`
	Db    orm.GroupOption     `json:"db" yaml:"db"`
	Redis []gedis.GroupOption `json:"redis" yaml:"redis"`
	Etcd  clientv3.Config     `json:"etcd" yaml:"etcd"`
}

// App 应用配置
type App struct {
	Name string `json:"name" yaml:"name"`
	Port uint16 `json:"port" yaml:"port"`
}
