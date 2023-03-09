package rpc

import (
	"context"
	"mydouyin/kitex_gen/douyinfavorite"
	"mydouyin/kitex_gen/douyinfavorite/favoriteservice"
	"mydouyin/pkg/consts"
	"mydouyin/pkg/mw"

	"github.com/cloudwego/kitex/client"
	"github.com/cloudwego/kitex/pkg/rpcinfo"
	"github.com/kitex-contrib/obs-opentelemetry/provider"
	"github.com/kitex-contrib/obs-opentelemetry/tracing"
	etcd "github.com/kitex-contrib/registry-etcd"
)

var favoriteClient favoriteservice.Client

func initFavorite() {
	r, err := etcd.NewEtcdResolver([]string{consts.ETCDAddress})
	if err != nil {
		panic(err)
	}
	provider.NewOpenTelemetryProvider(
		provider.WithServiceName(consts.ApiServiceName),
		provider.WithExportEndpoint(consts.ExportEndpoint),
		provider.WithInsecure(),
	)
	c, err := favoriteservice.NewClient(
		consts.FavoriteServiceName,
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
	favoriteClient = c
}

func FavoriteAction(ctx context.Context, req *douyinfavorite.FavoriteActionRequest) (r *douyinfavorite.FavoriteActionResponse, err error) {
	return favoriteClient.FavoriteAction(ctx, req)
}

func GetFavouriteList(ctx context.Context, req *douyinfavorite.GetListRequest) (r *douyinfavorite.GetListResponse, err error) {
	return favoriteClient.GetList(ctx, req)
}

func GetIsFavorite(ctx context.Context, req *douyinfavorite.GetIsFavoriteRequest) (r *douyinfavorite.GetIsFavoriteResponse, err error) {
	return favoriteClient.GetIsFavorite(ctx, req)
}
