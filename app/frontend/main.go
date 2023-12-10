package main

import (
	"context"
	"github.com/joho/godotenv"
	"os"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/app/server"
	"github.com/cloudwego/hertz/pkg/common/utils"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
	hertzprom "github.com/hertz-contrib/monitor-prometheus"
	hertzotelprovider "github.com/hertz-contrib/obs-opentelemetry/provider"
	hertzoteltracing "github.com/hertz-contrib/obs-opentelemetry/tracing"

	"github.com/baiyutang/gomall/app/frontend/infra/mtl"
	"github.com/baiyutang/gomall/app/frontend/routes"
)

func main() {
	_ = godotenv.Load()

	mtl.InitMtl()

	p := hertzotelprovider.NewOpenTelemetryProvider(
		hertzotelprovider.WithSdkTracerProvider(mtl.TracerProvider),
		hertzotelprovider.WithEnableMetrics(false),
	)
	defer p.Shutdown(context.Background())
	tracer, cfg := hertzoteltracing.NewServerTracer()
	h := server.Default(
		server.WithDisablePrintRoute(false),
		server.WithTracer(
			hertzprom.NewServerTracer(
				":10086",
				"/metrics",
				hertzprom.WithRegistry(mtl.Registry),
				//hertzprom.WithDisableServer(true),
			),
		),
		tracer,
	)
	h.Use(hertzoteltracing.ServerMiddleware(cfg))
	h.LoadHTMLGlob("template/*")
	h.Delims("{{", "}}")
	routes.RegisterProduct(h)
	routes.RegisterHome(h)
	routes.RegisterCategory(h)
	h.GET("sign-in", func(ctx context.Context, c *app.RequestContext) {
		c.HTML(consts.StatusOK, "sign-in", utils.H{
			"title": "Sign in",
		})

	})
	h.GET("sign-up", func(ctx context.Context, c *app.RequestContext) {
		c.HTML(consts.StatusOK, "sign-up", utils.H{
			"title": "Sign up",
		})
	})
	h.GET("/cart", func(ctx context.Context, c *app.RequestContext) {
		var items []string
		for i := 1; i <= 10; i++ {
			items = append(items, "hello")
		}
		c.HTML(consts.StatusOK, "cart", utils.H{
			"title":    "Cart",
			"items":    items,
			"cart_num": 10,
		})
	})
	h.POST("cart", func(ctx context.Context, c *app.RequestContext) {
		c.Redirect(consts.StatusFound, []byte("/cart"))
	})
	h.GET("/checkout/waiting", func(ctx context.Context, c *app.RequestContext) {
		c.HTML(consts.StatusOK, "waiting", utils.H{
			"title":    "waiting",
			"redirect": "/checkout/result",
		})
	})

	h.GET("/checkout/result", func(ctx context.Context, c *app.RequestContext) {
		c.HTML(consts.StatusOK, "result", utils.H{
			"title": "result",
		})
	})

	h.GET("/checkout", func(ctx context.Context, c *app.RequestContext) {
		var items []string
		for i := 1; i <= 10; i++ {
			items = append(items, "hello")
		}
		c.HTML(consts.StatusOK, "checkout", utils.H{
			"title":    "Checkout",
			"items":    items,
			"cart_num": 10,
		})

	})
	h.GET("/order", func(ctx context.Context, c *app.RequestContext) {
		c.HTML(consts.StatusOK, "order", utils.H{
			"title": "Order",
		})
	})
	h.GET("/about", func(ctx context.Context, c *app.RequestContext) {
		c.HTML(consts.StatusOK, "about", utils.H{
			"title": "About",
		})
	})
	if os.Getenv("GO_ENV") != "online" {
		h.GET("/robots.txt", func(ctx context.Context, c *app.RequestContext) {
			c.Data(consts.StatusOK, "text/plain", []byte(`User-agent: *
Disallow: /`))
		})
	}

	h.Static("/static", "./")
	h.Spin()
}
