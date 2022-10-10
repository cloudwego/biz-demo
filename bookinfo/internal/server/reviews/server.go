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

package reviews

import (
	"context"
	"net"

	"github.com/cloudwego/biz-demo/bookinfo/kitex_gen/cwg/bookinfo/ratings/ratingservice"
	"github.com/cloudwego/biz-demo/bookinfo/kitex_gen/cwg/bookinfo/reviews"
	"github.com/cloudwego/biz-demo/bookinfo/kitex_gen/cwg/bookinfo/reviews/reviewsservice"
	"github.com/cloudwego/biz-demo/bookinfo/pkg/constants"
	"github.com/cloudwego/biz-demo/bookinfo/pkg/utils/logutils"
	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/cloudwego/kitex/server"
	kitexlogrus "github.com/kitex-contrib/obs-opentelemetry/logging/logrus"
	"github.com/kitex-contrib/obs-opentelemetry/provider"
	"github.com/kitex-contrib/obs-opentelemetry/tracing"
)

// Server reviews server
type Server struct {
	opts          *ServerOptions
	svc           reviews.ReviewsService
	ratingsClient ratingservice.Client
}

// ServerOptions server opts
type ServerOptions struct {
	Addr     string         `mapstructure:"addr"`
	LogLevel logutils.Level `mapstructure:"logLevel"`
}

// DefaultServerOptions default opts
func DefaultServerOptions() *ServerOptions {
	return &ServerOptions{
		Addr:     ":8082",
		LogLevel: logutils.LevelInfo,
	}
}

// Run reviews server
func (s *Server) Run(ctx context.Context) error {
	klog.SetLogger(kitexlogrus.NewLogger())
	klog.SetLevel(s.opts.LogLevel.KitexLogLevel())

	p := provider.NewOpenTelemetryProvider(
		provider.WithServiceName(constants.ReviewsServiceName),
		provider.WithInsecure(),
	)
	defer func(p provider.OtelProvider, ctx context.Context) {
		_ = p.Shutdown(ctx)
	}(p, ctx)

	addr, err := net.ResolveTCPAddr("tcp", s.opts.Addr)
	if err != nil {
		klog.Fatal(err)
	}
	svr := reviewsservice.NewServer(
		s.svc,
		server.WithServiceAddr(addr),
		server.WithSuite(tracing.NewServerSuite()),
	)
	if err := svr.Run(); err != nil {
		klog.Fatal(err)
	}

	return nil
}
