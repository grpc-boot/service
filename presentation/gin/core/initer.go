package core

import (
	"strconv"
	"time"

	"service/common/components"
	"service/common/define/constant"

	"github.com/grpc-boot/base"
	"github.com/grpc-boot/base/core/shardmap"
	"github.com/grpc-boot/gateway"
	"go.uber.org/zap"
)

func Init() {
	conf := components.GetConf()
	//初始化日志
	components.InitLogger()

	//初始化组件
	components.InitComponents()

	//加载etcd conf
	if len(conf.EtcdConf.Option.PrefixList) > 0 {
		changeCh, err := components.SetEtcdConf(constant.ContEtcdConf, conf.EtcdConf)
		if err != nil {
			base.RedFatal("load %s err:%s", constant.ContEtcdConf, err.Error())
		}

		go initChangeChan(changeCh)
	}

	//加载网关
	components.SetGateway(constant.ContGateway, time.Second, func() (options []gateway.Option) {
		db, ok := components.GetDb(constant.ContDb)
		if !ok {
			base.Info("get db error")
			return
		}

		rows, err := db.Query(false, "SELECT * FROM `gateway`")
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
			bucketSize, _ := strconv.Atoi(row["bucket_size"])
			options[index] = gateway.Option{
				Name:        row["name"],
				Path:        row["path"],
				SecondLimit: limit,
				BucketSize:  bucketSize,
			}
		}

		return
	})
}

func initChangeChan(ch <-chan shardmap.ChangeEvent) {
	for {
		event, ok := <-ch
		if !ok {
			break
		}

		switch event.Type {
		case shardmap.Create:
			base.Fuchsia("create key:%s value:%+v", event.Key, event.Value)
		case shardmap.Update:
			base.Fuchsia("update key:%s value:%+v", event.Key, event.Value)
		case shardmap.Delete:
			base.Fuchsia("delete key:%s", event.Key)
		}
	}
}
