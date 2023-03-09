package service

import (
	"context"
	"mydouyin/cmd/favorite/dal/db"
	"mydouyin/kitex_gen/douyinfavorite"
)

type CreateFavoriteService struct {
	ctx context.Context
}

// NewCreateVideoService new CreateVideoService
func NewCreateFavoriteService(ctx context.Context) *CreateFavoriteService {
	return &CreateFavoriteService{ctx: ctx}
}

// CreateVideo create video info.
func (s *CreateFavoriteService) CreateFavorite(req *douyinfavorite.FavoriteActionRequest) error {
	return db.CreateFavorite(s.ctx, []*db.Favorite{{
		UserId: req.UserId,
		VideoId: req.VideoId,
	}})
}
