package service

import (
	"context"
	product "github.com/baiyutang/gomall/app/product/kitex_gen/product"
)

type ListProductsService struct {
	ctx context.Context
} // NewListProductsService new ListProductsService
func NewListProductsService(ctx context.Context) *ListProductsService {
	return &ListProductsService{ctx: ctx}
}

// Run create note info
func (s *ListProductsService) Run(req *product.Empty) (resp *product.ListProductsResponse, err error) {
	// Finish your business logic.

	return
}
