package main

import (
	"github.com/baiyutang/gomall/app/product/biz/dal"
	"log"
	"net"
	"os"

	"github.com/baiyutang/gomall/app/product/conf"
	"github.com/baiyutang/gomall/app/product/kitex_gen/product/productcatalogservice"
	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/cloudwego/kitex/pkg/rpcinfo"
	"github.com/cloudwego/kitex/server"
	kitexzap "github.com/kitex-contrib/obs-opentelemetry/logging/zap"
	consul "github.com/kitex-contrib/registry-consul"
)

func main() {
	opts := kitexInit()

	svr := productcatalogservice.NewServer(new(ProductCatalogServiceImpl), opts...)
	log.Fatal(svr.Run())
}

func kitexInit() (opts []server.Option) {
	// address
	dal.Init()
	addr, err := net.ResolveTCPAddr("tcp", conf.GetConf().Kitex.Address)
	if err != nil {
		panic(err)
	}
	opts = append(opts, server.WithServiceAddr(addr))

	// service info
	opts = append(opts, server.WithServerBasicInfo(&rpcinfo.EndpointBasicInfo{
		ServiceName: conf.GetConf().Kitex.Service,
	}))
	klog.Info(os.Getenv("REGISTRY_ADDR"))
	if os.Getenv("REGISTRY_ENABLE") == "true" {
		r, err := consul.NewConsulRegister(os.Getenv("REGISTRY_ADDR"))
		if err != nil {
			klog.Fatal(err)
		}
		opts = append(opts, server.WithRegistry(r))
	}

	// klog
	if os.Getenv("GO_ENV") != "online" {
		klog.SetLevel(klog.LevelDebug)
		return
	}
	logger := kitexzap.NewLogger()
	klog.SetLogger(logger)
	klog.SetOutput(os.Stdout)
	return
}
