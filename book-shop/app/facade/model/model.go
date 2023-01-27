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

package model

import (
	"net/http"

	"github.com/cloudwego/biz-demo/book-shop/pkg/errno"
	"github.com/cloudwego/hertz/pkg/app"

	"github.com/hertz-contrib/jwt"
)

var (
	UserAuthMiddleware *jwt.HertzJWTMiddleware
	ShopAuthMiddleware *jwt.HertzJWTMiddleware
)

type Response struct {
	Code    int64       `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

// SendResponse pack response
func SendResponse(c *app.RequestContext, err error, data interface{}) {
	Err := errno.ConvertErr(err)
	c.JSON(http.StatusOK, Response{
		Code:    Err.ErrCode,
		Message: Err.ErrMsg,
		Data:    data,
	})
}

type UserParam struct {
	UserName string `json:"username"`
	PassWord string `json:"password"`
}

type LoginResponse struct {
	Code   int64  `json:"code"`
	Expire string `json:"expire"`
	Token  string `json:"token"`
}

type AddProductRequest struct {
	Name        string `json:"name"`
	Pic         string `json:"pic"`
	Description string `json:"description"`
	ISBN        string `json:"isbn"`
	SpuName     string `json:"spu_name"`
	SpuPrice    int64  `json:"spu_price"`
	Price       int64  `json:"price"`
	Stock       int64  `json:"stock"`
}

type EditProductRequest struct {
	ProductId   string  `json:"product_id"`
	Name        *string `json:"name"`
	Pic         *string `json:"pic"`
	Description *string `json:"description"`
	ISBN        *string `json:"isbn"`
	SpuName     *string `json:"spu_name"`
	SpuPrice    *int64  `json:"spu_price"`
	Price       *int64  `json:"price"`
	Stock       *int64  `json:"stock"`
}

type OperateProductReq struct {
	ProductId string `json:"product_id"`
}

type SearchProductReq struct {
	Name        *string `json:"name"`
	Description *string `json:"description"`
	SpuName     *string `json:"spu_name"`
}

type ListProductReq struct {
	Name    *string `json:"name"`
	SpuName *string `json:"spu_name"`
	Status  *int64  `json:"status"`
}

type CreateOrderReq struct {
	Address   string `json:"address"`
	ProductId string `json:"product_id"`
	StockNum  int64  `json:"stock_num"`
}

type CancelOrderReq struct {
	OrderId string `json:"order_id"`
}

type ListOrderReq struct {
	Status *int64 `json:"status"`
}
