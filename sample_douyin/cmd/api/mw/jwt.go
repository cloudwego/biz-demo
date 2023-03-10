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

package mw

import (
	"context"
	"net/http"
	"time"

	"github.com/cloudwego/biz-demo/sample_douyin/cmd/api/rpc"
	douyinapi "github.com/cloudwego/biz-demo/sample_douyin/hertz_gen/douyinapi"
	"github.com/cloudwego/biz-demo/sample_douyin/kitex_gen/douyinuser"
	"github.com/cloudwego/biz-demo/sample_douyin/pkg/consts"
	"github.com/cloudwego/biz-demo/sample_douyin/pkg/errno"
	protocol_consts "github.com/cloudwego/hertz/pkg/protocol/consts"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/common/utils"
	"github.com/hertz-contrib/jwt"
)

var JwtMiddleware *HertzJWTMiddleware

func InitJWT() {
	JwtMiddleware, _ = New(&HertzJWTMiddleware{
		//用于设置签名密钥*
		Key: []byte(consts.SecretKey),
		//用于设置token的获取源，可选header、query、cookie、param、form 默认header: Authorization
		TokenLookup: "header: Authorization, query: token, form: token, cookie: jwt",
		//用于设置从header中获取token时的前缀，默认为Bearer
		TokenHeadName: "Bearer",
		//设置获取当前时间的函数
		TimeFunc: time.Now,
		//token的过期时间
		Timeout: 24 * time.Hour,
		//最大token刷新时间，允许客户端在tokenTime+MaxRefresh内刷新token的有效时间，追加一个timeout时长
		MaxRefresh: time.Hour,
		//用于设置检索身份的键 默认indentity
		IdentityKey: consts.IdentityKey,
		//设置获取身份信息的函数,与PayloadFunc一致
		IdentityHandler: func(ctx context.Context, c *app.RequestContext) interface{} {
			claims := ExtractClaims(ctx, c)
			return &douyinapi.User{
				ID: int64(claims[consts.IdentityKey].(float64)),
			}
		},
		//登陆成功后为向token中添加自定义负载信息的函数,额外存储了用户id，如不设置则只存储token的过期时间和创建时间
		PayloadFunc: func(data interface{}) MapClaims {
			if v, ok := data.(int64); ok {
				return MapClaims{
					consts.IdentityKey: v,
				}
			}
			return MapClaims{}
		},
		//设置登陆时认证用户信息的函数**
		Authenticator: func(ctx context.Context, c *app.RequestContext) (interface{}, error) {
			var err error
			var req douyinapi.CheckUserRequest
			if err = c.BindAndValidate(&req); err != nil {
				return "", jwt.ErrMissingLoginValues
			}
			if len(req.Username) == 0 || len(req.Password) == 0 {
				return "", jwt.ErrMissingLoginValues
			}
			resp, err := rpc.CheckUser(context.Background(), &douyinuser.CheckUserRequest{
				Username: req.Username,
				Password: req.Password,
			})
			if err != nil {
				return 0, err
			}
			if resp.BaseResp.StatusCode != 0 {
				return 0, errno.NewErrNo(resp.BaseResp.StatusCode, resp.BaseResp.StatusMessage)
			}
			return resp.UserId, nil
		},
		//设置登陆的响应函数
		LoginResponse: func(ctx context.Context, c *app.RequestContext, code int, token string, expire time.Time, identity interface{}) {
			res := &douyinapi.CheckUserResponse{
				StatusCode: errno.Success.ErrCode,
				StatusMsg:  "",
				UserID:     identity.(int64),
				Token:      token,
			}
			c.JSON(protocol_consts.StatusOK, res)
		},
		//验证流程失败的响应函数
		Unauthorized: func(ctx context.Context, c *app.RequestContext, code int, message string) {
			c.JSON(http.StatusOK, utils.H{
				"status_code": errno.AuthorizationFailedErr.ErrCode,
				"status_msg":  message,
				"token":       nil,
				"user_id":     nil,
			})
		},
		//设置jwt校验流程发生错误时相应所包含的错误信息
		HTTPStatusMessageFunc: func(e error, ctx context.Context, c *app.RequestContext) string {
			switch t := e.(type) {
			case errno.ErrNo:
				return t.ErrMsg
			default:
				return t.Error()
			}
		},
	})
}
