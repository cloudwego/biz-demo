package service

import (
	"context"
	"mydouyin/cmd/favorite/dal/db"
	"mydouyin/cmd/favorite/pack"
	"mydouyin/kitex_gen/douyinfavorite"
)

type GetListService struct {
	ctx context.Context
}

// NewGetFeedService new GetFeedService
func NewGetListService(ctx context.Context) *GetListService {
	return &GetListService{ctx: ctx}
}

// GetFeedService.
func (s *GetListService) GetList(req *douyinfavorite.GetListRequest) ([]int64, error) {
	favorites, err := db.GetFavoriteList(s.ctx, req.UserId)
	if err != nil {
		return nil, err
	} 
	return pack.FavoriteToVideoids(favorites) ,nil
}
