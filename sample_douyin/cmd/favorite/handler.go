package main

import (
	"context"
	"mydouyin/cmd/favorite/pack"
	"mydouyin/cmd/favorite/service"
	douyinfavorite "mydouyin/kitex_gen/douyinfavorite"
	"mydouyin/pkg/errno"
)

// FavoriteServiceImpl implements the last service interface defined in the IDL.
type FavoriteServiceImpl struct{}

// FavoriteAction implements the FavoriteServiceImpl interface.
func (s *FavoriteServiceImpl) FavoriteAction(ctx context.Context, req *douyinfavorite.FavoriteActionRequest) (resp *douyinfavorite.FavoriteActionResponse, err error) {
	resp = new(douyinfavorite.FavoriteActionResponse)
	if req.ActionType == "2" {
		err = service.NewCancleFavoriteService(ctx).CancleFavorite(req)
		if err != nil {
			resp.BaseResp = pack.BuildBaseResp(err)
			return resp, nil
		}
		resp.BaseResp = pack.BuildBaseResp(errno.Success)
		return resp, nil
	} else {
		err = service.NewCreateFavoriteService(ctx).CreateFavorite(req)
		if err != nil {
			resp.BaseResp = pack.BuildBaseResp(err)
			return resp, nil
		}
		resp.BaseResp = pack.BuildBaseResp(errno.Success)
		return resp, nil
	}
}

// GetList implements the FavoriteServiceImpl interface.
func (s *FavoriteServiceImpl) GetList(ctx context.Context, req *douyinfavorite.GetListRequest) (resp *douyinfavorite.GetListResponse, err error) {
	resp = new(douyinfavorite.GetListResponse)
	var vids []int64
	vids, err = service.NewGetListService(ctx).GetList(req)
	if err != nil {
		resp.BaseResp = pack.BuildBaseResp(err)
		return resp, nil
	}
	resp.VideoIds = vids
	resp.BaseResp = pack.BuildBaseResp(errno.Success)
	return resp, nil
}

// GetIsFavorite implements the FavoriteServiceImpl interface.
func (s *FavoriteServiceImpl) GetIsFavorite(ctx context.Context, req *douyinfavorite.GetIsFavoriteRequest) (resp *douyinfavorite.GetIsFavoriteResponse, err error) {
	resp = new(douyinfavorite.GetIsFavoriteResponse)
	var isfavorites []bool
	isfavorites, err = service.NewGetIsFavoriteService(ctx).GetIsFavorite(req)
	if err != nil {
		resp.BaseResp = pack.BuildBaseResp(err)
		return resp, nil
	}
	resp.IsFavorites = isfavorites
	resp.BaseResp = pack.BuildBaseResp(errno.Success)
	return resp, nil
}
