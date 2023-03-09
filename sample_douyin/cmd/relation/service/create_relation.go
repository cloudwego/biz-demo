package service

import (
	"context"
	"mydouyin/cmd/relation/dal/db"
	"mydouyin/kitex_gen/relation"
)

type CreateRelationService struct {
	ctx context.Context
}

func NewCreateRelationService(ctx context.Context) *CreateRelationService {
	return &CreateRelationService{
		ctx: ctx,
	}
}

func (s *CreateRelationService) CreateRelation(req *relation.CreateRelationRequest) error {
	return db.CreateRelation(s.ctx, &db.Relation{
		FollowId:   req.FollowId,
		FollowerId: req.FollowerId,
	})
}
