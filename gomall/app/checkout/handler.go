package main

import (
	"context"

	"github.com/cloudwego/biz-demo/gomall/app/checkout/biz/service"
	checkout "github.com/cloudwego/biz-demo/gomall/app/checkout/kitex_gen/checkout"
)

// CheckoutServiceImpl implements the last service interface defined in the IDL.
type CheckoutServiceImpl struct{}

// PlaceOrder implements the CheckoutServiceImpl interface.
func (s *CheckoutServiceImpl) PlaceOrder(ctx context.Context, req *checkout.CheckoutReq) (resp *checkout.CheckoutRes, err error) {
	resp, err = service.NewCheckoutService(ctx).Run(req)

	return resp, err
}

// Checkout implements the CheckoutServiceImpl interface.
func (s *CheckoutServiceImpl) Checkout(ctx context.Context, req *checkout.CheckoutReq) (resp *checkout.CheckoutRes, err error) {
	resp, err = service.NewCheckoutService(ctx).Run(req)

	return resp, err
}
