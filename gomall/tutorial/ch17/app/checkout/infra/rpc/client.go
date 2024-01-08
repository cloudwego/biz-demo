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

package rpc

import (
	"os"
	"sync"

	"github.com/cloudwego/biz-demo/gomall/app/checkout/conf"
	checkoututils "github.com/cloudwego/biz-demo/gomall/app/checkout/utils"
	"github.com/cloudwego/biz-demo/gomall/common/clientsuite"
	"github.com/cloudwego/biz-demo/gomall/common/mtl"
	"github.com/cloudwego/biz-demo/gomall/rpc_gen/kitex_gen/cart/cartservice"
	"github.com/cloudwego/biz-demo/gomall/rpc_gen/kitex_gen/order/orderservice"
	"github.com/cloudwego/biz-demo/gomall/rpc_gen/kitex_gen/payment/paymentservice"
	"github.com/cloudwego/biz-demo/gomall/rpc_gen/kitex_gen/product/productcatalogservice"
	"github.com/cloudwego/kitex/client"
	"github.com/cloudwego/kitex/pkg/rpcinfo"
	"github.com/cloudwego/kitex/pkg/transmeta"
	"github.com/cloudwego/kitex/transport"
	consul "github.com/kitex-contrib/registry-consul"
)

var (
	CartClient    cartservice.Client
	ProductClient productcatalogservice.Client
	PaymentClient paymentservice.Client
	OrderClient   orderservice.Client

	once        sync.Once
	serviceName = conf.GetConf().Kitex.Service
)

var commonOpts []client.Option

func InitClient() {
	once.Do(func() {
		initCartClient()
		initProductClient()
		initPaymentClient()
		initOrderClient()
	})
}

func initProductClient() {
	var opts []client.Option
	r, err := consul.NewConsulResolver(conf.GetConf().Registry.RegistryAddress[0])
	checkoututils.MustHandleError(err)
	opts = append(opts, client.WithResolver(r))
	opts = append(opts,
		client.WithClientBasicInfo(&rpcinfo.EndpointBasicInfo{ServiceName: conf.GetConf().Kitex.Service}),
		client.WithTransportProtocol(transport.GRPC),
		client.WithMetaHandler(transmeta.ClientHTTP2Handler),
		client.WithSuite(clientsuite.CommonGrpcClientSuite{
			CurrentServiceName: serviceName,
			TracerProvider:     mtl.TracerProvider,
		}),
	)

	ProductClient, err = productcatalogservice.NewClient("product", opts...)
	checkoututils.MustHandleError(err)
}

func initCartClient() {
	var opts []client.Option
	r, err := consul.NewConsulResolver(os.Getenv("REGISTRY_ADDR"))
	checkoututils.MustHandleError(err)
	opts = append(opts, client.WithResolver(r), client.WithSuite(clientsuite.CommonGrpcClientSuite{
		CurrentServiceName: serviceName,
		TracerProvider:     mtl.TracerProvider,
	}))
	opts = append(opts,
		client.WithClientBasicInfo(&rpcinfo.EndpointBasicInfo{ServiceName: conf.GetConf().Kitex.Service}),
		client.WithTransportProtocol(transport.GRPC),
		client.WithMetaHandler(transmeta.ServerHTTP2Handler),
	)
	opts = append(opts, commonOpts...)
	CartClient, err = cartservice.NewClient("cart", opts...)
	checkoututils.MustHandleError(err)
}

func initPaymentClient() {
	var opts []client.Option
	r, err := consul.NewConsulResolver(os.Getenv("REGISTRY_ADDR"))
	checkoututils.MustHandleError(err)
	opts = append(opts, client.WithResolver(r), client.WithSuite(clientsuite.CommonGrpcClientSuite{
		CurrentServiceName: serviceName,
		TracerProvider:     mtl.TracerProvider,
	}))
	opts = append(opts, client.WithClientBasicInfo(&rpcinfo.EndpointBasicInfo{ServiceName: conf.GetConf().Kitex.Service}), client.WithTransportProtocol(transport.GRPC), client.WithMetaHandler(transmeta.ClientHTTP2Handler))
	opts = append(opts, commonOpts...)
	PaymentClient, err = paymentservice.NewClient("payment", opts...)
	checkoututils.MustHandleError(err)
}

func initOrderClient() {
	var opts []client.Option
	r, err := consul.NewConsulResolver(conf.GetConf().Registry.RegistryAddress[0])
	checkoututils.MustHandleError(err)
	opts = append(opts, client.WithResolver(r), client.WithSuite(clientsuite.CommonGrpcClientSuite{
		CurrentServiceName: serviceName,
		TracerProvider:     mtl.TracerProvider,
	}))
	opts = append(opts, client.WithClientBasicInfo(&rpcinfo.EndpointBasicInfo{ServiceName: conf.GetConf().Kitex.Service}), client.WithTransportProtocol(transport.GRPC), client.WithMetaHandler(transmeta.ClientHTTP2Handler))
	opts = append(opts, commonOpts...)
	OrderClient, err = orderservice.NewClient("order", opts...)
	checkoututils.MustHandleError(err)
}
