package v1

import (
	"service/common/components"
	"service/common/define/constant"
	"service/common/model"

	"github.com/gin-gonic/gin"
)

// Gw 获取网关信息
func Gw(ctx *gin.Context) {
	gw, ok := components.GetGateway(constant.ContGateway)
	if !ok {
		json(ctx, components.Code2Response(constant.ErrGatewayNotExists))
		return
	}

	json(ctx, model.RespSuccess(gw.Info()))
}
