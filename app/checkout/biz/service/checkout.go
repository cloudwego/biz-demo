package service

import (
	"context"
	checkout "github.com/baiyutang/gomall/app/checkout/kitex_gen/checkout"
)

type CheckoutService struct {
	ctx context.Context
} // NewCheckoutService new CheckoutService
func NewCheckoutService(ctx context.Context) *CheckoutService {
	return &CheckoutService{ctx: ctx}
}

// Run create note info
func (s *CheckoutService) Run(req *checkout.CheckoutReq) (resp *checkout.CheckoutRes, err error) {
	// Finish your business logic.

	return
}
