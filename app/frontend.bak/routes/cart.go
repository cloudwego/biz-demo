package routes

import (
	"context"
	"fmt"
	"net/http"
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
