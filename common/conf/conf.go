package conf

import (
	"github.com/grpc-boot/base"
	_ "github.com/grpc-boot/betcd"
	_ "github.com/grpc-boot/gateway"
	"github.com/grpc-boot/gedis"
	"github.com/grpc-boot/orm"
	clientv3 "go.etcd.io/etcd/client/v3"
)

var (
	conf Conf
)

// Conf 配置
type Conf struct {
	Db    orm.GroupOption     `json:"db" yaml:"db"`
	Redis []gedis.GroupOption `json:"redis" yaml:"redis"`
	Etcd  clientv3.Config     `json:"etcd" yaml:"etcd"`
}

func InitJsonConf(filename string) error {
	return base.JsonDecodeFile(filename, &conf)
}

func InitYamlConf(filename string) error {
	return base.YamlDecodeFile(filename, &conf)
}

func GetConf() *Conf {
	return &conf
}
