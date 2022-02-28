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

func gateway(ctx context.Context, data interface{}, code int32) {
	gw, ok := components.GetGateway(constant.ContGateway)
	if ok {
		md, _ := metadata.FromIncomingContext(ctx)

		at := md.Get(constant.CtxAccessTime)
		path := md.Get(constant.CtxRequestPath)
		if len(at) > 0 {
			accessTime, _ := time.ParseInLocation(constant.NanoFormat, at[0], time.Local)
			_, _, _, err := gw.Out(accessTime, path[0], int(code))

			if err != nil {
				base.ZapError("gateway out error",
					zap.String(constant.ZapError, err.Error()),
				)
			}
		}
	}
}
