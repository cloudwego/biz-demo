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

package productpage

import (
	"context"

	"github.com/cloudwego/biz-demo/bookinfo/internal/handler/productpage"
	"github.com/cloudwego/biz-demo/bookinfo/pkg/constants"
	"github.com/cloudwego/biz-demo/bookinfo/pkg/utils/logutils"
	hertzserver "github.com/cloudwego/hertz/pkg/app/server"
	"github.com/cloudwego/kitex/pkg/klog"
	hertztracing "github.com/hertz-contrib/obs-opentelemetry/tracing"
	kitexlogrus "github.com/kitex-contrib/obs-opentelemetry/logging/logrus"
	"github.com/kitex-contrib/obs-opentelemetry/provider"
)

// Server hertz server
type Server struct {
	opts    *ServerOptions
	handler *productpage.Handler
}

// ServerOptions server opts
type ServerOptions struct {
	Addr     string         `mapstructure:"addr"`
	LogLevel logutils.Level `mapstructure:"logLevel"`
}

// DefaultServerOptions default opts
func DefaultServerOptions() *ServerOptions {
	return &ServerOptions{
		Addr:     ":8081",
		LogLevel: logutils.LevelInfo,
	}
}

// Run hertz server
func (s *Server) Run(ctx context.Context) error {
	klog.SetLogger(kitexlogrus.NewLogger())
	klog.SetLevel(s.opts.LogLevel.KitexLogLevel())

	p := provider.NewOpenTelemetryProvider(
		provider.WithServiceName(constants.ProductpageServiceName),
		provider.WithInsecure(),
	)

	tracer, cfg := hertztracing.NewServerTracer()
	h := hertzserver.Default(
		tracer,
		hertzserver.WithHostPorts(s.opts.Addr),
	)
	h.Use(hertztracing.ServerMiddleware(cfg))

	h.GET("/api/v1/products/:productID", s.handler.GetProduct)

	h.Spin()

	return p.Shutdown(ctx)
}
