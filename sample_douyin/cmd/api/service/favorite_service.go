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

	"github.com/cloudwego/biz-demo/sample_douyin/cmd/api/cache"
	"github.com/cloudwego/biz-demo/sample_douyin/cmd/api/pack"
	"github.com/cloudwego/biz-demo/sample_douyin/cmd/api/rpc"
	douyinapi "github.com/cloudwego/biz-demo/sample_douyin/hertz_gen/douyinapi"
	"github.com/cloudwego/biz-demo/sample_douyin/kitex_gen/douyinfavorite"
	"github.com/cloudwego/biz-demo/sample_douyin/kitex_gen/douyinuser"
	"github.com/cloudwego/biz-demo/sample_douyin/kitex_gen/douyinvideo"
	"github.com/cloudwego/biz-demo/sample_douyin/pkg/errno"
)

type FavoriteService struct {
	ctx context.Context
}

func NewFavoriteService(ctx context.Context) *FavoriteService {
	return &FavoriteService{
		ctx: ctx,
	}
}

func (s *FavoriteService) FavoriteAction(req douyinapi.FavoriteActionRequest, user *douyinapi.User) (*douyinapi.FavoriteActionResponse, error) {
	resp := new(douyinapi.FavoriteActionResponse)
	videoId, err := strconv.Atoi(req.VideoID)
	if err != nil {
		return resp, err
	}
	// 异步处理
	err = cache.FC.CommitFavoriteActionCommand(
		user.ID,
		int64(videoId),
		req.ActionType,
	)
	// 同步处理
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

func (s *FavoriteService) GetFavoriteList(req douyinapi.GetFavoriteListRequest, user *douyinapi.User) (*douyinapi.GetFavoriteListResponse, error) {
	resp := new(douyinapi.GetFavoriteListResponse)
	var err error
	userId, err := strconv.Atoi(req.UserID)
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
	resp.VideoList = make([]*douyinapi.Video, 0, 50)
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
		author := pack.PackUser(r.Users[0])
		video := pack.PackVideo(rpc_video)
		video.Author = author
		resp.VideoList = append(resp.VideoList, video)
	}
	return resp, nil
}
