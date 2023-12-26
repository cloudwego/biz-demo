package routes

import (
	"context"
	"strconv"

	"github.com/baiyutang/gomall/app/frontend/infra/rpc"
	"github.com/baiyutang/gomall/app/frontend/kitex_gen/cart"
	"github.com/baiyutang/gomall/app/frontend/kitex_gen/product"
	frontendutils "github.com/baiyutang/gomall/app/frontend/utils"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/app/server"
	"github.com/cloudwego/hertz/pkg/common/utils"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
)

func RegisterProduct(h *server.Hertz) {

	productClient := rpc.ProductClient
	h.GET("/product", func(ctx context.Context, c *app.RequestContext) {
		productId := c.Query("id")
		id64, _ := strconv.ParseUint(productId, 10, 32)

		p, _ := productClient.GetProduct(ctx, &product.GetProductRequest{Id: uint32(id64)})

		var cartNum int
		userId := uint32(ctx.Value(frontendutils.UserIdKey).(float64))
		cartResp, _ := rpc.CartClient.GetCart(ctx, &cart.GetCartRequest{UserId: userId})
		if cartResp != nil {
			cartNum = len(cartResp.Items)
		}

		c.HTML(consts.StatusOK, "product", frontendutils.WarpResponse(ctx, c, utils.H{
			"cart_num": cartNum,
			"item":     p,
		}))
	})
}
