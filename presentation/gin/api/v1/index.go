package v1

import (
	"service/common/define/constant"
	"service/common/model"

	"github.com/gin-gonic/gin"
	"github.com/grpc-boot/base"
)

func Index(ctx *gin.Context) {
	json(ctx, model.Response{
		Code: constant.Success,
		Msg:  "Hello World!",
		Data: base.SetValue,
	})
}
