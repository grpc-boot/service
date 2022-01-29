package core

import (
	"net/http"
	"time"

	"service/common/components"
	"service/common/define/constant"
	"service/common/model"

	"github.com/gin-gonic/gin"
	"github.com/grpc-boot/base"
	"github.com/grpc-boot/gateway"
	"go.uber.org/zap"
)

// WithGateway 加载网关
func WithGateway() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var (
			accessTime = time.Now()
			path       = ctx.FullPath()
		)

		//初始化ctx
		ctx.Set(constant.CtxAccessTime, accessTime)
		ctx.Set(constant.CtxRequestPath, path)

		//加载网关
		gw, ok := components.GetGateway(constant.ContGateway)
		if !ok {
			ctx.Next()
			return
		}

		status, _, _ := gw.InTimeout(time.Millisecond*100, path)
		if status == gateway.StatusYes {
			ctx.Next()
			return
		}

		response := model.Response{
			Code: http.StatusRequestTimeout,
			Msg:  "service is busy",
			Data: base.SetValue,
		}

		if status == gateway.StatusNo {
			response.Code = http.StatusInternalServerError
			response.Msg = "service not available"
		}

		ctx.JSON(http.StatusOK, response)

		_, _, _, err := gw.Out(accessTime, path, response.Code)
		if err != nil {
			base.ZapError("gateway out error",
				zap.String(constant.ZapError, err.Error()),
			)
		}

		ctx.Abort()
		return
	}
}
