package service

import (
	"context"
	order "github.com/baiyutang/gomall/app/order/kitex_gen/order"
)

type ListOrderService struct {
	ctx context.Context
} // NewListOrderService new ListOrderService
func NewListOrderService(ctx context.Context) *ListOrderService {
	return &ListOrderService{ctx: ctx}
}

// Run create note info
func (s *ListOrderService) Run(req *order.ListOrderRequest) (resp *order.ListOrderResponse, err error) {
	// Finish your business logic.

	return
}
