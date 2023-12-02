package main

import (
	"context"
	product "github.com/baiyutang/gomall/app/product/kitex_gen/product"
	"github.com/baiyutang/gomall/app/product/biz/service"
)

// ProductCatalogServiceImpl implements the last service interface defined in the IDL.
type ProductCatalogServiceImpl struct{}

// ListProducts implements the ProductCatalogServiceImpl interface.
func (s *ProductCatalogServiceImpl) ListProducts(ctx context.Context, req *product.Empty) (resp *product.ListProductsResponse, err error) {
	resp, err = service.NewListProductsService(ctx).Run(req)

	return resp, err
}

// GetProduct implements the ProductCatalogServiceImpl interface.
func (s *ProductCatalogServiceImpl) GetProduct(ctx context.Context, req *product.GetProductRequest) (resp *product.Product, err error) {
	resp, err = service.NewGetProductService(ctx).Run(req)

	return resp, err
}

// SearchProducts implements the ProductCatalogServiceImpl interface.
func (s *ProductCatalogServiceImpl) SearchProducts(ctx context.Context, req *product.SearchProductsRequest) (resp *product.SearchProductsResponse, err error) {
	resp, err = service.NewSearchProductsService(ctx).Run(req)

	return resp, err
}
