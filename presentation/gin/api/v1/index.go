package v1

import (
	"context"
	"net/http"
	"service/common/define/constant"

	"github.com/gin-gonic/gin"
	"github.com/grpc-boot/base"
	"go.uber.org/zap"
	"service/common/components"
)

func Index(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, map[string]interface{}{
		"code": 20000,
		"msg":  "ok",
		"data": struct{}{},
	})
}

func Conf(ctx *gin.Context) {
	c, ok := components.GetEtcdConf(constant.EtcdConf)
	if ok {
		v, exists := c.Get(`browser/conf/app/limit`)
		ctx.JSON(http.StatusOK, map[string]interface{}{
			"value":  v,
			"exists": exists,
		})
		return
	}

	ctx.JSON(http.StatusOK, map[string]interface{}{
		"code": 404,
		"msg":  "not exists",
		"data": struct{}{},
	})
}

func Service(ctx *gin.Context) {
	n, ok := components.GetEtcdNaming(constant.EtcdNaming)
	if ok {
		v, err := n.List(context.TODO())
		if err == nil {
			ctx.JSON(http.StatusOK, v)
			return
		}
		base.ZapError("got service error", zap.String("Error", err.Error()))
		return
	}

	ctx.JSON(http.StatusOK, map[string]interface{}{
		"code": 404,
		"msg":  "not exists",
		"data": struct{}{},
	})
}
