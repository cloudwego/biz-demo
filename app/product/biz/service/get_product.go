package service

import (
	"context"
	"github.com/baiyutang/gomall/app/product/biz/dal/mysql"
	"github.com/baiyutang/gomall/app/product/biz/model"
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
	p, err := model.GetProductById(mysql.DB, int(req.Id))
	return &product.Product{
		Id:          uint32(p.ID),
		Picture:     p.Picture,
		Price:       p.Price,
		Description: p.Description,
		Name:        p.Name,
	}, err
}
