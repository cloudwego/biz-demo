package mw

import (
	"context"
	"mydouyin/cmd/api/biz/apimodel"
	"mydouyin/cmd/api/biz/rpc"
	"mydouyin/kitex_gen/douyinuser"
	"mydouyin/pkg/consts"
	"mydouyin/pkg/errno"
	"net/http"
	"time"

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
			return &apimodel.User{
				UserID: int64(claims[consts.IdentityKey].(float64)),
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
			var req apimodel.CheckUserRequest
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
			res := &apimodel.CheckUserResponse{
				StatusCode: errno.Success.ErrCode,
				StatusMsg:  "",
				UserId:     identity.(int64),
				Token:      token,
			}
			res.Send(c)
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
