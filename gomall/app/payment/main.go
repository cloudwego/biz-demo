// Copyright 2024 CloudWeGo Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package main

import (
	"net"
	"os"
	"strings"

	"github.com/cloudwego/biz-demo/gomall/app/payment/biz/dal"
	"github.com/cloudwego/biz-demo/gomall/app/payment/conf"
	"github.com/cloudwego/biz-demo/gomall/app/payment/infra/mtl"
	"github.com/cloudwego/biz-demo/gomall/app/payment/middleware"
	"github.com/cloudwego/biz-demo/gomall/common/utils"
	"github.com/cloudwego/biz-demo/gomall/rpc_gen/kitex_gen/payment/paymentservice"
	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/cloudwego/kitex/pkg/rpcinfo"
	"github.com/cloudwego/kitex/pkg/transmeta"
	"github.com/cloudwego/kitex/server"
	"github.com/joho/godotenv"
	"github.com/kitex-contrib/config-etcd/etcd"
	etcdServer "github.com/kitex-contrib/config-etcd/server"
	prometheus "github.com/kitex-contrib/monitor-prometheus"
	"github.com/kitex-contrib/obs-opentelemetry/provider"
	"github.com/kitex-contrib/obs-opentelemetry/tracing"
	consul "github.com/kitex-contrib/registry-consul"
)

func main() {
	_ = godotenv.Load()
	mtl.InitMtl()
	dal.Init()
	opts := kitexInit()

	svr := paymentservice.NewServer(new(PaymentServiceImpl), opts...)

	err := svr.Run()
	if err != nil {
		klog.Error(err.Error())
	}
}

func kitexInit() (opts []server.Option) {
	// address
	address := conf.GetConf().Kitex.Address
	if strings.HasPrefix(address, ":") {
		localIp := utils.MustGetLocalIPv4()
		address = localIp + address
	}
	addr, err := net.ResolveTCPAddr("tcp", address)
	if err != nil {
		panic(err)
	}
	opts = append(opts, server.WithServiceAddr(addr))

	opts = append(opts,
		server.WithMiddleware(middleware.ServerMiddleware),
	)

	serviceName := conf.GetConf().Kitex.Service
	etcdClient, err := etcd.NewClient(etcd.Options{})
	if err != nil {
		panic(err)
	}

	opts = append(opts,
		server.WithSuite(etcdServer.NewSuite(serviceName, etcdClient)),
		server.WithMetaHandler(transmeta.ServerHTTP2Handler),
		server.WithServerBasicInfo(&rpcinfo.EndpointBasicInfo{ServiceName: serviceName}),
		server.WithTracer(prometheus.NewServerTracer("", "", prometheus.WithDisableServer(true), prometheus.WithRegistry(mtl.Registry))),
	)

	if os.Getenv("REGISTRY_ENABLE") == "true" {
		r, err := consul.NewConsulRegister(os.Getenv("REGISTRY_ADDR"))
		if err != nil {
			klog.Fatal(err)
		}
		opts = append(opts, server.WithRegistry(r))
	}
	_ = provider.NewOpenTelemetryProvider(
		provider.WithSdkTracerProvider(mtl.TracerProvider),
		provider.WithEnableMetrics(false),
	)
	opts = append(opts, server.WithSuite(tracing.NewServerSuite()))
	return
}
