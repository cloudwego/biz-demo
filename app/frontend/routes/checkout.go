package routes

import (
	"context"
	frontendutils "github.com/baiyutang/gomall/app/frontend/utils"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/app/server"
	"github.com/cloudwego/hertz/pkg/common/utils"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
)

func RegisterCheckout(h *server.Hertz) {
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
		c.HTML(consts.StatusOK, "checkout", frontendutils.WarpResponse(ctx, c, utils.H{
			"title":    "Checkout",
			"items":    items,
			"cart_num": 10,
		}))
	}
}
