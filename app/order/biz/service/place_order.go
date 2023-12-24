package service

import (
	"context"
	order "github.com/baiyutang/gomall/app/order/kitex_gen/order"
	"github.com/google/uuid"
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
	u, _ := uuid.NewRandom()

	return &order.PlaceOrderResponse{Order: &order.OrderResult{OrderId: u.String()}}, nil
}
