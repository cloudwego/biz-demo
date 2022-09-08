package main

import (
	"context"
	"net"

	"github.com/cloudwego/biz-demo/bookinfo/kitex_gen/cwg/bookinfo/details/detailsservice"
	"github.com/cloudwego/biz-demo/bookinfo/pkg/constants"
	"github.com/cloudwego/kitex/server"
	"github.com/kitex-contrib/obs-opentelemetry/provider"
	"github.com/kitex-contrib/obs-opentelemetry/tracing"
)

func main() {
	p := provider.NewOpenTelemetryProvider(
		provider.WithServiceName(constants.DetailsServiceName),
		provider.WithInsecure(),
	)
	defer p.Shutdown(context.Background())

	addr, err := net.ResolveTCPAddr("tcp", ":8084")
	if err != nil {
		return
	}

	svr := detailsservice.NewServer(
		NewHandler(),
		server.WithServiceAddr(addr),
		server.WithSuite(tracing.NewServerSuite()),
	)
	if err := svr.Run(); err != nil {
		return
	}
}
