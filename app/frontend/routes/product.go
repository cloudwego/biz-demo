package routes

import (
	"context"
	"github.com/baiyutang/gomall/app/frontend/infra/rpc"
	frontendutils "github.com/baiyutang/gomall/app/frontend/utils"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/app/server"
	"github.com/cloudwego/hertz/pkg/common/utils"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
	"strconv"

	"github.com/baiyutang/gomall/app/frontend/kitex_gen/product"
)

func RegisterProduct(h *server.Hertz) {

	productClient := rpc.ProductClient
	h.GET("/product", func(ctx context.Context, c *app.RequestContext) {
		productId := c.Query("id")
		id64, _ := strconv.ParseUint(productId, 10, 32)

		p, _ := productClient.GetProduct(context.Background(), &product.GetProductRequest{Id: uint32(id64)})
		c.HTML(consts.StatusOK, "product", frontendutils.WarpResponse(ctx, c, utils.H{
			"cart_num": 10,
			"item":     p,
		}))
	})
}
