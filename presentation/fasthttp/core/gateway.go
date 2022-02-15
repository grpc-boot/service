package core

import (
	"time"

	"service/common/components"
	"service/common/define/constant"
	"service/presentation/fasthttp/utils"

	"github.com/grpc-boot/base"
	"github.com/grpc-boot/gateway"
	routing "github.com/qiangxue/fasthttp-routing"
	"go.uber.org/zap"
)

// Gateway 加载网关
func Gateway(ctx *routing.Context) error {
	var (
		err        error
		accessTime = time.Now()
		path       = base.Bytes2String(ctx.Path())
	)

	//初始化ctx
	id, _ := components.GetId(constant.ContId)
	traceId, _ := id.IdTime(accessTime, 0)
	ctx.SetUserValue(constant.CtxTraceId, traceId)
	ctx.SetUserValue(constant.CtxAccessTime, accessTime)
	ctx.SetUserValue(constant.CtxRequestPath, path)

	//加载网关
	gw, ok := components.GetGateway(constant.ContGateway)
	if !ok {
		err = ctx.Next()
		if err != nil {
			base.ZapError("handler error",
				zap.String(constant.ZapError, err.Error()),
				zap.Int64(constant.ZapTraceId, traceId),
			)
		}

		return err
	}

	status, _ := gw.In(path)
	if status == gateway.StatusYes {
		err = ctx.Next()
		if err != nil {
			base.ZapError("handler error",
				zap.String(constant.ZapError, err.Error()),
				zap.Int64(constant.ZapTraceId, traceId),
			)
		}

		return err
	}

	code := constant.ErrLimit
	if status == gateway.StatusNo {
		code = constant.ErrNotAvailable
		_, err = utils.Jsonp(ctx, components.Code2Response(constant.ErrNotAvailable))
	} else {
		_, err = utils.Jsonp(ctx, components.Code2Response(constant.ErrLimit))
	}

	if err != nil {
		base.ZapError("write error",
			zap.String(constant.ZapError, err.Error()),
			zap.Int64(constant.ZapTraceId, traceId),
		)
	}

	_, _, _, err = gw.Out(accessTime, path, int(code))
	if err != nil {
		base.ZapError("gateway out error",
			zap.String(constant.ZapError, err.Error()),
			zap.Int64(constant.ZapTraceId, traceId),
		)
	}

	ctx.Abort()
	return nil
}
