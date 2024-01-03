package service

import (
	"context"
	"fmt"

	"github.com/cloudwego/biz-demo/gomall/app/cart/biz/dal/mysql"
	"github.com/cloudwego/biz-demo/gomall/app/cart/biz/model"
	"github.com/cloudwego/biz-demo/gomall/app/cart/infra/rpc"

	cart "github.com/cloudwego/biz-demo/gomall/app/cart/kitex_gen/cart"
	"github.com/cloudwego/biz-demo/gomall/app/cart/kitex_gen/product"
	"github.com/cloudwego/kitex/pkg/kerrors"
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
	getProduct, err := rpc.ProductClient.GetProduct(s.ctx, &product.GetProductRequest{Id: req.Item.GetProductId()})
	if err != nil {
		return nil, err
	}
	fmt.Printf("%#v", getProduct)

	if getProduct.Id == 0 {
		return nil, kerrors.NewBizStatusError(40004, "product not exist")
	}

	err = model.AddCart(mysql.DB, s.ctx, &model.Cart{
		UserId:    req.UserId,
		ProductId: req.Item.ProductId,
		Qty:       uint32(req.Item.Quantity),
	})

	return &cart.Empty{}, nil
}
