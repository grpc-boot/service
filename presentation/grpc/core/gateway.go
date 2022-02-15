package core

import (
	"context"
	"strconv"
	"time"

	"service/common/components"
	"service/common/define/constant"

	"github.com/grpc-boot/base"
	"github.com/grpc-boot/gateway"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

func Gateway(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp interface{}, err error) {
	var (
		accessTime = time.Now()
		path       = info.FullMethod
	)

	//初始化ctx
	id, _ := components.GetId(constant.ContId)
	traceId, _ := id.IdTime(accessTime, 0)

	md, _ := metadata.FromIncomingContext(ctx)
	md.Set(constant.CtxTraceId, strconv.FormatInt(traceId, 10))
	md.Set(constant.CtxAccessTime, accessTime.String())
	md.Set(constant.CtxRequestPath, path)

	//加载网关
	gw, ok := components.GetGateway(constant.ContGateway)
	if !ok {
		return handler(ctx, req)
	}

	gwStatus, _ := gw.In(path)
	if gwStatus == gateway.StatusYes {
		return handler(ctx, req)
	}

	code := constant.ErrLimit
	if gwStatus == gateway.StatusNo {
		code = constant.ErrNotAvailable
	}

	err = components.Code2Error(code)

	_, _, _, er := gw.Out(accessTime, path, int(code))
	if er != nil {
		base.ZapError("gateway out error",
			zap.String(constant.ZapError, err.Error()),
			zap.Int64(constant.ZapTraceId, traceId),
		)
	}

	return nil, err
}
