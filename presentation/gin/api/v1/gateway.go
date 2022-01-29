package v1

import (
	"net/http"

	"service/common/components"
	"service/common/define/constant"
	"service/common/model"

	"github.com/gin-gonic/gin"
	"github.com/grpc-boot/base"
)

// Gw 获取网关信息
func Gw(ctx *gin.Context) {
	gw, ok := components.GetGateway(constant.ContGateway)
	if !ok {
		json(ctx, model.Response{
			Code: constant.ErrGatewayNotExists,
			Msg:  "网关不存在",
			Data: base.SetValue,
		})
		return
	}

	json(ctx, model.Response{
		Code: http.StatusOK,
		Msg:  "ok",
		Data: gw.Info(),
	})
}
