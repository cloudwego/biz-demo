package suite

import (
	"fmt"
	"os"

	"github.com/baiyutang/gomall/app/common/utils"
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
		opts = append(opts, client.WithHostPorts("localhost:8881"))
	}
	opts = append(opts, client.WithClientBasicInfo(&rpcinfo.EndpointBasicInfo{
		ServiceName: fmt.Sprintf("%s-%s-client", s.CurrentServiceName, s.DestServiceName),
	}))

	return opts
}
