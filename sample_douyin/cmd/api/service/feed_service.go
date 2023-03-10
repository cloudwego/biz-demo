// Copyright 2023 CloudWeGo Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//

package service

import (
	"context"
	"strconv"
	"time"

	"github.com/cloudwego/biz-demo/sample_douyin/cmd/api/pack"
	"github.com/cloudwego/biz-demo/sample_douyin/cmd/api/rpc"
	douyinapi "github.com/cloudwego/biz-demo/sample_douyin/hertz_gen/douyinapi"
	"github.com/cloudwego/biz-demo/sample_douyin/kitex_gen/douyinfavorite"
	"github.com/cloudwego/biz-demo/sample_douyin/kitex_gen/douyinuser"
	"github.com/cloudwego/biz-demo/sample_douyin/kitex_gen/douyinvideo"
	"github.com/cloudwego/biz-demo/sample_douyin/pkg/errno"
)

type FeedService struct {
	ctx context.Context
}

func NewFeedService(ctx context.Context) *FeedService {
	return &FeedService{
		ctx: ctx,
	}
}

func (s *FeedService) GetFeed(req douyinapi.GetFeedRequest, userId int64) (*douyinapi.GetFeedResponse, error) {
	resp := new(douyinapi.GetFeedResponse)
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
	resp.VideoList = make([]*douyinapi.Video, 0, 30)
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
		var author *douyinapi.User
		if userId != -1 {
			author = pack.PackUserRelation(r.Users[0], userId)
		} else {
			author = pack.PackUser(r.Users[0])
		}
		video := pack.PackVideo(rpcResp.VideoList[i])
		video.Author = author
		video.IsFavorite = isFavorites.IsFavorites[i]
		resp.VideoList = append(resp.VideoList, video)
	}
	resp.NextTime = rpcResp.NextTime
	return resp, nil
}
