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

	"github.com/cloudwego/biz-demo/sample_douyin/cmd/message/dal/db"
	"github.com/cloudwego/biz-demo/sample_douyin/cmd/message/pack"
	"github.com/cloudwego/biz-demo/sample_douyin/kitex_gen/message"
)

type GetMessageListService struct {
	ctx context.Context
}

// NewGetFeedService new GetFeedService
func NewGetMessageListService(ctx context.Context) *GetMessageListService {
	return &GetMessageListService{ctx: ctx}
}

// GetFeedService.
func (s *GetMessageListService) GetMessageList(req *message.GetMessageListRequest) ([]*message.Message, error) {
	messages, err := db.QueryMessage(s.ctx, req.FromUserId, req.ToUserId, req.PreMsgTime)
	if err != nil {
		return nil, err
	}
	return pack.Messages(messages), nil
}
