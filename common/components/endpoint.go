package components

import (
	"service/common/define/constant"

	"github.com/grpc-boot/base"
	"go.etcd.io/etcd/client/v3/naming/endpoints"
	"go.uber.org/zap"
)

func EndpointWithLocalIp(addr string, metadata map[string]interface{}) (endpoint endpoints.Endpoint, err error) {
	var ip string
	ip, err = base.LocalIp()
	if err != nil {
		return
	}

	defaultMeta := map[string]interface{}{}

	if len(metadata) > 0 {
		for key, value := range metadata {
			defaultMeta[key] = value
		}
	}

	endpoint = endpoints.Endpoint{
		Addr:     ip + addr,
		Metadata: defaultMeta,
	}
	return
}

func Register(endpoint endpoints.Endpoint) error {
	if naming, ok := GetEtcdNaming(constant.ContEtcdNaming); ok {
		liveCh, err := naming.Register(60, endpoint)
		if err != nil {
			base.ZapFatal("register service error", zap.String(constant.ZapError, err.Error()))
			return err
		}

		go func() {
			for {
				res, getOk := <-liveCh
				if !getOk {
					break
				}

				base.ZapInfo("etcd keepAlive info",
					zap.String(constant.ZapInfo, res.String()),
				)
			}
		}()
	}

	return nil
}

func DeRegister(endpoint endpoints.Endpoint) error {
	if naming, ok := GetEtcdNaming(constant.ContEtcdNaming); ok {
		base.ZapInfo("begin remove service", zap.String(constant.ZapEndpoint, endpoint.Addr))
		if err := naming.Del(endpoint); err != nil {
			base.ZapError("remove service error",
				zap.String(constant.ZapEndpoint, endpoint.Addr),
				zap.String(constant.ZapError, err.Error()),
			)
			return err
		}
	}

	return nil
}
