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
	// Finish your business logic.
	err = model.Empty(mysql.DB, s.ctx, req.UserId)
	return
}
