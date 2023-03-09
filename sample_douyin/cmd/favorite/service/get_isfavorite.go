package service
import (
	"context"
	"mydouyin/cmd/favorite/dal/db"
	"mydouyin/cmd/favorite/pack"
	"mydouyin/kitex_gen/douyinfavorite"
)

type GetIsFavoriteService struct {
	ctx context.Context
}

// NewGetFeedService new GetFeedService
func NewGetIsFavoriteService(ctx context.Context) *GetIsFavoriteService {
	return &GetIsFavoriteService{ctx: ctx}
}

// GetFeedService.
func (s *GetIsFavoriteService) GetIsFavorite(req *douyinfavorite.GetIsFavoriteRequest) ([]bool, error) {
	return db.QueryFavoriteById(s.ctx, pack.Favorites(req.FavoriteList))
}
