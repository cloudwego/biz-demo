package service

import (
	"context"
	"mydouyin/cmd/message/dal/db"
	"mydouyin/cmd/message/pack"
	"mydouyin/kitex_gen/message"
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
