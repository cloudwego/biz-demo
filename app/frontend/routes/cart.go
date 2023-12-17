package routes

import (
	"context"
	frontendutils "github.com/baiyutang/gomall/app/frontend/utils"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/app/server"
	"github.com/cloudwego/hertz/pkg/common/utils"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
)

func RegisterCart(h *server.Hertz) {
	h.GET("/cart", func(ctx context.Context, c *app.RequestContext) {
		var items []string
		for i := 0; i <= 10; i++ {
			items = append(items, "hello")
		}
		c.HTML(consts.StatusOK, "cart", frontendutils.WarpResponse(ctx, c, utils.H{
			"title":    "Cart",
			"items":    items,
			"cart_num": 10,
		}))
	})
	h.POST("/cart", func(ctx context.Context, c *app.RequestContext) {
		c.Redirect(consts.StatusFound, []byte("/cart"))
	})
}
