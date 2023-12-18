package service

import (
	"context"
	order "github.com/baiyutang/gomall/app/order/kitex_gen/order"
)

type PlaceOrderService struct {
	ctx context.Context
} // NewPlaceOrderService new PlaceOrderService
func NewPlaceOrderService(ctx context.Context) *PlaceOrderService {
	return &PlaceOrderService{ctx: ctx}
}

// Run create note info
func (s *PlaceOrderService) Run(req *order.PlaceOrderRequest) (resp *order.PlaceOrderResponse, err error) {
	// Finish your business logic.

	return
}
