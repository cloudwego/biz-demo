package service

import (
	"context"
	"github.com/cloudwego/hertz/pkg/common/utils"

	common "github.com/baiyutang/gomall/app/frontend/hertz_gen/frontend/common"
	"github.com/cloudwego/hertz/pkg/app"
)

type CheckoutResultService struct {
	RequestContext *app.RequestContext
	Context        context.Context
}

func NewCheckoutResultService(Context context.Context, RequestContext *app.RequestContext) *CheckoutResultService {
	return &CheckoutResultService{RequestContext: RequestContext, Context: Context}
}

func (h *CheckoutResultService) Run(req *common.Empty) (resp map[string]any, err error) {
	return utils.H{}, nil
}
