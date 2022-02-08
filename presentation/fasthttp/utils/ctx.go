package utils

import (
	"service/common/model"

	"github.com/grpc-boot/base"
	routing "github.com/qiangxue/fasthttp-routing"
)

func Json(ctx *routing.Context, obj model.Response) (length int, err error) {
	ctx.Response.Header.Set("Content-Type", "application/json; charset=utf-8")

	data, _ := base.JsonEncode(obj)
	return ctx.Write(data)
}

func Jsonp(ctx *routing.Context, obj model.Response) (length int, err error) {
	callback := ctx.QueryArgs().Peek("callback")
	if len(callback) < 1 {
		return Json(ctx, obj)
	}

	ctx.Response.Header.Set("Content-Type", "application/javascript; charset=utf-8")

	jData, _ := base.JsonEncode(obj)
	data := make([]byte, 0, len(callback)+len(jData)+3)
	data = append(data, callback...)
	data = append(data, '(')
	data = append(data, jData...)
	data = append(data, ')', ';')
	return ctx.Write(data)
}
