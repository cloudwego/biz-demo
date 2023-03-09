package service

import (
	"context"
	"mydouyin/cmd/api/biz/apimodel"
	"mydouyin/cmd/api/biz/cache"
	"mydouyin/cmd/api/biz/rpc"
	"mydouyin/kitex_gen/douyinfavorite"
	"mydouyin/kitex_gen/douyinuser"
	"mydouyin/kitex_gen/douyinvideo"
	"mydouyin/pkg/errno"
	"strconv"
)

type FavoriteService struct {
	ctx context.Context
}

func NewFavoriteService(ctx context.Context) *FavoriteService {
	return &FavoriteService{
		ctx: ctx,
	}
}

func (s *FavoriteService) FavoriteAction(req apimodel.FavoriteActionRequest, user *apimodel.User) (*apimodel.FavoriteActionResponse, error) {
	resp := new(apimodel.FavoriteActionResponse)
	videoId, err := strconv.Atoi(req.VideoID)
	if err != nil {
		return resp, err
	}
	//异步处理
	err = cache.FC.CommitFavoriteActionCommand(
		user.UserID,
		int64(videoId),
		req.ActionType,
	)
	//同步处理
	// rpc_resp, err := rpc.FavoriteAction(s.ctx, &douyinfavorite.FavoriteActionRequest{
	// 	UserId:     user.UserID,
	// 	VideoId:    int64(videoId),
	// 	ActionType: req.ActionType,
	// })
	if err != nil {
		return resp, err
	}
	// if rpc_resp.BaseResp.StatusCode != 0 {
	// 	return resp, errno.NewErrNo(rpc_resp.BaseResp.StatusCode, rpc_resp.BaseResp.StatusMessage)
	// }
	return resp, nil
}

func (s *FavoriteService) GetFavoriteList(req apimodel.GetFavoriteListRequest, user *apimodel.User) (*apimodel.GetFavoriteListResponse, error) {
	resp := new(apimodel.GetFavoriteListResponse)
	var err error
	userId, err := strconv.Atoi(req.UserId)
	if err != nil {
		return resp, err
	}
	vids, err := rpc.GetFavouriteList(s.ctx, &douyinfavorite.GetListRequest{
		UserId: int64(userId),
	})
	if err != nil {
		return nil, err
	}
	if vids.BaseResp.StatusCode != 0 {
		return nil, errno.NewErrNo(vids.BaseResp.StatusCode, vids.BaseResp.StatusMessage)
	}
	resp.VideoList = make([]apimodel.Video, 0, 50)
	videos, err := rpc.MGetVideo(s.ctx, &douyinvideo.MGetVideoRequest{VideoIds: vids.VideoIds})
	if err != nil {
		return nil, err
	}
	if len(videos.Videos) < 1 {
		return resp, nil
	}
	for _, rpc_video := range videos.Videos {
		r, err := rpc.MGetUser(s.ctx, &douyinuser.MGetUserRequest{UserIds: []int64{rpc_video.Author}})
		if err != nil || r.BaseResp.StatusCode != 0 || len(r.Users) < 1 {
			continue
		}
		author := apimodel.PackUser(r.Users[0])
		video := apimodel.PackVideo(rpc_video)
		video.Author = *author
		resp.VideoList = append(resp.VideoList, *video)
	}
	return resp, nil
}
