package core

import (
	"strconv"

	"service/common/components"
	"service/common/define/constant"

	"github.com/grpc-boot/base"
	"github.com/grpc-boot/gateway"
	"go.uber.org/zap"
)

func Init() {
	//初始化日志
	components.InitLogger()

	//初始化组件
	components.InitComponents()

	//加载网关
	components.SetGateway(constant.ContGateway, 10, func() (options []gateway.Option) {
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
			options[index] = gateway.Option{
				Name:        row["name"],
				Path:        row["path"],
				SecondLimit: limit,
			}
		}

		return
	})
}
