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

package handlers

import (
	"context"
	"github.com/cloudwego/biz-demo/book-shop/app/facade/model"
	"github.com/cloudwego/hertz/pkg/app"
)

// ShopLogin godoc
// @Summary 商家登录
// @Description 商家登录
// @Tags 商家模块
// @Accept json
// @Produce json
// @Param userParam body handlers.UserParam true "店铺账号信息"
// @Success 200 {object} handlers.LoginResponse
// @Router /shop/login [post]
func ShopLogin(ctx context.Context, c *app.RequestContext) {
	model.UserAuthMiddleware.LoginHandler(ctx, c)
}
