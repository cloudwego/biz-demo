package service

import (
	"context"

	"github.com/baiyutang/gomall/app/cart/biz/dal/mysql"
	"github.com/baiyutang/gomall/app/cart/biz/model"
	cart "github.com/baiyutang/gomall/app/cart/kitex_gen/cart"
)

type GetCartService struct {
	ctx context.Context
} // NewGetCartService new GetCartService
func NewGetCartService(ctx context.Context) *GetCartService {
	return &GetCartService{ctx: ctx}
}

// Run create note info
func (s *GetCartService) Run(req *cart.GetCartRequest) (resp *cart.Cart, err error) {
	// resp = &cart.Cart{}
	// Finish your business logic.
	carts, err := model.GetCartByUserId(mysql.DB, s.ctx, req.GetUserId())
	var items []*cart.CartItem
	for _, v := range carts {
		items = append(items, &cart.CartItem{ProductId: v.ProductId, Quantity: int32(v.Qty)})
	}

	return &cart.Cart{UserId: req.GetUserId(), Items: items}, nil
}
