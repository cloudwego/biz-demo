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

type ListHandler struct {
	ctx   context.Context
	param *item.ListReq
}

func NewListHandler(ctx context.Context, req *item.ListReq) *ListHandler {
	return &ListHandler{
		ctx:   ctx,
		param: req,
	}
}

func (h *ListHandler) List() (*item.ListResp, error) {
	resp := &item.ListResp{
		BaseResp: errno.BuildBaseResp(errno.Success),
	}

	if h.param.Name == nil && h.param.SpuName == nil && h.param.Status == nil {
		resp.BaseResp = errno.BuildBaseResp(errno.ParamErr)
		return resp, nil
	}

	queryService := service.GetProductQueryServiceInstance()
	entities, err := queryService.ListProducts(h.ctx, h.param.Name, h.param.SpuName, (*int64)(h.param.Status))
	if err != nil {
		resp.BaseResp = errno.BuildBaseResp(err)
		return resp, nil
	}

	dtos := make([]*item.Product, 0)
	for _, e := range entities {
		dtos = append(dtos, converter.ConvertEntity2DTO(e))
	}

	resp.Products = dtos

	return resp, nil
}
