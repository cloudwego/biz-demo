package service

import (
	"context"
	"mydouyin/cmd/relation/dal/db"
	"mydouyin/kitex_gen/relation"
)

type ValidIfFollowService struct {
	ctx context.Context
}

func NewValidIfFollowService(ctx context.Context) *ValidIfFollowService {
	return &ValidIfFollowService{
		ctx: ctx,
	}
}

func (s *ValidIfFollowService) ValidIfFollowFollower(req *relation.ValidIfFollowRequest) (bool, error) {
	return db.ValidRelationIfExist(s.ctx, req.FollowId, req.FollowerId)
}
