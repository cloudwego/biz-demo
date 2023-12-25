package routes

import (
	"context"
	"net/http"

	"github.com/baiyutang/gomall/app/frontend/infra/rpc"
	"github.com/baiyutang/gomall/app/frontend/kitex_gen/user"
	"github.com/baiyutang/gomall/app/frontend/types"
	frontendutils "github.com/baiyutang/gomall/app/frontend/utils"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/app/server"
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
	"github.com/hertz-contrib/sessions"
)

func RegisterAuth(h *server.Hertz) {
	userClient := rpc.UserClient

	h.POST("/auth/register", func(ctx context.Context, c *app.RequestContext) {
		req := &types.RegisterReq{}
		if err := c.BindByContentType(req); err != nil {
			frontendutils.MustHandleError(err)
			return
		}
		resp, err := userClient.Register(ctx, &user.RegisterReq{
			Email:           req.Email,
			Password:        req.Password,
			ConfirmPassword: req.Password,
		})
		frontendutils.MustHandleError(err)

		session := sessions.Default(c)
		session.Set("user_id", resp.Userid)
		hlog.Info("session user_id:", session.Get("user_id"))
		err = session.Save()
		frontendutils.MustHandleError(err)

		c.Redirect(consts.StatusFound, []byte("/"))
	})

	h.POST("/auth/login", func(ctx context.Context, c *app.RequestContext) {
		req := &types.LoginReq{}
		if err := c.BindByContentType(req); err != nil {
			frontendutils.MustHandleError(err)
			return
		}
		resp, err := userClient.Login(ctx, &user.LoginReq{Email: req.Email, Password: req.Password})
		if err != nil {
			c.HTML(http.StatusOK, "error", map[string]interface{}{"message": err})
			return
		}

		session := sessions.Default(c)
		session.Set("user_id", resp.Userid)
		err = session.Save()
		if err != nil {
			c.HTML(http.StatusOK, "error", map[string]interface{}{"message": err})
			return
		}

		c.Redirect(consts.StatusFound, []byte("/"))
	})

	h.GET("/auth/logout", func(ctx context.Context, c *app.RequestContext) {
		session := sessions.Default(c)
		session.Clear()
		session.Save()

		c.Redirect(consts.StatusFound, []byte("/"))
	})
}
