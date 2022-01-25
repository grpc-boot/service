package main

import (
	"strconv"

	"service/common/conf"
	"service/presentation/gin/core"

	"github.com/grpc-boot/base"
)

func main() {
	//初始化配置文件
	err := conf.InitWithFlag()
	if err != nil {
		base.RedFatal("加载配置文件错误:%s", err.Error())
	}

	//应用初始化
	core.Init()

	//加载路由接收请求
	r := core.Route()

	//监听端口
	r.Run(":" + strconv.FormatInt(int64(conf.GetConf().App.Port), 10))
}
