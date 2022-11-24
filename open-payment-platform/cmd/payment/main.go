package main

import (
	"net"

	"github.com/cloudwego/biz-demo/open-payment-platform/kitex_gen/payment/paymentsvc"
	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/cloudwego/kitex/pkg/rpcinfo"
	"github.com/cloudwego/kitex/pkg/transmeta"
	"github.com/cloudwego/kitex/server"
	"github.com/kitex-contrib/registry-nacos/registry"
)

func main() {
	r, err := registry.NewDefaultNacosRegistry()
	if err != nil {
		klog.Fatal(err)
	}

	hdlr := initHandler()

	svr := paymentsvc.NewServer(
		hdlr,
		server.WithRegistry(r),
		server.WithServerBasicInfo(&rpcinfo.EndpointBasicInfo{ServiceName: "payment"}),
		server.WithServiceAddr(&net.TCPAddr{Port: 8081}),
		server.WithMetaHandler(transmeta.ServerTTHeaderHandler),
	)

	err = svr.Run()
	if err != nil {
		klog.Fatal(err)
	}
}
