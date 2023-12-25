package routes

import (
	"context"
	"strconv"

	"github.com/baiyutang/gomall/app/cart/kitex_gen/cart"
	"github.com/baiyutang/gomall/app/frontend/infra/rpc"
	frontendutils "github.com/baiyutang/gomall/app/frontend/utils"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/app/server"
	"github.com/cloudwego/hertz/pkg/common/utils"
	"github.com/cloudwego/hertz/pkg/protocol/consts"

	"github.com/baiyutang/gomall/app/frontend/kitex_gen/product"
)

func RegisterProduct(h *server.Hertz) {

	productClient := rpc.ProductClient
	h.GET("/product", func(ctx context.Context, c *app.RequestContext) {
		productId := c.Query("id")
		id64, _ := strconv.ParseUint(productId, 10, 32)

		p, _ := productClient.GetProduct(ctx, &product.GetProductRequest{Id: uint32(id64)})

		var cartNum int
		cartResp, _ := rpc.CartClient.GetCart(ctx, &cart.GetCartRequest{})
		if cartResp != nil {
			cartNum = len(cartResp.Items)
		}

		c.HTML(consts.StatusOK, "product", frontendutils.WarpResponse(ctx, c, utils.H{
			"cart_num": cartNum,
			"item":     p,
		}))
	})
}
