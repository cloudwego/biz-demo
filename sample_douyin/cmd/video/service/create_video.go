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
)

type CreateVideoService struct {
	ctx context.Context
}

// NewCreateVideoService new CreateVideoService
func NewCreateVideoService(ctx context.Context) *CreateVideoService {
	return &CreateVideoService{ctx: ctx}
}

// CreateVideo create video info.
func (s *CreateVideoService) CreateVideo(req *douyinvideo.CreateVideoRequest) ([]int64, error) {
	return db.CreateVideo(s.ctx, []*db.Video{{
		Author:   req.Author,
		PlayUrl:  req.PlayUrl,
		CoverUrl: req.CoverUrl,
		Title:    req.Title,
	}})
}
