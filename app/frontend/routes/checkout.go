package routes

import (
	"context"
	"fmt"
	"strconv"

	"github.com/baiyutang/gomall/app/frontend/infra/rpc"
	"github.com/baiyutang/gomall/app/frontend/kitex_gen/cart"
	"github.com/baiyutang/gomall/app/frontend/kitex_gen/checkout"
	"github.com/baiyutang/gomall/app/frontend/kitex_gen/payment"
	"github.com/baiyutang/gomall/app/frontend/kitex_gen/product"
	frontendutils "github.com/baiyutang/gomall/app/frontend/utils"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/app/server"
	"github.com/cloudwego/hertz/pkg/common/utils"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
	"github.com/cloudwego/kitex/pkg/klog"
)

func RegisterCheckout(h *server.Hertz) {
	type form struct {
		Firstname       string `json:"firstname" form:"firstname"`
		Lastname        string `json:"lastname" form:"lastname"`
		Street          string `json:"street" form:"street"`
		Zipcode         string `json:"zipcode" form:"city"`
		Province        string `json:"province" form:"province"`
		Country         string `json:"country" form:"country"`
		City            string `json:"city" form:"city"`
		CardNum         string `json:"cardNum" form:"cardNum"`
		ExpirationMonth int32  `json:"expirationMonth" form:"expirationMonth"`
		ExpirationYear  int32  `json:"expirationYear" form:"expirationYear"`
		Cvv             int32  `json:"cvv" form:"cvv"`
		Payment         string `json:"payment" form:"payment"`
	}
	h.POST("/checkout/waiting", func(ctx context.Context, c *app.RequestContext) {
		var f form
		err := c.BindAndValidate(&f)
		if err != nil {
			klog.Error(err)
		}
		r, err := rpc.CheckoutClient.Checkout(ctx, &checkout.CheckoutReq{
			UserId:    uint32(ctx.Value(frontendutils.UserIdKey).(float64)),
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
			fmt.Println(err)
		}
		fmt.Println(r)

		c.HTML(consts.StatusOK, "waiting", frontendutils.WarpResponse(ctx, c, utils.H{
			"title":    "waiting",
			"redirect": "/checkout/result",
		}))
	})

	h.GET("/checkout/result", func(ctx context.Context, c *app.RequestContext) {
		c.HTML(consts.StatusOK, "result", frontendutils.WarpResponse(ctx, c, utils.H{
			"title": "result",
		}))
	})

	h.GET("/checkout", func(ctx context.Context, c *app.RequestContext) {
		var items []map[string]string

		carts, err := rpc.CartClient.GetCart(ctx, &cart.GetCartRequest{UserId: uint32(ctx.Value(frontendutils.UserIdKey).(float64))})
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
			"cart_num": 10,
			"total":    strconv.FormatFloat(float64(total), 'f', 2, 64),
		}))
	})
}
