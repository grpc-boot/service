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
	"service/presentation/fasthttp/core"

	"github.com/grpc-boot/base"
	"github.com/valyala/fasthttp"
	"go.etcd.io/etcd/client/v3/naming/endpoints"
	"go.uber.org/zap"
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

	//加载路由接收请求
	r := core.Route()
	srv := fasthttp.Server{
		Handler: r.HandleRequest,
	}

	srv.ErrorHandler = func(ctx *fasthttp.RequestCtx, err error) {
		base.ZapError("error handler",
			zap.String(constant.ZapPath, base.Bytes2String(ctx.RequestURI())),
			zap.String(constant.ZapError, err.Error()),
			zap.String(constant.ZapIp, ctx.RemoteIP().String()),
		)
	}

	// 启动服务
	go func() {
		err = srv.ListenAndServe(conf.App.Addr)
		if err != nil && err != http.ErrServerClosed {
			base.ZapFatal("start server error", zap.String(constant.ZapError, err.Error()))
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

	//停止服务
	done := make(chan struct{}, 1)
	go func() {
		if err = srv.Shutdown(); err != nil {
			base.ZapFatal("stop server error", zap.String(constant.ZapError, err.Error()))
		}

		done <- base.SetValue
	}()

	//超时处理
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
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
