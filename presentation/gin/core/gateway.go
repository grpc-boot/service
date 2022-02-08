package core

import (
	"net/http"
	"time"

	"service/common/components"
	"service/common/define/constant"

	"github.com/gin-gonic/gin"
	"github.com/grpc-boot/base"
	"github.com/grpc-boot/gateway"
	"go.uber.org/zap"
)

// Gateway 加载网关
func Gateway(ctx *gin.Context) {
	var (
		accessTime = time.Now()
		path       = ctx.FullPath()
	)

	//初始化ctx
	id, _ := components.GetId(constant.ContId)
	traceId, _ := id.IdTime(accessTime, 0)
	ctx.Set(constant.CtxTraceId, traceId)
	ctx.Set(constant.CtxAccessTime, accessTime)
	ctx.Set(constant.CtxRequestPath, path)

	//加载网关
	gw, ok := components.GetGateway(constant.ContGateway)
	if !ok {
		ctx.Next()
		return
	}

	status, _ := gw.In(path)
	if status == gateway.StatusYes {
		ctx.Next()
		return
	}

	code := constant.ErrLimit
	if status == gateway.StatusNo {
		code = constant.ErrNotAvailable
		ctx.JSON(http.StatusOK, components.Code2Response(constant.ErrNotAvailable))
	} else {
		ctx.JSON(http.StatusOK, components.Code2Response(constant.ErrLimit))
	}

	_, _, _, err := gw.Out(accessTime, path, code)
	if err != nil {
		base.ZapError("gateway out error",
			zap.String(constant.ZapError, err.Error()),
			zap.Int64(constant.ZapTraceId, traceId),
		)
	}

	ctx.Abort()
	return
}
