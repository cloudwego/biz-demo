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

	"github.com/cloudwego/biz-demo/book-shop/app/item/common/converter"
	"github.com/cloudwego/biz-demo/book-shop/app/item/domain/service"
	"github.com/cloudwego/biz-demo/book-shop/kitex_gen/cwg/bookshop/item"
	"github.com/cloudwego/biz-demo/book-shop/pkg/errno"
)

type MGet2CHandler struct {
	ctx   context.Context
	param *item.MGet2CReq
}

func NewMGet2CHandler(ctx context.Context, req *item.MGet2CReq) *MGet2CHandler {
	return &MGet2CHandler{
		ctx:   ctx,
		param: req,
	}
}

func (h *MGet2CHandler) MGet() (*item.MGet2CResp, error) {
	resp := &item.MGet2CResp{
		BaseResp: errno.BuildBaseResp(errno.Success),
	}

	queryService := service.GetProductQueryServiceInstance()
	entities, err := queryService.MGet2C(h.ctx, h.param.ProductIds)
	if err != nil {
		resp.BaseResp = errno.BuildBaseResp(err)
		return resp, nil
	}
	dtoMap := make(map[int64]*item.Product)
	for _, e := range entities {
		dtoMap[e.ProductId] = converter.ConvertEntity2DTO(e)
	}

	resp.ProductMap = dtoMap

	return resp, nil
}
