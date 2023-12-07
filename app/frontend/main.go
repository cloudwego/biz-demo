package main

import (
	"context"
	"fmt"
	"net/http"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/app/server"
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"github.com/cloudwego/hertz/pkg/common/utils"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
)

func main() {
	h := server.Default(server.WithDisablePrintRoute(false))
	hlog.SetLevel(hlog.LevelTrace)
	h.LoadHTMLGlob("template/*")
	h.Delims("{{", "}}")
	h.GET("/", func(ctx context.Context, c *app.RequestContext) {
		var items []map[string]string
		for i := 1; i <= 10; i++ {
			items = append(items, map[string]string{
				"title":       fmt.Sprintf("product%d", i),
				"description": fmt.Sprintf("product description %d", i),
			})
		}
		c.HTML(http.StatusOK, "home", utils.H{
			"title":    "Hot sale",
			"cart_num": 10,
			"items":    items,
		})
	})
	h.GET("/category", func(ctx context.Context, c *app.RequestContext) {
		c.HTML(consts.StatusOK, "category", utils.H{
			"title": "Category",
		})
	})
	routes.RegisterProductRoute(h)
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
	h.Static("/static", "./")
	h.Spin()
}
