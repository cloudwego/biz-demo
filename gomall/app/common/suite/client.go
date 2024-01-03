package suite

import (
	"os"

	"github.com/cloudwego/biz-demo/gomall/app/common/utils"
	"github.com/cloudwego/kitex/client"
	"github.com/cloudwego/kitex/pkg/rpcinfo"
	"github.com/cloudwego/kitex/pkg/transmeta"
	"github.com/cloudwego/kitex/transport"
	consul "github.com/kitex-contrib/registry-consul"
)

type CommonGrpcClientSuite struct {
	DestServiceName    string
	DestServiceAddr    string
	CurrentServiceName string
}

func (s CommonGrpcClientSuite) Options() []client.Option {
	opts := []client.Option{
		client.WithMetaHandler(transmeta.ClientHTTP2Handler),
		client.WithTransportProtocol(transport.GRPC),
	}

	if os.Getenv("REGISTRY_ENABLE") == "true" {
		r, err := consul.NewConsulResolver(os.Getenv("REGISTRY_ADDR"))
		utils.MustHandleError(err)
		opts = append(opts, client.WithResolver(r))
	} else {
		opts = append(opts, client.WithHostPorts(s.DestServiceAddr))
	}
	opts = append(opts, client.WithClientBasicInfo(&rpcinfo.EndpointBasicInfo{
		ServiceName: s.CurrentServiceName,
	}))

	return opts
}
