package rpc

import (
	"context"
	"mydouyin/kitex_gen/douyinuser"
	"mydouyin/kitex_gen/douyinuser/userservice"
	"mydouyin/pkg/consts"
	"mydouyin/pkg/mw"

	"github.com/cloudwego/kitex/client"
	"github.com/cloudwego/kitex/pkg/rpcinfo"
	"github.com/kitex-contrib/obs-opentelemetry/provider"
	"github.com/kitex-contrib/obs-opentelemetry/tracing"
	etcd "github.com/kitex-contrib/registry-etcd"
)

var userClient userservice.Client

func initUser() {
	r, err := etcd.NewEtcdResolver([]string{consts.ETCDAddress})
	if err != nil {
		panic(err)
	}
	provider.NewOpenTelemetryProvider(
		provider.WithServiceName(consts.ApiServiceName),
		provider.WithExportEndpoint(consts.ExportEndpoint),
		provider.WithInsecure(),
	)
	c, err := userservice.NewClient(
		consts.UserServiceName,
		client.WithResolver(r),
		client.WithMuxConnection(1),
		client.WithMiddleware(mw.CommonMiddleware),
		client.WithInstanceMW(mw.ClientMiddleware),
		client.WithSuite(tracing.NewClientSuite()),
		client.WithClientBasicInfo(&rpcinfo.EndpointBasicInfo{ServiceName: consts.ApiServiceName}),
	)
	if err != nil {
		panic(err)
	}
	userClient = c
}

func CreateUser(ctx context.Context, req *douyinuser.CreateUserRequest) (r *douyinuser.CreateUserResponse, err error) {
	return userClient.CreateUser(ctx, req)
}

func CheckUser(ctx context.Context, req *douyinuser.CheckUserRequest) (r *douyinuser.CheckUserResponse, err error) {
	return userClient.CheckUser(ctx, req)
}

func MGetUser(ctx context.Context, req *douyinuser.MGetUserRequest) (r *douyinuser.MGetUserResponse, err error) {
	return userClient.MGetUser(ctx, req)
}
