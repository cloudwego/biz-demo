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

package mtl

import (
	"context"

	"github.com/cloudwego/biz-demo/gomall/app/email/conf"
	"github.com/cloudwego/kitex/server"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/exporters/otlp/otlptrace/otlptracegrpc"
	"go.opentelemetry.io/otel/sdk/resource"
	tracesdk "go.opentelemetry.io/otel/sdk/trace"
	semconv "go.opentelemetry.io/otel/semconv/v1.17.0"
)

var TracerProvider *tracesdk.TracerProvider

func initTracing() {
	exporter, err := otlptracegrpc.New(context.Background())
	if err != nil {
		panic(err)
	}
	server.RegisterShutdownHook(func() {
		_ = exporter.Shutdown(context.Background())
	})
	processor := tracesdk.NewBatchSpanProcessor(exporter)
	res, err := resource.New(context.Background(), resource.WithAttributes(semconv.ServiceNameKey.String(conf.GetConf().Kitex.Service)))
	if err != nil {
		res = resource.Default()
	}
	TracerProvider = tracesdk.NewTracerProvider(tracesdk.WithSpanProcessor(processor), tracesdk.WithResource(res))
	server.RegisterShutdownHook(func() {
		_ = TracerProvider.Shutdown(context.Background())
	})
	otel.SetTracerProvider(TracerProvider)
}
