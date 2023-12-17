package service

import (
	"context"
	checkout "github.com/baiyutang/gomall/app/checkout/kitex_gen/checkout"
)

type PlaceOrderService struct {
	ctx context.Context
} // NewPlaceOrderService new PlaceOrderService
func NewPlaceOrderService(ctx context.Context) *PlaceOrderService {
	return &PlaceOrderService{ctx: ctx}
}

// Run create note info
func (s *PlaceOrderService) Run(req *checkout.PlaceOrderReq) (resp *checkout.PlaceOrderRes, err error) {
	// Finish your business logic.

	return
}
