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

package serversuite

import (
	"github.com/cloudwego/biz-demo/gomall/common/mtl"
	"github.com/cloudwego/kitex/pkg/rpcinfo"
	"github.com/cloudwego/kitex/pkg/transmeta"
	"github.com/cloudwego/kitex/server"
	"github.com/hertz-contrib/obs-opentelemetry/provider"
	prometheus "github.com/kitex-contrib/monitor-prometheus"
	"github.com/kitex-contrib/obs-opentelemetry/tracing"
)

type CommonServerSuite struct {
	CurrentServiceName string
}

func (s CommonServerSuite) Options() []server.Option {
	opts := []server.Option{
		server.WithMetaHandler(transmeta.ServerHTTP2Handler),
	}

	_ = provider.NewOpenTelemetryProvider(provider.WithSdkTracerProvider(mtl.TracerProvider), provider.WithEnableMetrics(false))

	opts = append(opts,
		server.WithServerBasicInfo(&rpcinfo.EndpointBasicInfo{
			ServiceName: s.CurrentServiceName,
		}),
		server.WithSuite(tracing.NewServerSuite()),
		server.WithTracer(prometheus.NewServerTracer("", "", prometheus.WithDisableServer(true), prometheus.WithRegistry(mtl.Registry))),
	)

	return opts
}
