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

package handler_user

import (
	"context"

	"github.com/cloudwego/biz-demo/book-shop/app/facade/model"
	"github.com/cloudwego/hertz/pkg/app"
)

// UserLogin godoc
// @Summary user login
// @Description user login
// @Tags user module
// @Accept json
// @Produce json
// @Param userParam body model.UserParam true "login param"
// @Success 200 {object} model.LoginResponse
// @Router /user/login [post]
func UserLogin(ctx context.Context, c *app.RequestContext) {
	model.UserAuthMiddleware.LoginHandler(ctx, c)
}
