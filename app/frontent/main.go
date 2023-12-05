package main

import (
	"context"
	"fmt"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/app/server"
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"github.com/cloudwego/hertz/pkg/common/utils"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
	"net/http"
	"time"
)

func main() {
	h := server.Default(server.WithDisablePrintRoute(false), server.WithAutoReloadRender(true, time.Second))
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
			"title":    "hello cloudwego",
			"cart_num": 10,
			"items":    items,
		})
	})
	h.GET("/category", func(ctx context.Context, c *app.RequestContext) {
		c.HTML(consts.StatusOK, "category", utils.H{
			"title": "category",
		})
	})
	h.GET("/product", func(ctx context.Context, c *app.RequestContext) {
		c.HTML(consts.StatusOK, "product", utils.H{
			"title": "product",
		})
	})
	h.GET("/cart", func(ctx context.Context, c *app.RequestContext) {
		c.HTML(consts.StatusOK, "cart", utils.H{
			"title": "cart",
		})
	})
	h.GET("/order", func(ctx context.Context, c *app.RequestContext) {
		c.HTML(consts.StatusOK, "order", utils.H{
			"title": "order",
		})
	})
	h.GET("/about", func(ctx context.Context, c *app.RequestContext) {
		c.HTML(consts.StatusOK, "about", utils.H{
			"title": "about",
		})

	})
	h.Static("/static", "./")
	h.Spin()
}
