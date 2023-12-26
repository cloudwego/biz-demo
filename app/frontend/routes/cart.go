package routes

import (
	"context"
	"fmt"
	"strconv"

	"github.com/baiyutang/gomall/app/frontend/kitex_gen/product"
	"github.com/baiyutang/gomall/app/frontend/middleware"

	"github.com/baiyutang/gomall/app/frontend/infra/rpc"
	"github.com/baiyutang/gomall/app/frontend/kitex_gen/cart"
	frontendutils "github.com/baiyutang/gomall/app/frontend/utils"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/app/server"
	"github.com/cloudwego/hertz/pkg/common/utils"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
)

func RegisterCart(h *server.Hertz) {
	g := h.Group("/cart")
	g.Use(middleware.Auth())
	g.GET("/", func(ctx context.Context, c *app.RequestContext) {
		var items []map[string]string
		carts, err := rpc.CartClient.GetCart(ctx, &cart.GetCartRequest{UserId: uint32(ctx.Value(frontendutils.UserIdKey).(float64))})
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
		c.HTML(consts.StatusOK, "cart", frontendutils.WarpResponse(ctx, c, utils.H{
			"title":    "Cart",
			"items":    items,
			"total":    total,
			"cart_num": len(items),
		}))
	})
	type form struct {
		ProductId  uint32 `json:"productId" form:"productId"`
		ProductNum uint32 `json:"productNum" form:"productNum"`
	}
	g.POST("/", func(ctx context.Context, c *app.RequestContext) {
		var f form
		c.BindAndValidate(&f)
		_, err := rpc.CartClient.AddItem(ctx, &cart.AddItemRequest{UserId: uint32(ctx.Value(frontendutils.UserIdKey).(float64)), Item: &cart.CartItem{
			ProductId: f.ProductId,
			Quantity:  int32(f.ProductNum),
		}})
		fmt.Println(err)
		c.Redirect(consts.StatusFound, []byte("/cart"))
	})
}
