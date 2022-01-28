package components

import (
	"service/common/model"

	"github.com/grpc-boot/base"
	"go.etcd.io/etcd/client/v3/naming/endpoints"
)

func EndpointWithLocalIp(conf *model.Conf, metadata map[string]interface{}) (endpoint endpoints.Endpoint, err error) {
	var ip string
	ip, err = base.LocalIp()
	if err != nil {
		return
	}

	defaultMeta := map[string]interface{}{
		"name": conf.App.Name,
	}

	if len(metadata) > 0 {
		for key, value := range metadata {
			defaultMeta[key] = value
		}
	}

	endpoint = endpoints.Endpoint{
		Addr:     ip + conf.App.Addr,
		Metadata: defaultMeta,
	}
	return
}
