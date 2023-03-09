package service

import (
	"context"
	"mydouyin/cmd/api/biz/apimodel"
	"mydouyin/cmd/api/biz/rpc"
	"mydouyin/kitex_gen/douyinfavorite"
	"mydouyin/kitex_gen/douyinuser"
	"mydouyin/kitex_gen/douyinvideo"
	"mydouyin/pkg/errno"
	"strconv"
	"time"
)

type FeedService struct {
	ctx context.Context
}

func NewFeedService(ctx context.Context) *FeedService {
	return &FeedService{
		ctx: ctx,
	}
}

func (s *FeedService) GetFeed(req apimodel.GetFeedRequest, userId int64) (*apimodel.GetFeedResponse, error) {
	resp := new(apimodel.GetFeedResponse)
	var err error
	if req.LatestTime == "" {
		req.LatestTime = strconv.FormatInt(time.Now().Unix(), 10)
	}
	if len(req.LatestTime) > 10 {
		req.LatestTime = string([]rune(req.LatestTime[0 : len(req.LatestTime)-3]))
	}
	rpcResp, err := rpc.GetFeed(s.ctx, &douyinvideo.GetFeedRequest{
		LatestTime: req.LatestTime,
		UserId:     userId,
	})
	if err != nil {
		resp.NextTime = time.Now().Unix()
		return resp, err
	}
	if rpcResp.BaseResp.StatusCode != 0 {
		resp.NextTime = time.Now().Unix()
		return resp, errno.NewErrNo(rpcResp.BaseResp.StatusCode, rpcResp.BaseResp.StatusMessage)
	}
	resp.VideoList = make([]apimodel.Video, 0, 30)
	favorites := make([]*douyinfavorite.Favorite, 0)
	for _, rpcVideo := range rpcResp.VideoList {
		favorite := new(douyinfavorite.Favorite)
		favorite.UserId = userId
		favorite.VideoId = rpcVideo.VideoId
		favorites = append(favorites, favorite)
	}
	isFavorites, err := rpc.GetIsFavorite(s.ctx, &douyinfavorite.GetIsFavoriteRequest{FavoriteList: favorites})

	if err != nil {
		resp.NextTime = time.Now().Unix()
		return resp, err
	}

	if len(rpcResp.VideoList) != len(isFavorites.IsFavorites) {
		resp.NextTime = time.Now().Unix()
		return resp, errno.ServiceErr
	}

	for i := 0; i < len(rpcResp.VideoList); i++ {
		r, err := rpc.MGetUser(s.ctx, &douyinuser.MGetUserRequest{UserIds: []int64{rpcResp.VideoList[i].Author}})
		if err != nil || r.BaseResp.StatusCode != 0 || len(r.Users) < 1 {
			continue
		}
		var author *apimodel.User
		if userId != -1 {
			author = apimodel.PackUserRelation(r.Users[0], userId)
		} else {
			author = apimodel.PackUser(r.Users[0])
		}
		video := apimodel.PackVideo(rpcResp.VideoList[i])
		video.Author = *author
		video.IsFavorite = isFavorites.IsFavorites[i]
		resp.VideoList = append(resp.VideoList, *video)
	}
	resp.NextTime = rpcResp.NextTime
	return resp, nil
}
