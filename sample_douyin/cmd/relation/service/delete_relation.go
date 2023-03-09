package service

import (
	"context"
	"mydouyin/cmd/relation/dal/db"
	"mydouyin/kitex_gen/relation"
)

type DeleteRelationService struct {
	ctx context.Context
}

func NewDeleteRelationService(ctx context.Context) *DeleteRelationService {
	return &DeleteRelationService{
		ctx: ctx,
	}
}

func (s *DeleteRelationService) DeleteRelation(req *relation.DeleteRelationRequest) error {
	return db.DeleteRelation(s.ctx, req.FollowId, req.FollowerId)
}
