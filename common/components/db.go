package components

import (
	"sync/atomic"

	"github.com/grpc-boot/orm"
)

var (
	dbGroup atomic.Value //*map[string]orm.Group
)

// Db 获取数据库Group
func Db(logic string) (db orm.Group, exists bool) {
	p, ok := dbGroup.Load().(*map[string]orm.Group)
	if !ok {
		return nil, false
	}

	db, exists = (*p)[logic]
	return
}

// LoadDb 加载数据库
func LoadDb(logic string, option *orm.GroupOption) error {
	group, err := orm.NewMysqlGroup(option)
	if err != nil {
		return err
	}

	newGroup := make(map[string]orm.Group, 8)
	p, ok := dbGroup.Load().(*map[string]orm.Group)
	//拷贝
	if ok && len(*p) > 0 {
		for l, g := range *p {
			newGroup[l] = g
		}
	}

	newGroup[logic] = group
	dbGroup.Store(&newGroup)
	return nil
}
