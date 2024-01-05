package weall

import (
	"context"
	rpc_gen "github.com/cloudwego/biz-demo/gomall/rpc_gen/kitex_gen/rpc_gen"
	"github.com/cloudwego/kitex/client/callopt"
	"github.com/cloudwego/kitex/pkg/klog"
)

func Register(ctx context.Context, req *rpc_gen.RegisterReq, callOptions ...callopt.Option) (resp *rpc_gen.RegisterResp, err error) {
	resp, err = defaultClient.Register(ctx, req, callOptions...)
	if err != nil {
		klog.CtxErrorf(ctx, "Register call failed,err =%+v", err)
		return nil, err
	}
	return resp, nil
}

func Login(ctx context.Context, req *rpc_gen.LoginReq, callOptions ...callopt.Option) (resp *rpc_gen.LoginRes, err error) {
	resp, err = defaultClient.Login(ctx, req, callOptions...)
	if err != nil {
		klog.CtxErrorf(ctx, "Login call failed,err =%+v", err)
		return nil, err
	}
	return resp, nil
}
