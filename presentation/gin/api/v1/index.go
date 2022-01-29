package v1

import (
	"service/common/model"

	"github.com/gin-gonic/gin"
	"github.com/grpc-boot/base"
)

func Index(ctx *gin.Context) {
	json(ctx, model.Response{
		Code: 20000,
		Msg:  "ok",
		Data: base.SetValue,
	})
}
