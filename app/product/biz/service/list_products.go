package service

import (
	"context"
	"github.com/baiyutang/gomall/app/product/biz/dal/mysql"
	"github.com/baiyutang/gomall/app/product/biz/model"
	product "github.com/baiyutang/gomall/app/product/kitex_gen/product"
)

type ListProductsService struct {
	ctx context.Context
} // NewListProductsService new ListProductsService
func NewListProductsService(ctx context.Context) *ListProductsService {
	return &ListProductsService{ctx: ctx}
}

// Run create note info
func (s *ListProductsService) Run(req *product.ListProductsReq) (resp *product.ListProductsResponse, err error) {
	// Finish your business logic.
	c, err := model.GetProductsByCategoryName(mysql.DB, req.CategoryName)
	if err != nil {
		return nil, err
	}
	resp = &product.ListProductsResponse{}
	for _, v1 := range c {
		for _, v := range v1.Products {
			resp.Products = append(resp.Products, &product.Product{Id: uint32(v.ID), Name: v.Name, Description: v.Description, Picture: v.Picture, Price: v.Price})
		}
	}

	return resp, nil
}
