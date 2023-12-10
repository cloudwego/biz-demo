package service

import (
	"context"
	cart "github.com/baiyutang/gomall/app/cart/kitex_gen/cart"
)

type AddItemService struct {
	ctx context.Context
} // NewAddItemService new AddItemService
func NewAddItemService(ctx context.Context) *AddItemService {
	return &AddItemService{ctx: ctx}
}

// Run create note info
func (s *AddItemService) Run(req *cart.AddItemRequest) (resp *cart.Empty, err error) {
	// Finish your business logic.

	return
}
