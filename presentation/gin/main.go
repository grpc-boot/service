package main

import (
	"context"
	"math/rand"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"service/common/components"
	"service/common/define/constant"
	"service/presentation/gin/core"

	"github.com/grpc-boot/base"
	"go.etcd.io/etcd/client/v3/naming/endpoints"
	"go.uber.org/zap"
)

var (
	endpoint endpoints.Endpoint
	startAt  = time.Now()
)

func init() {
	components.SetInt64(constant.StartAt, startAt.Unix())
	rand.Seed(startAt.UnixNano())
}

func main() {
	//初始化配置文件
	err := components.InitConfWithFlag()
	if err != nil {
		base.ZapFatal("load config error", zap.String(constant.FieldError, err.Error()))
	}

	//初始化endpoint
	endpoint, err = components.EndpointWithLocalIp(components.GetConf(), map[string]interface{}{
		"startAt": startAt.Unix(),
	})
	if err != nil {
		base.ZapFatal("load endpoint error", zap.String(constant.FieldError, err.Error()))
	}

	//应用初始化
	core.Init()
	//加载路由接收请求
	r := core.Route()

	srv := http.Server{
		Addr:    components.GetConf().App.Addr,
		Handler: r,
	}

	// 注册停止回调
	srv.RegisterOnShutdown(shutdown)

	go func() {
		err = srv.ListenAndServe()
		if err != nil && err != http.ErrServerClosed {
			base.ZapFatal("start server error", zap.String(constant.FieldError, err.Error()))
		}
	}()

	//注册服务
	time.AfterFunc(time.Second, func() {
		if naming, ok := components.GetEtcdNaming(constant.EtcdNaming); ok {
			if _, err = naming.Register(10, endpoint); err != nil {
				base.ZapFatal("register service error", zap.String(constant.FieldError, err.Error()))
			}
		}
	})

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()
	if err = srv.Shutdown(ctx); err != nil {
		base.ZapFatal("stop server error", zap.String(constant.FieldError, err.Error()))
	}
}

func shutdown() {
	base.ZapInfo("begin remove service", zap.String(constant.FieldEndpoint, endpoint.Addr))
	if naming, ok := components.GetEtcdNaming(constant.EtcdNaming); ok {
		if err := naming.Del(endpoint); err != nil {
			base.ZapError("remove service error",
				zap.String(constant.FieldEndpoint, endpoint.Addr),
				zap.String(constant.FieldError, err.Error()),
			)
		}
	}
}
