package main

import (
	"context"
	cart "github.com/baiyutang/gomall/app/cart/kitex_gen/cart"
	"github.com/baiyutang/gomall/app/cart/biz/service"
)

// CartServiceImpl implements the last service interface defined in the IDL.
type CartServiceImpl struct{}

// AddItem implements the CartServiceImpl interface.
func (s *CartServiceImpl) AddItem(ctx context.Context, req *cart.AddItemRequest) (resp *cart.Empty, err error) {
	resp, err = service.NewAddItemService(ctx).Run(req)

	return resp, err
}

// GetCart implements the CartServiceImpl interface.
func (s *CartServiceImpl) GetCart(ctx context.Context, req *cart.GetCartRequest) (resp *cart.Cart, err error) {
	resp, err = service.NewGetCartService(ctx).Run(req)

	return resp, err
}

// EmptyCart implements the CartServiceImpl interface.
func (s *CartServiceImpl) EmptyCart(ctx context.Context, req *cart.EmptyCartRequest) (resp *cart.Empty, err error) {
	resp, err = service.NewEmptyCartService(ctx).Run(req)

	return resp, err
}
