package middleware

import (
	"context"

	"github.com/cloudwego/kitex/pkg/endpoint"
	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/cloudwego/kitex/pkg/rpcinfo"
)

func ServerMiddleware(next endpoint.Endpoint) endpoint.Endpoint {
	return func(ctx context.Context, req, resp interface{}) (err error) {
		ri := rpcinfo.GetRPCInfo(ctx)
		// get client serviceName
		klog.Infof("client serviceName: %v\n", ri.From().ServiceName())
		if err := next(ctx, req, resp); err != nil {
			return err
		}
		return nil
	}
}
