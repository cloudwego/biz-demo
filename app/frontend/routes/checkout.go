package routes

import (
	"context"
	"strconv"

	"github.com/baiyutang/gomall/app/frontend/infra/rpc"
	"github.com/baiyutang/gomall/app/frontend/kitex_gen/cart"
	"github.com/baiyutang/gomall/app/frontend/kitex_gen/checkout"
	"github.com/baiyutang/gomall/app/frontend/kitex_gen/payment"
	"github.com/baiyutang/gomall/app/frontend/kitex_gen/product"
	"github.com/baiyutang/gomall/app/frontend/middleware"
	"github.com/baiyutang/gomall/app/frontend/types"
	frontendutils "github.com/baiyutang/gomall/app/frontend/utils"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/app/server"
	"github.com/cloudwego/hertz/pkg/common/utils"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
	"github.com/cloudwego/kitex/pkg/klog"
)

func RegisterCheckout(h *server.Hertz) {
	h.POST("/checkout/waiting", middleware.Auth(), func(ctx context.Context, c *app.RequestContext) {
		var f types.CheckoutForm

		err := c.BindAndValidate(&f)
		if err != nil {
			klog.Error(err)
		}
		userId := frontendutils.GetUserIdFromCtx(ctx)

		_, err = rpc.CheckoutClient.Checkout(ctx, &checkout.CheckoutReq{
			UserId:    userId,
			Email:     f.Email,
			Firstname: f.Firstname,
			Lastname:  f.Lastname,
			Address: &checkout.Address{
				Country:       f.Country,
				ZipCode:       f.Zipcode,
				City:          f.City,
				State:         f.Province,
				StreetAddress: f.Street,
			},
			CreditCard: &payment.CreditCardInfo{
				CreditCardNumber:          f.CardNum,
				CreditCardExpirationYear:  f.ExpirationYear,
				CreditCardExpirationMonth: f.ExpirationMonth,
				CreditCardCvv:             f.Cvv,
			},
		})
		if err != nil {
			c.HTML(consts.StatusOK, "error", map[string]interface{}{"message": err})
			return
		}

		c.HTML(consts.StatusOK, "waiting", frontendutils.WarpResponse(ctx, c, utils.H{
			"title":    "waiting",
			"redirect": "/checkout/result",
		}))
	})

	h.GET("/checkout/result", middleware.Auth(), func(ctx context.Context, c *app.RequestContext) {
		c.HTML(consts.StatusOK, "result", frontendutils.WarpResponse(ctx, c, utils.H{
			"title": "result",
		}))
	})

	h.GET("/checkout", middleware.Auth(), func(ctx context.Context, c *app.RequestContext) {
		var items []map[string]string
		userId := frontendutils.GetUserIdFromCtx(ctx)

		carts, err := rpc.CartClient.GetCart(ctx, &cart.GetCartRequest{UserId: userId})
		if err != nil {
		}
		var total float32
		for _, v := range carts.Items {
			p, err := rpc.ProductClient.GetProduct(ctx, &product.GetProductRequest{Id: v.ProductId})
			if err != nil {
				continue
			}
			items = append(items, map[string]string{"Name": p.Name, "Price": strconv.FormatFloat(float64(p.Price), 'f', 2, 64), "Picture": p.Picture, "Qty": strconv.Itoa(int(v.Quantity))})
			total += float32(v.Quantity) * p.Price
		}

		c.HTML(consts.StatusOK, "checkout", frontendutils.WarpResponse(ctx, c, utils.H{
			"title":    "Checkout",
			"items":    items,
			"cart_num": len(items),
			"total":    strconv.FormatFloat(float64(total), 'f', 2, 64),
		}))
	})
}
