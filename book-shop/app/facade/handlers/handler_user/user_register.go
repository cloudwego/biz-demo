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
	"github.com/cloudwego/biz-demo/book-shop/app/facade/infras/client"
	"github.com/cloudwego/biz-demo/book-shop/app/facade/model"
	"github.com/cloudwego/biz-demo/book-shop/kitex_gen/cwg/bookshop/user"
	"github.com/cloudwego/biz-demo/book-shop/pkg/errno"
	"github.com/cloudwego/hertz/pkg/app"
)

// UserRegister godoc
// @Summary 用户注册
// @Description 用户注册
// @Tags 用户模块
// @Accept json
// @Produce json
// @Param userParam body handlers.UserParam true "注册信息"
// @Success 200 {object} handlers.Response
// @Router /user/register [post]
func UserRegister(ctx context.Context, c *app.RequestContext) {
	var registerParam model.UserParam
	if err := c.BindAndValidate(&registerParam); err != nil {
		model.SendResponse(c, errno.ConvertErr(err), nil)
		return
	}

	if len(registerParam.UserName) == 0 || len(registerParam.PassWord) == 0 {
		model.SendResponse(c, errno.ParamErr, nil)
		return
	}

	err := client.CreateUser(ctx, &user.CreateUserReq{
		UserName: registerParam.UserName,
		Password: registerParam.PassWord,
	})
	if err != nil {
		model.SendResponse(c, errno.ConvertErr(err), nil)
		return
	}
	model.SendResponse(c, errno.Success, nil)
}
