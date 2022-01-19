package dal

import (
	"service/common/conf"

	"github.com/grpc-boot/orm"
)

var (
	userDb orm.Group
)

func InitDb() (err error) {
	config := conf.GetConf()
	userDb, err = orm.NewMysqlGroup(&config.Db)

	return err
}

func UserDb() orm.Group {
	return userDb
}
