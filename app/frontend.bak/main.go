package main

import (
	"context"
	"os"
	"time"

	"github.com/baiyutang/gomall/app/frontend/infra/mtl"
	"github.com/baiyutang/gomall/app/frontend/infra/rpc"
	"github.com/baiyutang/gomall/app/frontend/middleware"
	"github.com/baiyutang/gomall/app/frontend/routes"
	frontendutils "github.com/baiyutang/gomall/app/frontend/utils"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/app/server"
	"github.com/cloudwego/hertz/pkg/common/utils"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
	hertzprom "github.com/hertz-contrib/monitor-prometheus"
	hertzotelprovider "github.com/hertz-contrib/obs-opentelemetry/provider"
	hertzoteltracing "github.com/hertz-contrib/obs-opentelemetry/tracing"
	"github.com/hertz-contrib/sessions"
	"github.com/hertz-contrib/sessions/redis"
	"github.com/joho/godotenv"
	oteltrace "go.opentelemetry.io/otel/trace"
)

func main() {
	_ = godotenv.Load()

	mtl.InitMtl()
	rpc.InitClient()

	p := hertzotelprovider.NewOpenTelemetryProvider(
		hertzotelprovider.WithSdkTracerProvider(mtl.TracerProvider),
		hertzotelprovider.WithEnableMetrics(false),
	)
	defer p.Shutdown(context.Background())
	tracer, cfg := hertzoteltracing.NewServerTracer(hertzoteltracing.WithCustomResponseHandler(func(ctx context.Context, c *app.RequestContext) {
		c.Header("shop-trace-id", oteltrace.SpanFromContext(ctx).SpanContext().TraceID().String())
	}))
	h := server.Default(
		server.WithExitWaitTime(time.Second),
		server.WithDisablePrintRoute(false),
		server.WithTracer(
			hertzprom.NewServerTracer(
				"",
				"",
				hertzprom.WithRegistry(mtl.Registry),
				hertzprom.WithDisableServer(true),
			),
		),
		server.WithHostPorts(":8080"),
		tracer,
	)
	h.OnShutdown = append(h.OnShutdown, mtl.Hooks...)



	frontendutils.MustHandleError(err)

	middleware.RegisterMiddleware(h)

	h.Use(hertzoteltracing.ServerMiddleware(cfg))

	routes.RegisterProduct(h)
	routes.RegisterHome(h)
	routes.RegisterCategory(h)
	routes.RegisterAuth(h)
	routes.RegisterCart(h)
	routes.RegisterCheckout(h)
	routes.RegisterOrder(h)

	h.LoadHTMLGlob("template/*")
	h.Delims("{{", "}}")


	h.Static("/static", "./")
	h.Spin()
}
