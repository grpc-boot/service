package core

import (
	"strconv"

	"service/common/components"
	"service/common/define/constant"

	"github.com/grpc-boot/base"
	"github.com/grpc-boot/gateway"
	clientv3 "go.etcd.io/etcd/client/v3"
	"go.uber.org/zap"
)

func Init() {
	var (
		err  error
		conf = components.GetConf()
	)

	//初始化日志
	components.InitLogger()

	//加载数据库
	err = components.SetDb(constant.ContDb, &conf.Db)
	if err != nil {
		base.RedFatal("load %s err:%s", constant.ContDb, err.Error())
	}

	//加载Etcd
	err = components.SetEtcd(constant.ContEtcd, conf.Etcd.ToConfig())
	if err != nil {
		base.RedFatal("load %s err:%s", constant.ContEtcd, err.Error())
	}

	//加载etcd conf
	err = components.SetEtcdConf(constant.ContEtcdConf, conf.EtcdConf, clientv3.WithPrefix())
	if err != nil {
		base.RedFatal("load %s err:%s", constant.ContEtcdConf, err.Error())
	}

	//加载etcd naming
	err = components.SetEtcdNaming(constant.ContEtcdNaming, conf.EtcdNaming)
	if err != nil {
		base.RedFatal("load %s err:%s", constant.ContEtcdNaming, err.Error())
	}

	//加载redis
	components.SetRedis(constant.ContRedis, &conf.Redis)

	//加载网关
	components.SetGateway(constant.ContGateway, 10, func() (options []gateway.Option) {
		db, ok := components.GetDb(constant.ContDb)
		if !ok {
			base.Info("get db error")
			return
		}

		var rows []map[string]string
		rows, err = db.Query(false, "SELECT * FROM `gateway`")
		if err != nil {
			base.ZapError("OptionsWithMysql: get options form mysql table error",
				zap.String(constant.ZapError, err.Error()),
			)
			return
		}

		if len(rows) < 1 {
			base.ZapError("OptionsWithMysql: get options form mysql table error",
				zap.String(constant.ZapError, "table empty"),
			)
			return
		}

		options = make([]gateway.Option, len(rows))
		for index, row := range rows {
			limit, _ := strconv.Atoi(row["second_limit"])
			options[index] = gateway.Option{
				Name:        row["name"],
				Path:        row["path"],
				SecondLimit: limit,
			}
		}

		return
	})
}
