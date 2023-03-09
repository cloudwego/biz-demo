package service

import (
	"context"
	"mydouyin/cmd/relation/dal/db"
	"mydouyin/cmd/relation/pack"
	"mydouyin/kitex_gen/relation"
)

type GetFollowService struct {
	ctx context.Context
}

func NewGetFollowService(ctx context.Context) *GetFollowService {
	return &GetFollowService{
		ctx: ctx,
	}
}

func (s *GetFollowService) GetFollow(req *relation.GetFollowListRequest) ([]int64, error) {
	relations, err := db.GetFollowsByFollower(s.ctx, req.FollowerId)
	if err != nil {
		return nil, err
	}
	return pack.Relation2Follow(relations), nil
}
