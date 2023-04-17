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

package main

import (
	"context"

	"github.com/cloudwego/biz-demo/sample_douyin/cmd/favorite/pack"
	"github.com/cloudwego/biz-demo/sample_douyin/cmd/favorite/service"
	douyinfavorite "github.com/cloudwego/biz-demo/sample_douyin/kitex_gen/douyinfavorite"
	"github.com/cloudwego/biz-demo/sample_douyin/pkg/errno"
)

// FavoriteServiceImpl implements the last service interface defined in the IDL.
type FavoriteServiceImpl struct{}

// FavoriteAction implements the FavoriteServiceImpl interface.
func (s *FavoriteServiceImpl) FavoriteAction(ctx context.Context, req *douyinfavorite.FavoriteActionRequest) (resp *douyinfavorite.FavoriteActionResponse, err error) {
	resp = new(douyinfavorite.FavoriteActionResponse)
	if req.ActionType == "2" {
		err = service.NewCancelFavoriteService(ctx).CancelFavorite(req)
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
