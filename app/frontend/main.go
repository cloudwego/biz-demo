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
	h.GET("/cart", func(ctx context.Context, c *app.RequestContext) {
		c.HTML(consts.StatusOK, "cart", utils.H{
			"title": "Cart",
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
