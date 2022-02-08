package v1

import (
	"time"

	"service/common/components"
	"service/common/define/constant"
	"service/common/model"
	"service/presentation/fasthttp/utils"

	"github.com/grpc-boot/base"
	routing "github.com/qiangxue/fasthttp-routing"
	"github.com/valyala/fasthttp"
	"go.uber.org/zap"
)

func stringArg(key string, args *fasthttp.Args) (value string) {
	val := args.Peek(key)
	if len(val) < 1 {
		return
	}

	return base.Bytes2String(val)
}

func gateway(ctx *routing.Context, obj model.Response) {
	gw, ok := components.GetGateway(constant.ContGateway)
	if ok {
		at := ctx.UserValue(constant.CtxAccessTime)
		path := ctx.UserValue(constant.CtxRequestPath)
		_, _, _, err := gw.Out(at.(time.Time), path.(string), int(obj.Code))
		if err != nil {
			base.ZapError("gateway out error",
				zap.String(constant.ZapError, err.Error()),
			)
		}
	}
}

func json(ctx *routing.Context, obj model.Response) (length int, err error) {
	length, err = utils.Json(ctx, obj)
	gateway(ctx, obj)
	return
}

func jsonp(ctx *routing.Context, obj model.Response) (length int, err error) {
	length, err = jsonp(ctx, obj)
	gateway(ctx, obj)
	return
}
