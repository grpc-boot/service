package components

import (
	"errors"
	"fmt"
	"sync/atomic"

	"service/common/model"

	"github.com/grpc-boot/betcd"
	"github.com/grpc-boot/gedis"
	"github.com/grpc-boot/orm"
	clientv3 "go.etcd.io/etcd/client/v3"
)

var (
	container atomic.Value //*map[string]interface{}
)

func setContainer(name string, val interface{}) {
	newContainer := make(map[string]interface{})
	oldContainer, ok := container.Load().(*map[string]interface{})
	if ok && len(*oldContainer) > 0 {
		for n, v := range *oldContainer {
			newContainer[n] = v
		}
	}
	newContainer[name] = val
	container.Store(&newContainer)
}

func getContainer(name string) (val interface{}, exists bool) {
	c, ok := container.Load().(*map[string]interface{})
	if !ok {
		return nil, ok
	}

	val, exists = (*c)[name]
	return
}

// SetRedis 加载Redis连接池
func SetRedis(name string, option *gedis.Option) {
	setContainer(name, gedis.NewPool(*option))
}

// GetRedis 获取Redis连接池
func GetRedis(name string) (red gedis.Pool, ok bool) {
	val, exists := getContainer(name)
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

	setContainer(name, group)
	return nil
}

// GetRedisGroup 获取RedisGroup
func GetRedisGroup(name string) (group gedis.Group, ok bool) {
	val, exists := getContainer(name)
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

	setContainer(name, group)
	return nil
}

// GetDb 获取数据库
func GetDb(name string) (red orm.Group, ok bool) {
	val, exists := getContainer(name)
	if !exists {
		return nil, exists
	}

	red, ok = val.(orm.Group)
	return
}

// SetEtcd 加载Etcd
func SetEtcd(name string, conf clientv3.Config) error {
	client, err := clientv3.New(conf)
	if err != nil {
		return err
	}

	setContainer(name, client)
	return nil
}

// GetEtcd 获取
func GetEtcd(name string) (client *clientv3.Client, ok bool) {
	val, exists := getContainer(name)
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

	setContainer(name, c)
	return nil
}

// GetEtcdConf 获取EtcdConf
func GetEtcdConf(name string) (conf betcd.Config, ok bool) {
	val, exists := getContainer(name)
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

	setContainer(name, naming)
	return nil
}

// GetEtcdNaming 获取etcd naming
func GetEtcdNaming(name string) (naming betcd.Naming, ok bool) {
	val, exists := getContainer(name)
	if !exists {
		return nil, exists
	}

	naming, ok = val.(betcd.Naming)
	return
}
