package main

import (
	"context"
	"os"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/app/server"
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"github.com/cloudwego/hertz/pkg/common/utils"
	"github.com/cloudwego/hertz/pkg/protocol/consts"

	"github.com/baiyutang/gomall/app/frontend/routes"
)

func main() {
	h := server.Default(server.WithDisablePrintRoute(false))
	hlog.SetLevel(hlog.LevelTrace)
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
