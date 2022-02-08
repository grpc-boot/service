package v1

import (
	"service/common/components"
	"service/common/define/constant"
	"service/common/model"

	routing "github.com/qiangxue/fasthttp-routing"
)

// Gw 获取网关信息
func Gw(ctx *routing.Context) error {
	gw, ok := components.GetGateway(constant.ContGateway)
	if !ok {
		_, err := json(ctx, components.Code2Response(constant.ErrGatewayNotExists))
		return err
	}

	_, err := json(ctx, model.RespSuccess(gw.Info()))
	return err
}
