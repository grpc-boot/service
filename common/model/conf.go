package model

import (
	"github.com/grpc-boot/gedis"
	"github.com/grpc-boot/orm"
)

// Conf 配置
type Conf struct {
	App        App                    `json:"app" yaml:"app"`
	Id         SnowFlake              `json:"id" yaml:"id"`
	Log        Log                    `json:"log" yaml:"log"`
	Db         orm.GroupOption        `json:"db" yaml:"db"`
	Redis      gedis.Option           `json:"redis" yaml:"redis"`
	Etcd       Etcd                   `json:"etcd" yaml:"etcd"`
	EtcdConf   EtcdConf               `json:"etcdConf" yaml:"etcdConf"`
	EtcdNaming EtcdNaming             `json:"etcdNaming" yaml:"etcdNaming"`
	Var        map[string]interface{} `json:"var" yaml:"var"`
}

func (c *Conf) Float64Var(key string) (val float64, ok bool) {
	value, exists := c.Var[key]
	if !exists {
		return
	}

	switch v := value.(type) {
	case float64:
		return v, true
	case int:
		return float64(v), true
	case int64:
		return float64(v), true
	}
	return
}

// Int64Var 获取int64类型变量
func (c *Conf) Int64Var(key string) (val int64, ok bool) {
	value, exists := c.Var[key]
	if !exists {
		return
	}

	switch v := value.(type) {
	case int:
		return int64(v), true
	case int64:
		return v, true
	case float64:
		return int64(v), true
	}
	return
}

// StringVar 获取string类型变量
func (c *Conf) StringVar(key string) (val string, ok bool) {
	value, exists := c.Var[key]
	if !exists {
		return
	}

	val, ok = value.(string)
	return
}

// Build 配置初始化
func (c *Conf) Build() {
	c.App.Build()
	c.Id.Build()
}
