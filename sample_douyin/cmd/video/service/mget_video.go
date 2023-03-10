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
	"github.com/cloudwego/biz-demo/sample_douyin/cmd/video/dal/db"
	"github.com/cloudwego/biz-demo/sample_douyin/kitex_gen/douyinvideo"
	"github.com/cloudwego/biz-demo/sample_douyin/cmd/video/pack"
)

type MGetVideoService struct {
	ctx context.Context
}

// MGetVideoService new MGetVideoService
func NewMGetVideoService(ctx context.Context) *MGetVideoService {
	return &MGetVideoService{ctx: ctx}
}

// MGetVideo.
func (s *MGetVideoService) MGetVideo(req *douyinvideo.MGetVideoRequest) ([]*douyinvideo.Video, error) {
	videos, err := db.MGetVideos(s.ctx, req.VideoIds)
	if err != nil {
		return nil, err
	} 
	return pack.Videos(videos) ,nil
}
