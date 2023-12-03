package service

import (
	"context"

	product "github.com/baiyutang/gomall/app/product/kitex_gen/product"
)

type GetProductService struct {
	ctx context.Context
} // NewGetProductService new GetProductService
func NewGetProductService(ctx context.Context) *GetProductService {
	return &GetProductService{ctx: ctx}
}

// Run create note info
func (s *GetProductService) Run(req *product.GetProductRequest) (resp *product.Product, err error) {
	// Finish your business logic.

	return &product.Product{Id: "first", Name: "first goods"}, nil
}
