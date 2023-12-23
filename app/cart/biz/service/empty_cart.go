package service

import (
	"context"
	"github.com/baiyutang/gomall/app/cart/biz/dal/mysql"
	"github.com/baiyutang/gomall/app/cart/biz/model"
	cart "github.com/baiyutang/gomall/app/cart/kitex_gen/cart"
	"github.com/cloudwego/kitex/pkg/kerrors"
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
	err = model.EmptyCart(mysql.DB, s.ctx, req.GetUserId())
	if err != nil {
		return &cart.Empty{}, kerrors.NewBizStatusError(50001, "empty cart error")
	}

	return &cart.Empty{}, nil
}
