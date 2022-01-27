package core

import (
	"service/common/components"
	"service/common/define"

	"github.com/grpc-boot/base"
)

func Init() {
	var (
		err error
		conf = components.GetConf()
	)

	components.InitLogger()

	//加载数据库
	err = components.SetDb(define.UserDb, &conf.Db)
	if err != nil {
		base.RedFatal("load user-db err:%s", err.Error())
	}

	//加载redis
	components.SetRedis(define.UserRedis, &conf.Redis)
}
