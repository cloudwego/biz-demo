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
	"github.com/cloudwego/biz-demo/sample_douyin/kitex_gen/douyinuser"
	"github.com/cloudwego/biz-demo/sample_douyin/kitex_gen/message"
	"github.com/cloudwego/biz-demo/sample_douyin/kitex_gen/relation"
	"github.com/cloudwego/biz-demo/sample_douyin/pkg/errno"
)

type RelationService struct {
	ctx context.Context
}

func NewRelationService(ctx context.Context) *RelationService {
	return &RelationService{
		ctx: ctx,
	}
}

func (s *RelationService) RelationAction(req douyinapi.RelationActionRequest, user *douyinapi.User) (*douyinapi.RelationActionResponse, error) {
	resp := new(douyinapi.RelationActionResponse)
	userId := user.ID
	to_user_id, err := strconv.Atoi(req.ToUserID)
	if err != nil {
		return resp, errno.ParamErr
	}

	switch req.ActionType {
	case "1":
		rpc_resp, err := rpc.CreateRelation(s.ctx, &relation.CreateRelationRequest{
			FollowId:   int64(to_user_id),
			FollowerId: userId,
		})
		if err != nil {
			return resp, err
		}
		if rpc_resp.StatusCode != 0 {
			return resp, errno.NewErrNo(rpc_resp.StatusCode, rpc_resp.StatusMessage)
		}
	case "2":
		rpc_resp, err := rpc.DeleteRelation(s.ctx, &relation.DeleteRelationRequest{
			FollowId:   int64(to_user_id),
			FollowerId: userId,
		})
		if err != nil {
			return resp, err
		}
		if rpc_resp.StatusCode != 0 {
			return resp, errno.NewErrNo(rpc_resp.StatusCode, rpc_resp.StatusMessage)
		}
	default:
		err = errno.ParamErr
		return resp, err
	}
	return resp, nil
}

// 获取关注或粉丝列表，option表示操作类型(1：关注列表，2：粉丝列表)
func (s *RelationService) FollowAndFollowerList(req douyinapi.FollowAndFollowerListRequest, user *douyinapi.User, option int) (*douyinapi.FollowAndFollowerListResponse, error) {
	resp := new(douyinapi.FollowAndFollowerListResponse)
	var err error
	// users := make([]*douyinapi.User, 0)
	userIds := make([]int64, 0)
	userId, err := strconv.Atoi(req.UserID)
	if err != nil {
		return resp, err
	}
	switch option {
	case 1:
		rpc_resp, err := rpc.GetFollow(s.ctx, &relation.GetFollowListRequest{FollowerId: int64(userId)})
		if err != nil {
			return resp, err
		}
		if rpc_resp.BaseResp.StatusCode != 0 {
			return resp, errno.NewErrNo(rpc_resp.BaseResp.StatusCode, rpc_resp.BaseResp.StatusMessage)
		}
		if len(rpc_resp.FollowIds) < 1 {
			return resp, nil
		}
		userIds = rpc_resp.FollowIds
	case 2:
		rpc_resp, err := rpc.GetFollower(s.ctx, &relation.GetFollowerListRequest{FollowId: int64(userId)})
		if err != nil {
			return resp, err
		}
		if rpc_resp.BaseResp.StatusCode != 0 {
			return resp, errno.NewErrNo(rpc_resp.BaseResp.StatusCode, rpc_resp.BaseResp.StatusMessage)
		}
		if len(rpc_resp.FollowerIds) < 1 {
			return resp, nil
		}
		userIds = rpc_resp.FollowerIds
	}
	ur, err := rpc.MGetUser(s.ctx, &douyinuser.MGetUserRequest{
		UserIds: userIds,
	})
	if err != nil {
		return nil, err
	}
	if ur.BaseResp.StatusCode != 0 {
		return nil, errno.NewErrNo(ur.BaseResp.StatusCode, ur.BaseResp.StatusMessage)
	}
	for _, rpc_user := range ur.Users {
		switch option {
		case 1:
			if user.ID == int64(userId) {
				// 看的自己的关注列表，IsFollow肯定都是true
				u := pack.PackUser(rpc_user)
				u.IsFollow = true
				resp.UserList = append(resp.UserList, u)
			} else {
				// 看的别人的关注列表，需要掉rpc查IsFollow
				u := pack.PackUserRelation(rpc_user, int64(user.ID))
				resp.UserList = append(resp.UserList, u)
			}
		case 2:
			u := pack.PackUserRelation(rpc_user, int64(user.ID))
			resp.UserList = append(resp.UserList, u)
		}
	}
	// resp.UserList = users
	return resp, errno.Success
}

