package service

import (
	"context"
	product "github.com/baiyutang/gomall/app/product/kitex_gen/product"
)

type SearchProductsService struct {
	ctx context.Context
} // NewSearchProductsService new SearchProductsService
func NewSearchProductsService(ctx context.Context) *SearchProductsService {
	return &SearchProductsService{ctx: ctx}
}

// Run create note info
func (s *SearchProductsService) Run(req *product.SearchProductsRequest) (resp *product.SearchProductsResponse, err error) {
	// Finish your business logic.

	return
}
