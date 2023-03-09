package rpc

import (
	"context"
	"mydouyin/kitex_gen/douyinvideo"
	"mydouyin/kitex_gen/douyinvideo/videoservice"
	"mydouyin/pkg/consts"
	"mydouyin/pkg/mw"

	"github.com/cloudwego/kitex/client"
	"github.com/cloudwego/kitex/pkg/rpcinfo"
	"github.com/kitex-contrib/obs-opentelemetry/provider"
	"github.com/kitex-contrib/obs-opentelemetry/tracing"
	etcd "github.com/kitex-contrib/registry-etcd"
)

var videoClient videoservice.Client

func initVideo() {
	r, err := etcd.NewEtcdResolver([]string{consts.ETCDAddress})
	if err != nil {
		panic(err)
	}
	provider.NewOpenTelemetryProvider(
		provider.WithServiceName(consts.ApiServiceName),
		provider.WithExportEndpoint(consts.ExportEndpoint),
		provider.WithInsecure(),
	)
	c, err := videoservice.NewClient(
		consts.VideoServiceName,
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
	videoClient = c
}

func CreateVideo(ctx context.Context, req *douyinvideo.CreateVideoRequest) (r *douyinvideo.CreateVideoResponse, err error) {
	return videoClient.CreateVideo(ctx, req)
}

func GetFeed(ctx context.Context, req *douyinvideo.GetFeedRequest) (r *douyinvideo.GetFeedResponse, err error) {
	return videoClient.GetFeed(ctx, req)
}

func GetList(ctx context.Context, req *douyinvideo.GetListRequest) (r *douyinvideo.GetListResponse, err error) {
	return videoClient.GetList(ctx, req)
}

func MGetVideo(ctx context.Context, req *douyinvideo.MGetVideoRequest) (r *douyinvideo.MGetVideoResponse, err error) {
	return videoClient.MGetVideoUser(ctx, req)
}

func DeleteVideo(ctx context.Context, req *douyinvideo.DeleteVideoRequest) (r *douyinvideo.DeleteVideoResponse, err error) {
	return videoClient.DeleteVideo(ctx, req)
}

func GetTimeVideos(ctx context.Context, req *douyinvideo.GetTimeVideosRequest) (r *douyinvideo.GetTimeVideosResponse, err error) {
	return videoClient.GetTimeVideos(ctx, req)
}
