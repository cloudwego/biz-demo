package utils

import (
	"context"

	"github.com/cloudwego/biz-demo/gomall/app/frontend/infra/rpc"
	"github.com/cloudwego/biz-demo/gomall/app/frontend/kitex_gen/cart"
	frontendutils "github.com/cloudwego/biz-demo/gomall/app/frontend/utils"

	"github.com/cloudwego/hertz/pkg/app"
)

// SendErrResponse  pack error response
func SendErrResponse(ctx context.Context, c *app.RequestContext, code int, err error) {
	// todo edit custom code
	c.String(code, err.Error())
}

// SendSuccessResponse  pack success response
func SendSuccessResponse(ctx context.Context, c *app.RequestContext, code int, data interface{}) {
	// todo edit custom code
	c.JSON(code, data)
}

func WarpResponse(ctx context.Context, c *app.RequestContext, content map[string]any) map[string]any {
	var cartNum int
	userId := frontendutils.GetUserIdFromCtx(ctx)
	cartResp, _ := rpc.CartClient.GetCart(ctx, &cart.GetCartRequest{UserId: userId})
	if cartResp != nil {
		cartNum = len(cartResp.Items)
	}
	content["user_id"] = ctx.Value(frontendutils.UserIdKey)
	content["cart_num"] = cartNum
	return content
}
