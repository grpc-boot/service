package v1

import (
	"github.com/gin-gonic/gin"
	"github.com/grpc-boot/base"
	"go.uber.org/zap"
	"service/common/components"
	"service/common/define/constant"
	"service/common/model"
	"time"
)

func gateway(ctx *gin.Context, obj model.Response) {
	gw, ok := components.GetGateway(constant.ContGateway)
	if ok {
		at, _ := ctx.Get(constant.CtxAccessTime)
		path, _ := ctx.Get(constant.CtxRequestPath)
		_, _, _, err := gw.Out(at.(time.Time), path.(string), int(obj.Code))
		if err != nil {
			base.ZapError("gateway out error",
				zap.String(constant.ZapError, err.Error()),
			)
		}
	}
}
