package main

import (
	"strconv"

	"service/common/components"
	"service/presentation/gin/core"

	"github.com/grpc-boot/base"
)

func main() {
	//初始化配置文件
	err := components.InitConfWithFlag()
	if err != nil {
		base.RedFatal("加载配置文件错误:%s", err.Error())
	}

	//应用初始化
	core.Init()

	//加载路由接收请求
	r := core.Route()

	//监听端口
	r.Run(":" + strconv.FormatInt(int64(components.GetConf().App.Port), 10))
}
