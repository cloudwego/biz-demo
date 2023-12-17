package routes

import (
	"context"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/app/server"
	"github.com/cloudwego/hertz/pkg/common/utils"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
	"github.com/cloudwego/kitex/client"

	"github.com/baiyutang/gomall/app/frontend/kitex_gen/user"
	"github.com/baiyutang/gomall/app/frontend/kitex_gen/user/userservice"
)

func RegisterAuth(h *server.Hertz) {
	userClient, _ := userservice.NewClient("user", client.WithHostPorts("localhost:8882"))
	h.GET("/auth/register", func(ctx context.Context, c *app.RequestContext) {
		p, _ := userClient.Register(context.Background(), &user.RegisterReq{Email: "abc@abc.com", Password: "hello@password"})
		c.HTML(consts.StatusOK, "home", utils.H{
			"cart_num": 10,
			"user_id":  p.Userid,
		})
	})
	h.GET("/auth/login", func(ctx context.Context, c *app.RequestContext) {
		p, _ := userClient.Login(context.Background(), &user.LoginReq{Email: "abc@abc.com", Password: "hello@password"})
		c.HTML(consts.StatusOK, "home", utils.H{
			"cart_num": 10,
			"user_id":  p.Userid,
		})
	})
}
