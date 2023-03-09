package service

import (
	"context"
	"mydouyin/cmd/relation/dal/db"
	"mydouyin/cmd/relation/pack"
	"mydouyin/kitex_gen/relation"
)

type GetFollowerService struct {
	ctx context.Context
}

func NewGetFollowerService(ctx context.Context) *GetFollowerService {
	return &GetFollowerService{
		ctx: ctx,
	}
}

func (s *GetFollowerService) GetFollower(req *relation.GetFollowerListRequest) ([]int64, error) {
	relations, err := db.GetFollowersByFollow(s.ctx, req.FollowId)
	if err != nil {
		return nil, err
	}
	return pack.Relation2Follower(relations), nil
}
