package routes

import (
	"context"
	"fmt"
	"net/http"

	"github.com/baiyutang/gomall/app/frontend/infra/rpc"
	"github.com/baiyutang/gomall/app/frontend/kitex_gen/cart"
	"github.com/baiyutang/gomall/app/frontend/kitex_gen/order"
	"github.com/baiyutang/gomall/app/frontend/middleware"
	"github.com/baiyutang/gomall/app/frontend/types"
	frontendutils "github.com/baiyutang/gomall/app/frontend/utils"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/app/server"
	"github.com/cloudwego/hertz/pkg/common/utils"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
)

func RegisterOrder(h *server.Hertz) {
	g := h.Group("/order")
	g.Use(middleware.Auth())

	g.GET("/", func(ctx context.Context, c *app.RequestContext) {
		userId := frontendutils.GetUserIdFromCtx(ctx)
		var orders []*types.Order
		listOrderResp, err := rpc.OrderClient.ListOrder(ctx, &order.ListOrderRequest{UserId: userId})
		if err != nil {
			c.HTML(http.StatusOK, "error", map[string]any{"message": fmt.Sprintf("ListOrder.err:%v", err)})
			return
		}
		if listOrderResp != nil && len(listOrderResp.Orders) > 0 {
			for _, v := range listOrderResp.Orders {
				var items []types.OrderItem
				var total float32
				if len(v.OrderItems) > 0 {
					for _, vv := range v.OrderItems {
						total += vv.Cost
						i := vv.Item
						items = append(items, types.OrderItem{
							ProductId:   i.ProductId,
							Qty:         uint32(i.Quantity),
							ProductName: fmt.Sprintf("product - name %d", i.ProductId),
							Cost:        vv.Cost,
						})
					}
				}
				orders = append(orders, &types.Order{
					Cost:      total,
					Items:     items,
					CreatedAt: v.CreatedAt,
					OrderId:   v.OrderId,
					Consignee: types.Consignee{Email: v.Email},
				})
			}
		}

		var cartNum int
		cartResp, _ := rpc.CartClient.GetCart(ctx, &cart.GetCartRequest{UserId: userId})
		if cartResp != nil {
			cartNum = len(cartResp.Items)
		}

		c.HTML(consts.StatusOK, "order", frontendutils.WarpResponse(ctx, c, utils.H{
			"title":    "Order",
			"orders":   orders,
			"cart_num": cartNum,
		}))
	})
}
