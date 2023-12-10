package service

import (
	"context"
	cart "github.com/baiyutang/gomall/app/cart/kitex_gen/cart"
)

type EmptyCartService struct {
	ctx context.Context
} // NewEmptyCartService new EmptyCartService
func NewEmptyCartService(ctx context.Context) *EmptyCartService {
	return &EmptyCartService{ctx: ctx}
}

// Run create note info
func (s *EmptyCartService) Run(req *cart.EmptyCartRequest) (resp *cart.Empty, err error) {
	// Finish your business logic.

	return
}
