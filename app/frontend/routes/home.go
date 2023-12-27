package routes

import (
	"context"
	"net/http"

	"github.com/baiyutang/gomall/app/frontend/infra/rpc"
	"github.com/baiyutang/gomall/app/frontend/kitex_gen/cart"
	"github.com/baiyutang/gomall/app/frontend/kitex_gen/product"
	frontendutils "github.com/baiyutang/gomall/app/frontend/utils"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/app/server"
	"github.com/cloudwego/hertz/pkg/common/utils"
	"github.com/cloudwego/kitex/pkg/klog"
)

func RegisterHome(h *server.Hertz) {
	productClient := rpc.ProductClient
	h.GET("/", func(ctx context.Context, c *app.RequestContext) {
		p, err := productClient.ListProducts(ctx, &product.ListProductsReq{})
		if err != nil {
			klog.Error(err)
		}
		var items []*product.Product
		if p != nil {
			items = p.Products
		}
		var cartNum int
		tryUserId := ctx.Value(frontendutils.UserIdKey)
		if tryUserId != nil {
			userId := uint32(tryUserId.(float64))
			cartResp, _ := rpc.CartClient.GetCart(ctx, &cart.GetCartRequest{UserId: userId})
			if cartResp != nil {
				cartNum = len(cartResp.Items)
			}
		}
		c.HTML(http.StatusOK, "home", frontendutils.WarpResponse(ctx, c, utils.H{
			"title":    "Hot sale",
			"cart_num": cartNum,
			"items":    items,
		}))
	})
}
