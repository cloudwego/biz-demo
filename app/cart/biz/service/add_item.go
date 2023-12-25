package service

import (
	"context"

	"github.com/baiyutang/gomall/app/cart/biz/dal/mysql"
	"github.com/baiyutang/gomall/app/cart/biz/model"
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
	err = model.Add(mysql.DB, s.ctx, &model.CartItem{
		UserId:    req.UserId,
		ProductId: req.Item.ProductId,
		Quantity:  req.Item.Quantity,
	})
	return
}
