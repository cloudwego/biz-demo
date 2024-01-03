package service

import (
	"context"

	"github.com/cloudwego/biz-demo/gomall/app/product/biz/dal/mysql"
	"github.com/cloudwego/biz-demo/gomall/app/product/biz/dal/redis"
	"github.com/cloudwego/biz-demo/gomall/app/product/biz/model"
	product "github.com/cloudwego/biz-demo/gomall/app/product/kitex_gen/product"
	"github.com/cloudwego/kitex/pkg/kerrors"
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
	if req.Id == 0 {
		return nil, kerrors.NewBizStatusError(40000, "product id is required")
	}

	p, err := model.NewCachedProductQuery(model.NewProductQuery(s.ctx, mysql.DB), redis.RedisClient).GetById(int(req.Id))
	return &product.Product{
		Id:          uint32(p.ID),
		Picture:     p.Picture,
		Price:       p.Price,
		Description: p.Description,
		Name:        p.Name,
	}, err
}
