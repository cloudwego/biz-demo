// Copyright 2023 CloudWeGo Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//

package handler

import (
	"context"

	"github.com/cloudwego/biz-demo/book-shop/app/item/common/constant"
	"github.com/cloudwego/biz-demo/book-shop/app/item/domain/service"
	"github.com/cloudwego/biz-demo/book-shop/kitex_gen/cwg/bookshop/item"
	"github.com/cloudwego/biz-demo/book-shop/pkg/errno"
)

type OnlineHandler struct {
	ctx   context.Context
	param *item.OnlineReq
}

func NewOnlineHandler(ctx context.Context, req *item.OnlineReq) *OnlineHandler {
	return &OnlineHandler{
		ctx:   ctx,
		param: req,
	}
}

func (h *OnlineHandler) Online() (*item.OnlineResp, error) {
	resp := &item.OnlineResp{
		BaseResp: errno.BuildBaseResp(errno.Success),
	}

	stateService := service.GetProductStateService()
	queryService := service.GetProductQueryServiceInstance()
	updateService := service.GetProductUpdateServiceInstance()

	// 0. get origin info
	originEntity, err := queryService.GetProduct(h.ctx, h.param.ProductId)
	if err != nil {
		resp.BaseResp = errno.BuildBaseResp(err)
		return resp, nil
	}

	// 1. validate
	validateFunc, err := stateService.GetCanTransferFunc(constant.StateOperationTypeOnline)
	if err != nil {
		resp.BaseResp = errno.BuildBaseResp(err)
		return resp, nil
	}
	err = validateFunc(&service.ProductStateInfo{Status: originEntity.Status})
	if err != nil {
		resp.BaseResp = errno.BuildBaseResp(err)
		return resp, nil
	}

	// 2. construct do
	targetEntity, err := stateService.ConstructTargetInfo(originEntity, constant.StateOperationTypeOnline)
	if err != nil {
		resp.BaseResp = errno.BuildBaseResp(err)
		return resp, nil
	}

	// 3. process
	err = updateService.EditProduct(h.ctx, originEntity, targetEntity)
	if err != nil {
		resp.BaseResp = errno.BuildBaseResp(err)
		return resp, nil
	}

	return resp, nil
}
