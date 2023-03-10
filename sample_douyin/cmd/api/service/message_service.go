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
	"sort"
	"time"

	"github.com/cloudwego/biz-demo/sample_douyin/cmd/api/cache"
	"github.com/cloudwego/biz-demo/sample_douyin/cmd/api/pack"
	"github.com/cloudwego/biz-demo/sample_douyin/cmd/api/rpc"
	douyinapi "github.com/cloudwego/biz-demo/sample_douyin/hertz_gen/douyinapi"
	"github.com/cloudwego/biz-demo/sample_douyin/kitex_gen/message"
	"github.com/cloudwego/biz-demo/sample_douyin/pkg/errno"
)

type MessageService struct {
	ctx context.Context
}

func NewMessageService(ctx context.Context) *MessageService {
	return &MessageService{
		ctx: ctx,
	}
}

func (s *MessageService) MessageAction(req douyinapi.MessageActionRequest, user *douyinapi.User) (resp *douyinapi.MessageActionResponse, err error) {
	resp = new(douyinapi.MessageActionResponse)
	err = cache.MC.CommitCreateMessageCommand(user.ID, req.ToUserID, req.Content)
	//err = cache.MC.CommitCreateMessageCommandV0(user.UserID, req.ToUserId, req.Content)
	if err != nil {
		return resp, err
	}
	return resp, nil
}

func (s *MessageService) MessageChat(req douyinapi.MessageChatRequest, user *douyinapi.User) (resp *douyinapi.MessageChatResponse, err error) {
	resp = new(douyinapi.MessageChatResponse)
	if time.Now().Unix() < req.PreMsgTime {
		//客户端返回的毫秒级时间戳，需要转化成秒级
		// req.PreMsgTime = req.PreMsgTime / 1e3
		req.PreMsgTime, err = cache.MC.GetLastedMsg(user.ID, req.ToUserID)
		// log.Println(req.PreMsgTime, err)
		if err != nil {
			return
		}
	}
	// log.Println(req.PreMsgTime, time.Now().Unix() > req.PreMsgTime)
	messageList, hit, err := cache.MC.GetMessage(user.ID, req.ToUserID, req.PreMsgTime)
	if err != nil {
		return
	}
	if hit {
		resp.MessageList = messageList
		return
	}
	rpc_resp_from, err := rpc.GetMessageList(s.ctx, &message.GetMessageListRequest{
		FromUserId: user.ID,
		ToUserId:   req.ToUserID,
		PreMsgTime: req.PreMsgTime,
	})
	if err != nil {
		return resp, err
	}
	if rpc_resp_from.BaseResp.StatusCode != 0 {
		return nil, errno.NewErrNo(rpc_resp_from.BaseResp.StatusCode, rpc_resp_from.BaseResp.StatusMessage)
	}
	message_list_from := pack.PackMessages(rpc_resp_from.MessageList)
	rpc_resp_to, err := rpc.GetMessageList(s.ctx, &message.GetMessageListRequest{
		FromUserId: req.ToUserID,
		ToUserId:   user.ID,
		PreMsgTime: req.PreMsgTime,
	})
	if err != nil {
		return resp, err
	}
	if rpc_resp_to.BaseResp.StatusCode != 0 {
		return nil, errno.NewErrNo(rpc_resp_to.BaseResp.StatusCode, rpc_resp_to.BaseResp.StatusMessage)
	}
	message_list_to := pack.PackMessages(rpc_resp_to.MessageList)
	resp.MessageList = append(message_list_from, message_list_to...)
	sort.Sort(pack.MessageSorter(resp.MessageList))
	if len(resp.MessageList) == 0 {
		//表示两个用户之间第一次聊天，没有消息记录，向对应kv缓存中加一条空消息，防止service轮询rpc接口
		cache.MC.SaveMessage(append([]*douyinapi.Message{}, &douyinapi.Message{FromUserID: user.ID, ToUserID: req.ToUserID, CreateTime: 0}))
	} else {
		cache.MC.SaveMessage(resp.MessageList)
	}
	return
}
