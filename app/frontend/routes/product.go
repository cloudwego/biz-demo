package routes

import (
	"context"
	frontendutils "github.com/baiyutang/gomall/app/frontend/utils"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/app/server"
	"github.com/cloudwego/hertz/pkg/common/utils"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
	"github.com/cloudwego/kitex/client"
	"strconv"

	"github.com/baiyutang/gomall/app/frontend/kitex_gen/product"
	"github.com/baiyutang/gomall/app/frontend/kitex_gen/product/productcatalogservice"
)

func RegisterProduct(h *server.Hertz) {
	productClient, _ := productcatalogservice.NewClient("product", client.WithHostPorts("localhost:8881"))
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
