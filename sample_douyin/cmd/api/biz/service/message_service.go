package service

import (
	"context"
	"mydouyin/cmd/api/biz/apimodel"
	"mydouyin/cmd/api/biz/cache"
	"mydouyin/cmd/api/biz/rpc"
	"mydouyin/kitex_gen/message"
	"mydouyin/pkg/errno"
	"sort"
	"time"
	//"time"
)

type MessageService struct {
	ctx context.Context
}

func NewMessageService(ctx context.Context) *MessageService {
	return &MessageService{
		ctx: ctx,
	}
}

func (s *MessageService) MessageAction(req apimodel.MessageActionRequest, user *apimodel.User) (resp *apimodel.MessageActionResponse, err error) {
	resp = new(apimodel.MessageActionResponse)
	err = cache.MC.CommitCreateMessageCommand(user.UserID, req.ToUserId, req.Content)
	//err = cache.MC.CommitCreateMessageCommandV0(user.UserID, req.ToUserId, req.Content)
	if err != nil {
		return resp, err
	}
	return resp, nil
}

func (s *MessageService) MessageChat(req apimodel.MessageChatRequest, user *apimodel.User) (resp *apimodel.MessageChatResponse, err error) {
	resp = new(apimodel.MessageChatResponse)
	if time.Now().Unix() < req.PreMsgTime {
		//客户端返回的毫秒级时间戳，需要转化成秒级
		// req.PreMsgTime = req.PreMsgTime / 1e3
		req.PreMsgTime, err = cache.MC.GetLastedMsg(user.UserID, req.ToUserId)
		// log.Println(req.PreMsgTime, err)
		if err != nil {
			return
		}
	}
	// log.Println(req.PreMsgTime, time.Now().Unix() > req.PreMsgTime)
	messageList, hit, err := cache.MC.GetMessage(user.UserID, req.ToUserId, req.PreMsgTime)
	if err != nil {
		return
	}
	if hit {
		resp.MessageList = messageList
		return
	}
	rpc_resp_from, err := rpc.GetMessageList(s.ctx, &message.GetMessageListRequest{
		FromUserId: user.UserID,
		ToUserId:   req.ToUserId,
		PreMsgTime: req.PreMsgTime,
	})
	if err != nil {
		return resp, err
	}
	if rpc_resp_from.BaseResp.StatusCode != 0 {
		return nil, errno.NewErrNo(rpc_resp_from.BaseResp.StatusCode, rpc_resp_from.BaseResp.StatusMessage)
	}
	message_list_from := apimodel.PackMessages(rpc_resp_from.MessageList)
	rpc_resp_to, err := rpc.GetMessageList(s.ctx, &message.GetMessageListRequest{
		FromUserId: req.ToUserId,
		ToUserId:   user.UserID,
		PreMsgTime: req.PreMsgTime,
	})
	if err != nil {
		return resp, err
	}
	if rpc_resp_to.BaseResp.StatusCode != 0 {
		return nil, errno.NewErrNo(rpc_resp_to.BaseResp.StatusCode, rpc_resp_to.BaseResp.StatusMessage)
	}
	message_list_to := apimodel.PackMessages(rpc_resp_to.MessageList)
	resp.MessageList = append(message_list_from, message_list_to...)
	sort.Sort(apimodel.MessageSorter(resp.MessageList))
	if len(resp.MessageList) == 0 {
		//表示两个用户之间第一次聊天，没有消息记录，向对应kv缓存中加一条空消息，防止service轮询rpc接口
		cache.MC.SaveMessage(append([]*apimodel.Message{}, &apimodel.Message{FromUserId: user.UserID, ToUserId: req.ToUserId, CreateTime: 0}))
	} else {
		cache.MC.SaveMessage(resp.MessageList)
	}
	return
}
