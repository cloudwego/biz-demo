package main

import (
	"context"
	"github.com/cloudwego/biz-demo/mall/cmd/cmp_ecom_shop/app"
	"github.com/cloudwego/biz-demo/mall/cmd/cmp_ecom_shop/common"
	"github.com/cloudwego/biz-demo/mall/cmd/cmp_ecom_shop/kitex_gen/cmp/ecom/shop"
	"github.com/cloudwego/biz-demo/mall/pkg/errno"
)

// ShopServiceImpl implements the last service interface defined in the IDL.
type ShopServiceImpl struct{}

// SettleShop implements the ShopServiceImpl interface.
func (s *ShopServiceImpl) SettleShop(ctx context.Context, req *shop.SettleShopReq) (resp *shop.SettleShopResp, err error) {
	resp = shop.NewSettleShopResp()
	if req.ShopName == "" {
		resp.BaseResp = common.BuildBaseResp(errno.ParamErr)
		return resp, nil
	}

	shopId, err := app.NewShopService(ctx).SettleShop(req)
	if err != nil {
		resp.BaseResp = common.BuildBaseResp(err)
		return resp, nil
	}
	resp.ShopId = shopId
	resp.BaseResp = common.BuildBaseResp(errno.Success)
	return resp, nil
}

// GetShopIdByUserId implements the ShopServiceImpl interface.
func (s *ShopServiceImpl) GetShopIdByUserId(ctx context.Context, req *shop.GetShopIdByUserIdReq) (resp *shop.GetShopIdByUserIdResp, err error) {
	resp = shop.NewGetShopIdByUserIdResp()
	shopId, err := app.NewShopService(ctx).GetShopIdByUserId(req)
	if err != nil {
		resp.BaseResp = common.BuildBaseResp(err)
		return resp, nil
	}
	resp.ShopId = shopId
	resp.BaseResp = common.BuildBaseResp(errno.Success)
	return resp, nil
}
