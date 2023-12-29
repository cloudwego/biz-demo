package service

import (
	"context"
	"github.com/baiyutang/gomall/app/frontend/hertz_gen/frontend/cart"
	common "github.com/baiyutang/gomall/app/frontend/hertz_gen/frontend/common"
	"github.com/baiyutang/gomall/app/frontend/infra/rpc"
	rpccart "github.com/baiyutang/gomall/app/frontend/kitex_gen/cart"
	frontendutils "github.com/baiyutang/gomall/app/frontend/utils"
	"github.com/cloudwego/hertz/pkg/app"
)

type AddCartItemService struct {
	RequestContext *app.RequestContext
	Context        context.Context
}

func NewAddCartItemService(Context context.Context, RequestContext *app.RequestContext) *AddCartItemService {
	return &AddCartItemService{RequestContext: RequestContext, Context: Context}
}

func (h *AddCartItemService) Run(req *cart.AddCartReq) (resp *common.Empty, err error) {
	_, err = rpc.CartClient.AddItem(h.Context, &rpccart.AddItemRequest{
		UserId: uint32(h.Context.Value(frontendutils.UserIdKey).(float64)),
		Item: &rpccart.CartItem{
			ProductId: req.ProductId,
			Quantity:  req.ProductNum,
		}})
	return
}
