package main

import (
	"context"
	"math/rand"
	"net"
	"os"
	"os/signal"
	"syscall"
	"time"

	"service/common/components"
	"service/common/define/constant"
	"service/presentation/grpc/core"

	"github.com/grpc-boot/base"
	"go.etcd.io/etcd/client/v3/naming/endpoints"
	"go.uber.org/zap"
	"google.golang.org/grpc"
)

var (
	endpoint endpoints.Endpoint
	startAt  = time.Now()
)

func init() {
	components.SetInt64(constant.ContStartAt, startAt.Unix())
	rand.Seed(startAt.UnixNano())
}

func main() {
	//初始化配置文件
	err := components.InitConfWithFlag()
	if err != nil {
		base.ZapFatal("load config error", zap.String(constant.ZapError, err.Error()))
	}

	//获取配置文件
	conf := components.GetConf()

	// 开启了服务自动注册
	if conf.App.AutoRegis {
		//初始化endpoint
		endpoint, err = components.EndpointWithLocalIp(conf.App.Addr, map[string]interface{}{
			"startAt": startAt.Unix(),
		})
		if err != nil {
			base.ZapFatal("load endpoint error", zap.String(constant.ZapError, err.Error()))
		}
	}

	//应用初始化
	core.Init()

	opts := []grpc.ServerOption{
		grpc.UnaryInterceptor(core.Gateway),
	}
	srv := grpc.NewServer(opts...)

	// 注册服务
	core.Route(srv)

	go func() {
		lis, err := net.Listen("tcp", conf.App.Addr)
		if err != nil {
			base.ZapFatal("listen port error", zap.String(constant.ZapError, err.Error()))
		}

		if er := srv.Serve(lis); er != nil {
			base.ZapFatal("serve error", zap.String(constant.ZapError, er.Error()))
		}
	}()

	if endpoint.Addr != "" {
		//注册服务
		time.AfterFunc(time.Second, func() {
			if err = components.Register(endpoint); err != nil {
				base.ZapFatal("register server failed")
			}
		})
	}

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	// 取消注册
	shutdown()

	done := make(chan struct{}, 1)
	go func() {
		srv.GracefulStop()
		done <- base.SetValue
	}()

	//超时处理
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	select {
	case <-done:
	case <-ctx.Done():
		break
	}
}

func shutdown() {
	if endpoint.Addr != "" {
		//注销服务
		_ = components.DeRegister(endpoint)
	}
}
