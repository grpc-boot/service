package components

import (
	"service/common/define/constant"

	"github.com/grpc-boot/base"
)

// InitComponents 初始化组件
func InitComponents() {
	var err error

	//初始化id配置
	conf.Id.Build()

	//加载id生成器
	err = SetIdIp(constant.ContId, conf.Id)
	if err != nil {
		base.RedFatal("load snowflake err:%s", err.Error())
	}

	//加载数据库
	if len(conf.Db.Masters) > 0 {
		err = SetDb(constant.ContDb, &conf.Db)
		if err != nil {
			base.RedFatal("load %s err:%s", constant.ContDb, err.Error())
		}
	}

	//加载redis
	if conf.Redis.Host != "" {
		SetRedis(constant.ContRedis, &conf.Redis)
	}

	//加载Etcd
	if len(conf.Etcd.Endpoints) > 0 {
		err = SetEtcd(constant.ContEtcd, conf.Etcd.ToConfig())
		if err != nil {
			base.RedFatal("load %s err:%s", constant.ContEtcd, err.Error())
		}
	}

	//加载etcd naming
	if conf.EtcdNaming.Service != "" {
		err = SetEtcdNaming(constant.ContEtcdNaming, conf.EtcdNaming)
		if err != nil {
			base.RedFatal("load %s err:%s", constant.ContEtcdNaming, err.Error())
		}
	}
}
