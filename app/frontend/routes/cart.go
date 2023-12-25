package routes

import (
	"context"
	"fmt"
	"net/http"

	"github.com/baiyutang/gomall/app/cart/kitex_gen/cart"
	"github.com/baiyutang/gomall/app/frontend/infra/rpc"
	"github.com/baiyutang/gomall/app/frontend/types"
	frontendutils "github.com/baiyutang/gomall/app/frontend/utils"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/app/server"
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"github.com/cloudwego/hertz/pkg/common/utils"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
)

func RegisterCart(h *server.Hertz) {
	cartSvc := rpc.CartClient
	h.GET("/cart", func(ctx context.Context, c *app.RequestContext) {
		userId := ctx.Value(frontendutils.UserIdKey).(int32)
		resp, err := cartSvc.GetCart(ctx, &cart.GetCartRequest{UserId: uint32(userId)})
		if err != nil {
			c.HTML(http.StatusOK, "error", map[string]interface{}{"message": err})
		}
		var (
			items []*types.CartItem
			// productIdList []uint32
		)
		itemsMap := make(map[uint32]*types.CartItem, 0)

		if resp == nil {
			for _, v := range resp.Items {
				if v.ProductId == 0 {
					continue
				}
				// productIdList = append(productIdList, v.ProductId)
				itemsMap[v.ProductId] = &types.CartItem{
					ProductId:   v.ProductId,
					Quantity:    v.Quantity,
					ProductName: fmt.Sprintf("producnt %d", v.ProductId),
				}
			}
		}
		// rpc.ProductClient.ListProducts(ctx, &product.ListProductsReq{})

		c.HTML(consts.StatusOK, "cart", frontendutils.WarpResponse(ctx, c, utils.H{
			"title":    "Cart",
			"items":    items,
			"cart_num": len(items),
		}))
	})

	h.POST("/cart", func(ctx context.Context, c *app.RequestContext) {
		req := &types.AddCartReq{}
		_ = c.BindByContentType(req)

		hlog.Info("req---->", req)
		userId := ctx.Value(frontendutils.UserIdKey).(int32)
		_, err := cartSvc.AddItem(ctx, &cart.AddItemRequest{
			UserId: uint32(userId),
			Item:   &cart.CartItem{ProductId: req.ProductId, Quantity: int32(req.Quantity)},
		})
		if err != nil {
			c.HTML(http.StatusOK, "error", map[string]interface{}{"message": err})
		}
		c.Redirect(consts.StatusFound, []byte("/cart"))
	})
}
