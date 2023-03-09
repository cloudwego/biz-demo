package service

import (
	"context"
	"mydouyin/cmd/message/dal/db"
	"mydouyin/kitex_gen/message"
)

type CreateMessageService struct {
	ctx context.Context
}

// NewCreateMessageService new CreateMessageService
func NewCreateMessageService(ctx context.Context) *CreateMessageService {
	return &CreateMessageService{ctx: ctx}
}

// CreateMessage create message info.
func (s *CreateMessageService) CreateMessage(req *message.CreateMessageRequest) (id, create_time int64, err error) {
	return db.CreateMessage(s.ctx, []*db.Message{{
		FromUserID: int(req.FromUserId),
		ToUserID:   int(req.ToUserId),
		Content:    req.Content,
	}})
}
