package v1

import (
	"net/http"
	"time"

	"service/common/components"
	"service/common/define/constant"
	"service/common/model"

	"github.com/gin-gonic/gin"
	"github.com/grpc-boot/base"
	"go.uber.org/zap"
)

func gateway(ctx *gin.Context, obj model.Response) {
	gw, ok := components.GetGateway(constant.ContGateway)
	if ok {
		at, _ := ctx.Get(constant.CtxAccessTime)
		path, _ := ctx.Get(constant.CtxRequestPath)
		_, _, _, err := gw.Out(at.(time.Time), path.(string), obj.Code)
		if err != nil {
			base.ZapError("gateway out error",
				zap.String(constant.ZapError, err.Error()),
			)
		}
	}
}

func json(ctx *gin.Context, obj model.Response) {
	ctx.JSON(http.StatusOK, obj)
	gateway(ctx, obj)
}

func jsonp(ctx *gin.Context, obj model.Response) {
	ctx.JSONP(http.StatusOK, obj)
	gateway(ctx, obj)
}
