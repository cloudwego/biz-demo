// Copyright 2023 CloudWeGo Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//

package rpc

import (
	"context"

	"github.com/cloudwego/biz-demo/sample_douyin/kitex_gen/douyinfavorite"
	"github.com/cloudwego/biz-demo/sample_douyin/kitex_gen/douyinfavorite/favoriteservice"
	"github.com/cloudwego/biz-demo/sample_douyin/pkg/consts"
	"github.com/cloudwego/biz-demo/sample_douyin/pkg/mw"

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