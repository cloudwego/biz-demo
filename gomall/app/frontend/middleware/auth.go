package middleware

import (
	"context"
	"fmt"

	"github.com/cloudwego/biz-demo/gomall/app/frontend/utils"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/hertz-contrib/sessions"
)

func GlobalAuth() app.HandlerFunc {
	return func(ctx context.Context, c *app.RequestContext) {
		session := sessions.Default(c)
		userId := session.Get("user_id")
		if userId == nil {
			c.Next(ctx)
			return
		}
		ctx = context.WithValue(ctx, utils.UserIdKey, userId)
		c.Next(ctx)
	}
}

func Auth() app.HandlerFunc {
	return func(ctx context.Context, c *app.RequestContext) {
		session := sessions.Default(c)
		userId := session.Get("user_id")
		if userId == nil {
			byteRef := c.GetHeader("Referer")
			ref := string(byteRef)
			next := "/sign-in"
			if ref != "" {
				if utils.ValidateNext(ref) {
					next = fmt.Sprintf("%s?next=%s", next, ref)
				}
			}
			c.Redirect(302, []byte(next))
			c.Abort()
			c.Next(ctx)
			return
		}
		ctx = context.WithValue(ctx, utils.UserIdKey, userId)
		c.Next(ctx)
	}
}
