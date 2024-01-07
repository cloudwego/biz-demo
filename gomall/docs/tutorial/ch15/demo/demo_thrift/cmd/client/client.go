package main

import (
	"context"
	"fmt"

	"github.com/cloudwego/biz-demo/gomall/demo/demo_thrift/conf"
	"github.com/cloudwego/biz-demo/gomall/demo/demo_thrift/kitex_gen/api"
	"github.com/cloudwego/biz-demo/gomall/demo/demo_thrift/kitex_gen/api/echo"
	"github.com/cloudwego/kitex/client"
	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/cloudwego/kitex/pkg/rpcinfo"
	"github.com/cloudwego/kitex/pkg/transmeta"
	"github.com/cloudwego/kitex/transport"
	consul "github.com/kitex-contrib/registry-consul"
)

func main() {
	r, err := consul.NewConsulResolver(conf.GetConf().Registry.RegistryAddress[0])
	if err != nil {
		panic(err)
	}
	c, err := echo.NewClient("demo_thrift", client.WithResolver(r),
		client.WithTransportProtocol(transport.TTHeader),
		client.WithMetaHandler(transmeta.ClientTTHeaderHandler),
		client.WithClientBasicInfo(&rpcinfo.EndpointBasicInfo{ServiceName: "demo_thrift_client"}),
	)
	if err != nil {
		panic(err)
	}
	res, err := c.Echo(context.Background(), &api.Request{Message: "hello"})
	if err != nil {
		klog.Fatal(err)
	}
	fmt.Printf("%v", res)
}
