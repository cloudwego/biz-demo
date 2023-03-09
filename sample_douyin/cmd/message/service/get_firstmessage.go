package service

import (
	"context"
	"mydouyin/cmd/message/dal/db"
	"mydouyin/cmd/message/pack"
	"mydouyin/kitex_gen/message"
)

type GetFirstMessageService struct {
	ctx context.Context
}

// NewGetFeedService new GetFeedService
func NewGetFirstMessageService(ctx context.Context) *GetFirstMessageService {
	return &GetFirstMessageService{ctx: ctx}
}

// GetFeedService.
func (s *GetFirstMessageService) GetFirstMessage(req *message.GetFirstMessageRequest) ([]*message.FirstMessage, error) {
	firstmessages, err := db.MGetFirstMessage(s.ctx, req.Id, req.FriendIds)
	if err != nil {
		return nil, err
	}
	return pack.FirstMessages(firstmessages), nil
}
