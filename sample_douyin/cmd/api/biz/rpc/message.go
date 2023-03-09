package rpc

import (
	"context"
	"mydouyin/kitex_gen/message"
	"mydouyin/kitex_gen/message/messageservice"
	"mydouyin/pkg/consts"
	"mydouyin/pkg/mw"

	"github.com/cloudwego/kitex/client"
	"github.com/cloudwego/kitex/pkg/rpcinfo"
	"github.com/kitex-contrib/obs-opentelemetry/provider"
	"github.com/kitex-contrib/obs-opentelemetry/tracing"
	etcd "github.com/kitex-contrib/registry-etcd"
)

var messageClient messageservice.Client

func initMessage() {
	r, err := etcd.NewEtcdResolver([]string{consts.ETCDAddress})
	if err != nil {
		panic(err)
	}
	provider.NewOpenTelemetryProvider(
		provider.WithServiceName(consts.ApiServiceName),
		provider.WithExportEndpoint(consts.ExportEndpoint),
		provider.WithInsecure(),
	)
	c, err := messageservice.NewClient(
		consts.MessageServiceName,
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
	messageClient = c
}

func CreateMessage(ctx context.Context, req *message.CreateMessageRequest) (r *message.CreateMessageResponse, err error) {
	return messageClient.CreateMessage(ctx, req)
}

func GetMessageList(ctx context.Context, req *message.GetMessageListRequest) (r *message.GetMessageListResponse, err error) {
	return messageClient.GetMessageList(ctx, req)
}

func GetFirstMessage(ctx context.Context, req *message.GetFirstMessageRequest) (r *message.GetFirstMessageResponse, err error) {
	return messageClient.GetFirstMessage(ctx, req)
}
