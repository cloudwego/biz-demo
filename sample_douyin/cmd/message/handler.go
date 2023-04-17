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

package main

import (
	"context"

	"github.com/cloudwego/biz-demo/sample_douyin/cmd/message/pack"
	"github.com/cloudwego/biz-demo/sample_douyin/cmd/message/service"
	message "github.com/cloudwego/biz-demo/sample_douyin/kitex_gen/message"
	"github.com/cloudwego/biz-demo/sample_douyin/pkg/errno"
)

// MessageServiceImpl implements the last service interface defined in the IDL.
type MessageServiceImpl struct{}

// CreateMessage implements the MessageServiceImpl interface.
func (s *MessageServiceImpl) CreateMessage(ctx context.Context, req *message.CreateMessageRequest) (resp *message.CreateMessageResponse, err error) {
	// TODO: Your code here...
	resp = new(message.CreateMessageResponse)
	if err = req.IsValid(); err != nil {
		resp.BaseResp = pack.BuildBaseResp(errno.ParamErr)
		return resp, nil
	}
	id, create_time, err := service.NewCreateMessageService(ctx).CreateMessage(req)
	if err != nil {
		resp.BaseResp = pack.BuildBaseResp(err)
		return resp, nil
	}
	resp.BaseResp = pack.BuildBaseResp(errno.Success)
	resp.CreateTime = create_time
	resp.Id = id
	return resp, nil
}

// GetMessageList implements the MessageServiceImpl interface.
func (s *MessageServiceImpl) GetMessageList(ctx context.Context, req *message.GetMessageListRequest) (resp *message.GetMessageListResponse, err error) {
	// TODO: Your code here...
	resp = new(message.GetMessageListResponse)
	var messages []*message.Message
	messages, err = service.NewGetMessageListService(ctx).GetMessageList(req)
	if err != nil {
		resp.BaseResp = pack.BuildBaseResp(err)
		return resp, nil
	}
	resp.MessageList = messages
	resp.BaseResp = pack.BuildBaseResp(errno.Success)
	return resp, nil
}

// GetFirstMessage implements the MessageServiceImpl interface.
func (s *MessageServiceImpl) GetFirstMessage(ctx context.Context, req *message.GetFirstMessageRequest) (resp *message.GetFirstMessageResponse, err error) {
	// TODO: Your code here...
	resp = new(message.GetFirstMessageResponse)
	var firstmessages []*message.FirstMessage
	firstmessages, err = service.NewGetFirstMessageService(ctx).GetFirstMessage(req)
	if err != nil {
		resp.BaseResp = pack.BuildBaseResp(err)
		return resp, nil
	}
	resp.FirstMessageList = firstmessages
	resp.BaseResp = pack.BuildBaseResp(errno.Success)
	return resp, nil
}
