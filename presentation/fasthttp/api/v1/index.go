package v1

import (
	"service/common/define/constant"
	"service/common/model"

	"github.com/grpc-boot/base"
	routing "github.com/qiangxue/fasthttp-routing"
)

func Index(ctx *routing.Context) error {
	_, err := json(ctx, model.Response{
		Code: constant.Success,
		Msg:  "Hello World!",
		Data: base.SetValue,
	})

	return err
}
