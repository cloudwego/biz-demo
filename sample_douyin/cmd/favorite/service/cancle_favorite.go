package service

import (
	"context"
	"mydouyin/cmd/favorite/dal/db"
	"mydouyin/kitex_gen/douyinfavorite"
)

type CancleFavoriteService struct {
	ctx context.Context
}

// NewCrceateVideoService new CreateVideoService
func NewCancleFavoriteService(ctx context.Context) *CancleFavoriteService {
	return &CancleFavoriteService{ctx: ctx}
}

// CreateVideo create video info.
func (s *CancleFavoriteService) CancleFavorite(req *douyinfavorite.FavoriteActionRequest) error {
	return db.CancleFavorite(s.ctx, []*db.Favorite{{
		UserId: req.UserId,
		VideoId: req.VideoId,
	}})
}
