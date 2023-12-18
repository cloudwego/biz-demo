package routes

import (
	"context"
	frontendutils "github.com/baiyutang/gomall/app/frontend/utils"
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"github.com/hertz-contrib/sessions"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/app/server"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
	"github.com/cloudwego/kitex/client"

	"github.com/baiyutang/gomall/app/frontend/kitex_gen/user"
	"github.com/baiyutang/gomall/app/frontend/kitex_gen/user/userservice"
)

func RegisterAuth(h *server.Hertz) {
	userClient, _ := userservice.NewClient("user", client.WithHostPorts("localhost:8882"))
	h.POST("/auth/register", func(ctx context.Context, c *app.RequestContext) {
		_, err := userClient.Register(context.Background(), &user.RegisterReq{Email: "abc@abc.com", Password: "hello@password"})
		frontendutils.MustHandleError(err)

		session := sessions.Default(c)
		session.Set("user_id", 1)
		hlog.Info(session.Get("user_id"), "session user_id")
		err = session.Save()
		frontendutils.MustHandleError(err)

		c.Redirect(consts.StatusFound, []byte("/"))
	})
	h.POST("/auth/login", func(ctx context.Context, c *app.RequestContext) {
		//p, _ := userClient.Login(context.Background(), &user.LoginReq{Email: "abc@abc.com", Password: "hello@password"})
		session := sessions.Default(c)
		session.Set("user_id", 1)
		err := session.Save()
		frontendutils.MustHandleError(err)

		c.Redirect(consts.StatusFound, []byte("/"))
	})
	h.GET("/auth/logout", func(ctx context.Context, c *app.RequestContext) {
		session := sessions.Default(c)
		session.Clear()
		session.Save()

		c.Redirect(consts.StatusFound, []byte("/"))
	})
}
