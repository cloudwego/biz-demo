package service

import (
	"context"

	product "github.com/baiyutang/gomall/app/frontend/hertz_gen/frontend/product"
	"github.com/baiyutang/gomall/app/frontend/infra/rpc"
	rpcproduct "github.com/baiyutang/gomall/app/frontend/kitex_gen/product"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/common/utils"
)

type GetProductService struct {
	RequestContext *app.RequestContext
	Context        context.Context
}

func NewGetProductService(Context context.Context, RequestContext *app.RequestContext) *GetProductService {
	return &GetProductService{RequestContext: RequestContext, Context: Context}
}

func (h *GetProductService) Run(req *product.ProductReq) (resp map[string]any, err error) {
	p, _ := rpc.ProductClient.GetProduct(h.Context, &rpcproduct.GetProductRequest{Id: req.GetId()})
	return utils.H{
		"item": p,
	}, nil
}
