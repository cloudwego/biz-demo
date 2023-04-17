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

package service

import (
	"context"
	"fmt"
	"log"
	"strconv"
	"time"

	"github.com/cloudwego/biz-demo/sample_douyin/cmd/video/dal/db"
	"github.com/cloudwego/biz-demo/sample_douyin/cmd/video/pack"
	"github.com/cloudwego/biz-demo/sample_douyin/kitex_gen/douyinvideo"

	"github.com/cloudwego/kitex/pkg/klog"
)

type GetFeedService struct {
	ctx context.Context
}

// NewGetFeedService new GetFeedService
func NewGetFeedService(ctx context.Context) *GetFeedService {
	return &GetFeedService{ctx: ctx}
}

// GetFeedService.
func (s *GetFeedService) GetFeed(req *douyinvideo.GetFeedRequest) (int64, []*douyinvideo.Video, error) {
	fmt.Printf("video")
	log.Println(req.LatestTime)
	latestTime, err := strconv.Atoi(req.LatestTime)
	if err != nil {
		return time.Now().Unix(), nil, err
	}
	latestTimeStr := time.Unix(int64(latestTime), 0).Format("2006-01-02 15:04:05")
	videos, err := db.GetFeed(s.ctx, latestTimeStr)
	klog.Infof("请求时间：%v", req.LatestTime)
	if err != nil {
		return time.Now().Unix(), nil, err
	}
	if len(videos) < 1 {
		return time.Now().Unix(), nil, nil
	}
	var index int = len(videos) - 1
	return videos[index].CreatedAt.Unix(), pack.Videos(videos), nil
}
