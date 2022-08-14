package main

import (
	"context"
	"github.com/cloudwego/biz-demo/mall/cmd/cmp_ecom_product/kitex_gen/cmp/ecom/product"
)

// ProductServiceImpl implements the last service interface defined in the IDL.
type ProductServiceImpl struct{}

// AddBrand implements the ProductServiceImpl interface.
func (s *ProductServiceImpl) AddBrand(ctx context.Context, req *product.AddBrandReq) (resp *product.AddBrandResp, err error) {
	// TODO: Your code here...
	return
}

// UpdateBrand implements the ProductServiceImpl interface.
func (s *ProductServiceImpl) UpdateBrand(ctx context.Context, req *product.UpdateBrandReq) (resp *product.UpdateBrandResp, err error) {
	// TODO: Your code here...
	return
}

// DeleteBrand implements the ProductServiceImpl interface.
func (s *ProductServiceImpl) DeleteBrand(ctx context.Context, req *product.DeleteBrandReq) (resp *product.DeleteBrandResp, err error) {
	// TODO: Your code here...
	return
}

// GetBrandsByShopId implements the ProductServiceImpl interface.
func (s *ProductServiceImpl) GetBrandsByShopId(ctx context.Context, req *product.GetBrandsByShopIdReq) (resp *product.GetBrandsByShopIdResp, err error) {
	// TODO: Your code here...
	return
}

// PassProduct implements the ProductServiceImpl interface.
func (s *ProductServiceImpl) PassProduct(ctx context.Context, req *product.PassProductReq) (resp *product.PassProductResp, err error) {
	// TODO: Your code here...
	return
}

// RejectProduct implements the ProductServiceImpl interface.
func (s *ProductServiceImpl) RejectProduct(ctx context.Context, req *product.RejectProductReq) (resp *product.RejectProductResp, err error) {
	// TODO: Your code here...
	return
}