func (s *RelationService) FriendList(req douyinapi.FriendListRequest) (*douyinapi.FriendListResponse, error) {
	resp := new(douyinapi.FriendListResponse)
	var err error
	// users := make([]*douyinapi.User, 0)
	rpc_resp, err := rpc.GetFriend(s.ctx, &relation.GetFriendRequest{
		MeId: req.UserID,
	})
	if err != nil {
		return resp, err
	}
	if rpc_resp.BaseResp.StatusCode != 0 {
		return resp, errno.NewErrNo(rpc_resp.BaseResp.StatusCode, rpc_resp.BaseResp.StatusMessage)
	}
	if len(rpc_resp.FriendIds) == 0 {
		resp.UserList = make([]*douyinapi.FriendUser, 0)
		return resp, nil
	}
	ur, err := rpc.MGetUser(s.ctx, &douyinuser.MGetUserRequest{
		UserIds: rpc_resp.FriendIds,
	})
	if err != nil {
		return resp, err
	}
	if ur.BaseResp.StatusCode != 0 {
		return resp, errno.NewErrNo(ur.BaseResp.StatusCode, ur.BaseResp.StatusMessage)
	}
	friend_map := make(map[int64]*douyinapi.FriendUser, 0)
	for _, rpc_user := range ur.Users {
		u := pack.PackFriendUser(rpc_user)
		u.IsFollow = true
		resp.UserList = append(resp.UserList, u)
		friend_map[u.ID] = u
	}

	// 先走缓存，从缓存中查看能不能得到friendlist
	first_msg_list := cache.MC.GetFirstMessage(req.UserID, rpc_resp.FriendIds)
	missFriendId := make([]int64, 0)
	missFriend := make(map[int64]*douyinapi.FriendUser, 0)
	// log.Println("cache中查到的fristlist:")
	for _, first_msg := range first_msg_list {
		if first_msg.MsgType == -1 {
			// miss了
			missFriendId = append(missFriendId, first_msg.FriendId)
			missFriend[first_msg.FriendId] = friend_map[first_msg.FriendId]
		} else {
			friend_map[first_msg.FriendId].Message = first_msg.Content
			friend_map[first_msg.FriendId].MsgType = int64(first_msg.MsgType)
			// resp.UserList[i].Message = first_msg.Content
			// resp.UserList[i].MsgType = int64(first_msg.MsgType)
		}
	}
	// log.Println("miss的friendId", missFriendId)
	if len(missFriend) != 0 {
		// 调RPC方法走MySql数据库
		gfm_resp, err := rpc.GetFirstMessage(s.ctx, &message.GetFirstMessageRequest{
			Id:        req.UserID,
			FriendIds: missFriendId,
		})
		if err != nil {
			return resp, err
		}
		if gfm_resp.BaseResp.StatusCode != 0 {
			return resp, errno.NewErrNo(gfm_resp.BaseResp.StatusCode, gfm_resp.BaseResp.StatusMessage)
		}
		if len(gfm_resp.FirstMessageList) != len(missFriend) {
			return resp, errno.QueryErr
		}
		for _, message := range gfm_resp.FirstMessageList {
			// log.Printf("%v:%v\n", *missFriend[message.FriendId], *message)
			if missFriend[message.FriendId].ID == message.FriendId {
				if message.MsgType == -1 {
					missFriend[message.FriendId].Message = ""
				} else {
					missFriend[message.FriendId].Message = message.Message
				}
				missFriend[message.FriendId].MsgType = message.MsgType
			}
			// miss了就更新缓存
			// log.Println("miss了,更新redis缓存", missFriendId)
			if message.MsgType == 1 {
				err = cache.MC.SetFirstMessage(&douyinapi.Message{
					ToUserID:   message.FriendId,
					FromUserID: req.UserID,
					Content:    message.Message,
				})
				if err != nil {
					return nil, err
				}
			} else if message.MsgType == 0 {
				err = cache.MC.SetFirstMessage(&douyinapi.Message{
					ToUserID:   req.UserID,
					FromUserID: message.FriendId,
					Content:    message.Message,
				})
				if err != nil {
					return nil, err
				}
			} else {
				err = cache.MC.SetFirstMessage(&douyinapi.Message{
					ToUserID:   req.UserID,
					FromUserID: message.FriendId,
					Content:    "",
				})
				if err != nil {
					return nil, err
				}
			}
		}
	}

	return resp, nil
}
