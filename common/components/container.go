package components

import (
	"sync/atomic"

	"github.com/grpc-boot/gedis"
	"github.com/grpc-boot/orm"
)

var (
	container atomic.Value //*map[string]interface{}
)

// Set 向容器里添加值
func Set(name string, val interface{}) {
	newContainer := make(map[string]interface{})
	oldContainer, ok := container.Load().(*map[string]interface{})
	if ok && len(*oldContainer) > 0 {
		for n, v := range *oldContainer {
			newContainer[n] = v
		}
	}
	newContainer[name] = val
}

// Get 从容器中取值
func Get(name string) (val interface{}, exists bool) {
	c, ok := container.Load().(*map[string]interface{})
	if !ok {
		return nil, false
	}

	val, exists = (*c)[name]
	return
}

// SetRedis 加载Redis连接池
func SetRedis(name string, option *gedis.Option) {
	Set(name, gedis.NewPool(*option))
}

// GetRedis 获取Redis连接池
func GetRedis(name string) (red gedis.Pool, exists bool) {
	val, ok := Get(name)
	if !ok {
		return nil, ok
	}

	red, exists = val.(gedis.Pool)
	return
}

// SetRedisGroup 加载RedisGroup
func SetRedisGroup(name string, options ...gedis.GroupOption) error {
	group, err := gedis.NewGroup(options...)
	if err != nil {
		return err
	}

	Set(name, group)
	return nil
}

// GetRedisGroup 获取RedisGroup
func GetRedisGroup(name string) (group gedis.Group, exists bool) {
	val, ok := Get(name)
	if !ok {
		return nil, ok
	}

	group, exists = val.(gedis.Group)
	return
}

// SetDb 加载数据库
func SetDb(name string, option *orm.GroupOption) error {
	group, err := orm.NewMysqlGroup(option)
	if err != nil {
		return err
	}

	Set(name, group)
	return nil
}

// GetDb 获取数据库
func GetDb(name string) (red orm.Group, exists bool) {
	val, ok := Get(name)
	if !ok {
		return nil, ok
	}

	red, exists = val.(orm.Group)
	return
}
