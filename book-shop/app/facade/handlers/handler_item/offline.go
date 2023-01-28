// Copyright 2022 CloudWeGo Authors
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

package handler_item

import (
	"context"
	"strconv"

	"github.com/cloudwego/biz-demo/book-shop/app/facade/infras/client"
	"github.com/cloudwego/biz-demo/book-shop/app/facade/model"
	"github.com/cloudwego/biz-demo/book-shop/pkg/errno"
	"github.com/cloudwego/hertz/pkg/app"
)

// OfflineProduct godoc
// @Summary offline product
// @Description offline product
// @Tags product module
// @Accept json
// @Produce json
// @Param offlineProductRequest body model.OperateProductReq true "request param of operating product"
// @Security TokenAuth
// @Success 200 {object} model.Response
// @Router /item2b/offline [post]
func OfflineProduct(ctx context.Context, c *app.RequestContext) {
	var offlineReq model.OperateProductReq
	if err := c.BindAndValidate(&offlineReq); err != nil {
		model.SendResponse(c, errno.ConvertErr(err), nil)
		return
	}

	pid, err := strconv.ParseInt(offlineReq.ProductId, 10, 64)
	if err != nil {
		model.SendResponse(c, errno.ConvertErr(err), nil)
		return
	}

	err = client.OperateProduct(ctx, pid, "offline")
	if err != nil {
		model.SendResponse(c, errno.ConvertErr(err), nil)
		return
	}
	model.SendResponse(c, errno.Success, nil)
}
