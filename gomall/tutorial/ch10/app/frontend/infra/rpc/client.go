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
	"context"
	"sync"

	"github.com/cloudwego/biz-demo/gomall/app/frontend/conf"
	frontendutils "github.com/cloudwego/biz-demo/gomall/app/frontend/utils"
	"github.com/cloudwego/biz-demo/gomall/rpc_gen/kitex_gen/product"
	"github.com/cloudwego/biz-demo/gomall/rpc_gen/kitex_gen/product/productcatalogservice"
	"github.com/cloudwego/biz-demo/gomall/rpc_gen/kitex_gen/user/userservice"
	"github.com/cloudwego/kitex/client"
	"github.com/cloudwego/kitex/pkg/circuitbreak"
	"github.com/cloudwego/kitex/pkg/fallback"
	"github.com/cloudwego/kitex/pkg/rpcinfo"
	consul "github.com/kitex-contrib/registry-consul"
)

var (
	UserClient    userservice.Client
	ProductClient productcatalogservice.Client

	once sync.Once
)

func InitClient() {
	once.Do(func() {
		initUserClient()
		initProductClient()
	})
}

func initUserClient() {
	var opts []client.Option
	r, err := consul.NewConsulResolver(conf.GetConf().Hertz.RegistryAddr)
	frontendutils.MustHandleError(err)
	opts = append(opts, client.WithResolver(r))

	UserClient, err = userservice.NewClient("user", opts...)
	frontendutils.MustHandleError(err)
}

func initProductClient() {
	var opts []client.Option
	r, err := consul.NewConsulResolver(conf.GetConf().Hertz.RegistryAddr)
	frontendutils.MustHandleError(err)
	opts = append(opts, client.WithResolver(r))

	cbs := circuitbreak.NewCBSuite(func(ri rpcinfo.RPCInfo) string {
		return circuitbreak.RPCInfo2Key(ri)
	})

	cbs.UpdateServiceCBConfig("shop-frontend/product/GetProduct", circuitbreak.CBConfig{Enable: true, ErrRate: 0.5, MinSample: 2})

	opts = append(opts, client.WithCircuitBreaker(cbs), client.WithFallback(fallback.NewFallbackPolicy(fallback.UnwrapHelper(func(ctx context.Context, req, resp interface{}, err error) (fbResp interface{}, fbErr error) {
		methodName := rpcinfo.GetRPCInfo(ctx).To().Method()
		if err == nil {
			return resp, err
		}
		if methodName != "ListProducts" {
			return resp, err
		}
		return &product.ListProductsResp{
			Products: []*product.Product{
				{
					Price:       6.6,
					Id:          3,
					Picture:     "/static/image/t-shirt.jpeg",
					Name:        "T-Shirt",
					Description: "CloudWeGo T-Shirt",
				},
			},
		}, nil
	}))))
	ProductClient, err = productcatalogservice.NewClient("product", opts...)
	frontendutils.MustHandleError(err)
}
