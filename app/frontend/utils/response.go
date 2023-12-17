package utils

import (
	"context"
	"github.com/cloudwego/hertz/pkg/app"
)

func WarpResponse(ctx context.Context, c *app.RequestContext, content map[string]any) map[string]any {
	content["user_id"] = ctx.Value(UserIdKey)
	return content
}
