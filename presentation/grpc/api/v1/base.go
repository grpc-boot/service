package v1

import (
	"context"
	"time"

	"service/common/components"
	"service/common/define/constant"

	"github.com/grpc-boot/base"
	"go.uber.org/zap"
	"google.golang.org/grpc/metadata"
)

func gateway(ctx context.Context, data interface{}, code int) {
	gw, ok := components.GetGateway(constant.ContGateway)
	if ok {
		md, _ := metadata.FromIncomingContext(ctx)

		at := md.Get(constant.CtxAccessTime)
		path := md.Get(constant.CtxRequestPath)

		accessTime, _ := time.Parse("2006-01-02 15:04:05.999999999 -0700 MST", at[0])

		_, _, _, err := gw.Out(accessTime, path[0], code)

		if err != nil {
			base.ZapError("gateway out error",
				zap.String(constant.ZapError, err.Error()),
			)
		}
	}
}
