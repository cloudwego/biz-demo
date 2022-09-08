// Copyright 2022 CloudWeGo Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//

package injectors

import (
	"github.com/cloudwego/biz-demo/bookinfo/kitex_gen/cwg/bookinfo/details/detailsservice"
	"github.com/cloudwego/biz-demo/bookinfo/pkg/constants"
	kclient "github.com/cloudwego/kitex/client"
	"github.com/cloudwego/kitex/pkg/xds"
	"github.com/kitex-contrib/obs-opentelemetry/tracing"
	"github.com/kitex-contrib/xds/xdssuite"
)

type DetailsClientOptions struct {
	Endpoint  string `mapstructure:"endpoint"`
	EnableXDS bool   `mapstructure:"enableXDS"`
}

func DefaultDetailsClientOptions() *DetailsClientOptions {
	return &DetailsClientOptions{
		Endpoint:  ":8084",
		EnableXDS: false,
	}
}

func ProvideDetailsClient(opts *DetailsClientOptions) (detailsservice.Client, error) {
	if opts.EnableXDS {
		return detailsservice.NewClient(
			constants.DetailsServiceName,
			kclient.WithHostPorts(opts.Endpoint),
			kclient.WithSuite(tracing.NewClientSuite()),
			kclient.WithXDSSuite(xds.ClientSuite{
				RouterMiddleware: xdssuite.NewXDSRouterMiddleware(),
				Resolver:         xdssuite.NewXDSResolver(),
			}),
		)
	}

	return detailsservice.NewClient(
		constants.DetailsServiceName,
		kclient.WithHostPorts(opts.Endpoint),
		kclient.WithSuite(tracing.NewClientSuite()),
	)
}
