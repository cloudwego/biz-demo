package routes

import (
	"context"
	"fmt"
	"strconv"

	"github.com/baiyutang/gomall/app/frontend/infra/rpc"
	"github.com/baiyutang/gomall/app/frontend/kitex_gen/cart"
	"github.com/baiyutang/gomall/app/frontend/kitex_gen/order"
	"github.com/baiyutang/gomall/app/frontend/kitex_gen/product"
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
		var items []map[string]string
		userId := uint32(ctx.Value(frontendutils.UserIdKey).(float64))
		carts, err := rpc.CartClient.GetCart(ctx, &cart.GetCartRequest{UserId: userId})
		if err != nil {
			c.JSON(500, "get cart error")
		}
		var total float32
		for _, v := range carts.Items {
			p, err := rpc.ProductClient.GetProduct(ctx, &product.GetProductRequest{Id: v.GetProductId()})
			if err != nil {
				continue
			}
			items = append(items, map[string]string{"Name": p.Name, "Description": p.Description, "Picture": p.Picture, "Price": strconv.FormatFloat(float64(p.Price), 'f', 2, 64), "Qty": strconv.Itoa(int(v.Quantity))})
			total += float32(v.Quantity) * p.Price
		}
		var orders []*types.Order
		listOrderResp, err := rpc.OrderClient.ListOrder(ctx, &order.ListOrderRequest{UserId: userId})
		if err != nil {
			c.JSON(500, "get cart error")
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
							ProductName: fmt.Sprintf("product %d", i.ProductId),
							OrderId:     v.OrderId,
							Cost:        vv.Cost,
						})
					}
				}
				orders = append(orders, &types.Order{Cost: total, Items: items, CreatedAt: v.CreatedAt})
			}
		}

		c.HTML(consts.StatusOK, "order", frontendutils.WarpResponse(ctx, c, utils.H{
			"title":    "Order",
			"orders":   orders,
			"total":    total,
			"cart_num": len(items),
		}))
	})

	g.POST("/", func(ctx context.Context, c *app.RequestContext) {

	})
}
