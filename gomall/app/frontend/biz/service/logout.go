package service

import (
	"context"

	common "github.com/cloudwego/biz-demo/gomall/app/frontend/hertz_gen/frontend/common"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/hertz-contrib/sessions"
)

type LogoutService struct {
	RequestContext *app.RequestContext
	Context        context.Context
}

func NewLogoutService(Context context.Context, RequestContext *app.RequestContext) *LogoutService {
	return &LogoutService{RequestContext: RequestContext, Context: Context}
}

func (h *LogoutService) Run(req *common.Empty) (resp *common.Empty, err error) {
	session := sessions.Default(h.RequestContext)
	session.Clear()
	session.Save()
	return
}
