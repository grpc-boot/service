package main

import (
	"context"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"service/common/components"
	"service/common/define"
	"service/presentation/gin/core"

	"github.com/grpc-boot/base"
	"go.etcd.io/etcd/client/v3/naming/endpoints"
	"go.uber.org/zap"
)

func main() {
	//初始化配置文件
	err := components.InitConfWithFlag()
	if err != nil {
		base.ZapFatal("load config error", zap.String("Error", err.Error()))
	}

	//应用初始化
	core.Init()
	//加载路由接收请求
	r := core.Route()

	srv := http.Server{
		Addr:    components.GetConf().App.Addr,
		Handler: r,
	}

	var ip string
	ip, err = base.LocalIp()
	if err != nil {
		base.ZapFatal("get local ip error", zap.String("Error", err.Error()))
	}
	endpoint := endpoints.Endpoint{
		Addr: ip + components.GetConf().App.Addr,
		Metadata: map[string]interface{}{
			"level":    100,
			"protocal": "http",
			"name":     "service",
		},
	}

	// 注册停止回调
	srv.RegisterOnShutdown(func() {
		base.ZapInfo("begin remove service", zap.String("Addr", endpoint.Addr))

		if naming, ok := components.GetEtcdNaming(define.EtcdNaming); ok {
			if err = naming.Del(endpoint); err != nil {
				base.ZapError("remove service error",
					zap.String("Addr", endpoint.Addr),
					zap.String("Error", err.Error()),
				)
			}
		}
	})

	go func() {
		err = srv.ListenAndServe()
		if err != nil && err != http.ErrServerClosed {
			base.ZapFatal("start server error", zap.String("Error", err.Error()))
		}
	}()

	//注册服务
	time.AfterFunc(time.Second, func() {
		if naming, ok := components.GetEtcdNaming(define.EtcdNaming); ok {
			if _, err = naming.Register(10, endpoint); err != nil {
				base.ZapFatal("register service error", zap.String("Error", err.Error()))
			}
		}
	})

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()
	if err = srv.Shutdown(ctx); err != nil {
		base.ZapFatal("stop server error", zap.String("Error", err.Error()))
	}
}
