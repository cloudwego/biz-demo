package middleware

import (
	"context"
	"github.com/baiyutang/gomall/app/frontend/utils"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"github.com/hertz-contrib/sessions"
)

var authPath = []string{
	"/cart",
}

func Auth() app.HandlerFunc {
	return func(ctx context.Context, c *app.RequestContext) {
		session := sessions.Default(c)
		hlog.Info("full path", c.FullPath())
		hlog.Info(utils.InArray(c.FullPath(), authPath))
		//if !utils.InArray(c.FullPath(), authPath) {
		//	c.Next(ctx)
		//	return
		//}
		userId := session.Get("user_id")
		hlog.Info("user_id", userId)
		if userId == nil {
			//c.Redirect(302, []byte("/sign-in"))
			//c.Abort()
			c.Next(ctx)
			return
		}
		ctx = context.WithValue(ctx, utils.UserIdKey, userId)
		c.Next(ctx)
	}
}
