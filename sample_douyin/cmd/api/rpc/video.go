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

	"github.com/cloudwego/biz-demo/sample_douyin/kitex_gen/douyinvideo"
	"github.com/cloudwego/biz-demo/sample_douyin/kitex_gen/douyinvideo/videoservice"
	"github.com/cloudwego/biz-demo/sample_douyin/pkg/consts"
	"github.com/cloudwego/biz-demo/sample_douyin/pkg/mw"

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
