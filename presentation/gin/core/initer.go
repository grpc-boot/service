package core

import (
	"service/common/components"
	"service/common/define/constant"

	"github.com/grpc-boot/base"
	clientv3 "go.etcd.io/etcd/client/v3"
)

func Init() {
	var (
		err  error
		conf = components.GetConf()
	)

	//初始化日志
	components.InitLogger()

	//加载数据库
	err = components.SetDb(constant.Db, &conf.Db)
	if err != nil {
		base.RedFatal("load %s err:%s", constant.Db, err.Error())
	}

	//加载Etcd
	err = components.SetEtcd(constant.Etcd, conf.Etcd.ToConfig())
	if err != nil {
		base.RedFatal("load %s err:%s", constant.Etcd, err.Error())
	}

	//加载etcd conf
	err = components.SetEtcdConf(constant.EtcdConf, conf.EtcdConf, clientv3.WithPrefix())
	if err != nil {
		base.RedFatal("load %s err:%s", constant.EtcdConf, err.Error())
	}

	//加载etcd naming
	err = components.SetEtcdNaming(constant.EtcdNaming, conf.EtcdNaming)
	if err != nil {
		base.RedFatal("load %s err:%s", constant.EtcdNaming, err.Error())
	}

	//加载redis
	components.SetRedis(constant.Redis, &conf.Redis)
}
