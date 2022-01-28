package core

import (
	clientv3 "go.etcd.io/etcd/client/v3"
	"service/common/components"
	"service/common/define"

	"github.com/grpc-boot/base"
)

func Init() {
	var (
		err  error
		conf = components.GetConf()
	)

	//初始化日志
	components.InitLogger()

	//加载数据库
	err = components.SetDb(define.Db, &conf.Db)
	if err != nil {
		base.RedFatal("load %s err:%s", define.Db, err.Error())
	}

	//加载Etcd
	err = components.SetEtcd(define.Etcd, conf.Etcd.ToConfig())
	if err != nil {
		base.RedFatal("load %s err:%s", define.Etcd, err.Error())
	}

	//加载etcd conf
	err = components.SetEtcdConf(define.EtcdConf, conf.EtcdConf, clientv3.WithPrefix())
	if err != nil {
		base.RedFatal("load %s err:%s", define.EtcdConf, err.Error())
	}

	//加载etcd naming
	err = components.SetEtcdNaming(define.EtcdNaming, conf.EtcdNaming)
	if err != nil {
		base.RedFatal("load %s err:%s", define.EtcdNaming, err.Error())
	}

	//加载redis
	components.SetRedis(define.Redis, &conf.Redis)
}
