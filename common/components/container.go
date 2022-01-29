package components

import (
	"errors"
	"fmt"
	"sync"
	"time"

	"service/common/model"

	"github.com/grpc-boot/betcd"
	"github.com/grpc-boot/gateway"
	"github.com/grpc-boot/gedis"
	"github.com/grpc-boot/orm"
	clientv3 "go.etcd.io/etcd/client/v3"
)

// 存储全局对象，适用于：极少量的insert，大量的query
// 注意：1.禁止存储大量用户或者连接数据 2.防止Key冲突
var container sync.Map

// GetAny 获取any
func GetAny(name string) (val interface{}, exists bool) {
	return container.Load(name)
}

// SetAny 存储any
func SetAny(name string, val interface{}) {
	container.Store(name, val)
}

// SetFloat64 存储float64值
func SetFloat64(name string, val float64) {
	container.Store(name, val)
}

// GetFloat64 获取float64值
func GetFloat64(name string) (val float64, ok bool) {
	value, exists := container.Load(name)
	if !exists {
		return
	}
	val, ok = value.(float64)
	return
}

// SetInt64 存储int64值
func SetInt64(name string, val int64) {
	container.Store(name, val)
}

// GetInt64 获取int64值
func GetInt64(name string) (val int64, ok bool) {
	value, exists := container.Load(name)
	if !exists {
		return
	}
	val, ok = value.(int64)
	return
}

// SetString 存储string值
func SetString(name, val string) {
	container.Store(name, val)
}

// SetBytes 存储[]byte
func SetBytes(name string, val []byte) {
	container.Store(name, val)
}

// GetBytes 获取[]byte
func GetBytes(name string) (val []byte, ok bool) {
	value, exists := container.Load(name)
	if !exists {
		return
	}
	val, ok = value.([]byte)
	return
}

// GetString 获取
func GetString(name string) (val string, ok bool) {
	value, exists := container.Load(name)
	if !exists {
		return
	}
	val, ok = value.(string)
	return
}

// SetRedis 加载Redis连接池
func SetRedis(name string, option *gedis.Option) {
	container.Store(name, gedis.NewPool(*option))
}

// GetRedis 获取Redis连接池
func GetRedis(name string) (red gedis.Pool, ok bool) {
	val, exists := container.Load(name)
	if !exists {
		return nil, exists
	}

	red, ok = val.(gedis.Pool)
	return
}

// SetRedisGroup 加载RedisGroup
func SetRedisGroup(name string, options ...gedis.GroupOption) error {
	group, err := gedis.NewGroup(options...)
	if err != nil {
		return err
	}

	container.Store(name, group)
	return nil
}

// GetRedisGroup 获取RedisGroup
func GetRedisGroup(name string) (group gedis.Group, ok bool) {
	val, exists := container.Load(name)
	if !exists {
		return nil, exists
	}

	group, ok = val.(gedis.Group)
	return
}

// SetDb 加载数据库
func SetDb(name string, option *orm.GroupOption) error {
	group, err := orm.NewMysqlGroup(option)
	if err != nil {
		return err
	}

	container.Store(name, group)
	return nil
}

// GetDb 获取数据库
func GetDb(name string) (db orm.Group, ok bool) {
	val, exists := container.Load(name)
	if !exists {
		return nil, exists
	}

	db, ok = val.(orm.Group)
	return
}

// SetEtcd 加载Etcd
func SetEtcd(name string, conf clientv3.Config) error {
	client, err := clientv3.New(conf)
	if err != nil {
		return err
	}

	container.Store(name, client)
	return nil
}

// GetEtcd 获取
func GetEtcd(name string) (client *clientv3.Client, ok bool) {
	val, exists := container.Load(name)
	if !exists {
		return nil, exists
	}

	client, ok = val.(*clientv3.Client)
	return
}

// SetEtcdConf 加载EtcdConf
func SetEtcdConf(name string, conf model.EtcdConf, options ...clientv3.OpOption) error {
	client, ok := GetEtcd(conf.Name)
	if !ok {
		return errors.New(fmt.Sprintf("没有找到name为%s的etcd服务", conf.Name))
	}

	c, err := betcd.NewConfigWithClient(client, conf.Prefix, conf.KeyOption, options...)
	if err != nil {
		return err
	}

	container.Store(name, c)
	return nil
}

// GetEtcdConf 获取EtcdConf
func GetEtcdConf(name string) (conf betcd.Config, ok bool) {
	val, exists := container.Load(name)
	if !exists {
		return nil, exists
	}

	conf, ok = val.(betcd.Config)
	return
}

// SetEtcdNaming 加载etcd naming
func SetEtcdNaming(name string, conf model.EtcdNaming) error {
	client, ok := GetEtcd(conf.Name)
	if !ok {
		return errors.New(fmt.Sprintf("没有找到name为%s的etcd服务", conf.Name))
	}

	naming, err := betcd.NewNamingWithClient(client, conf.Service)
	if err != nil {
		return err
	}

	container.Store(name, naming)
	return nil
}

// GetEtcdNaming 获取etcd naming
func GetEtcdNaming(name string) (naming betcd.Naming, ok bool) {
	val, exists := container.Load(name)
	if !exists {
		return nil, exists
	}

	naming, ok = val.(betcd.Naming)
	return
}

// SetGateway 存储网关
func SetGateway(name string, duration time.Duration, optionsFunc gateway.OptionsFunc) {
	gw := gateway.NewGateway(duration, optionsFunc)
	container.Store(name, gw)
}

// GetGateway 获取网关
func GetGateway(name string) (gw gateway.Gateway, ok bool) {
	val, exists := container.Load(name)
	if !exists {
		return
	}

	gw, ok = val.(gateway.Gateway)

	return
}
