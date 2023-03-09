package service

import (
	"context"
	"mydouyin/cmd/relation/dal/db"
	"mydouyin/kitex_gen/relation"
)

type GetFriendService struct {
	ctx context.Context
}

func NewGetFriendService(ctx context.Context) *GetFriendService {
	return &GetFriendService{
		ctx: ctx,
	}
}

func (s *GetFriendService) GetFriend(req *relation.GetFriendRequest) ([]int64, error) {
	return db.GetFriend(s.ctx, req.MeId)
}
