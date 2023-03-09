package service

import (
	"context"
	"mydouyin/cmd/api/biz/apimodel"
	"mydouyin/cmd/api/biz/cache"
	"mydouyin/cmd/api/biz/rpc"
	"mydouyin/kitex_gen/douyinuser"
	"mydouyin/kitex_gen/message"
	"mydouyin/kitex_gen/relation"
	"mydouyin/pkg/errno"
	"strconv"
)

type RelationService struct {
	ctx context.Context
}

func NewRelationService(ctx context.Context) *RelationService {
	return &RelationService{
		ctx: ctx,
	}
}

func (s *RelationService) RelationAction(req apimodel.RelationActionRequest, user *apimodel.User) (*apimodel.RelationActionResponse, error) {
	resp := new(apimodel.RelationActionResponse)
	userId := user.UserID
	to_user_id, err := strconv.Atoi(req.ToUserId)
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

//获取关注或粉丝列表，option表示操作类型(1：关注列表，2：粉丝列表)
func (s *RelationService) FollowAndFollowerList(req apimodel.FollowAndFollowerListRequest, user *apimodel.User, option int) (*apimodel.FollowAndFollowerListReponse, error) {
	resp := new(apimodel.FollowAndFollowerListReponse)
	var err error
	// users := make([]*apimodel.User, 0)
	userIds := make([]int64, 0)
	userId, err := strconv.Atoi(req.UserId)
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
			if user.UserID == int64(userId) {
				//看的自己的关注列表，IsFollow肯定都是true
				u := apimodel.PackUser(rpc_user)
				u.IsFollow = true
				resp.UserList = append(resp.UserList, u)
			} else {
				//看的别人的关注列表，需要掉rpc查IsFollow
				u := apimodel.PackUserRelation(rpc_user, int64(user.UserID))
				resp.UserList = append(resp.UserList, u)
			}
		case 2:
			u := apimodel.PackUserRelation(rpc_user, int64(user.UserID))
			resp.UserList = append(resp.UserList, u)
		}
	}
	// resp.UserList = users
	return resp, errno.Success
}

func (s *RelationService) FriendList(req apimodel.FriendListRequest) (*apimodel.FriendListReponse, error) {
	resp := new(apimodel.FriendListReponse)
	var err error
	// users := make([]*apimodel.User, 0)
	rpc_resp, err := rpc.GetFriend(s.ctx, &relation.GetFriendRequest{
		MeId: req.UserId,
	})
	if err != nil {
		return resp, err
	}
	if rpc_resp.BaseResp.StatusCode != 0 {
		return resp, errno.NewErrNo(rpc_resp.BaseResp.StatusCode, rpc_resp.BaseResp.StatusMessage)
	}
	if len(rpc_resp.FriendIds) == 0 {
		resp.UserList = make([]*apimodel.FriendUser, 0)
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
	friend_map := make(map[int64]*apimodel.FriendUser, 0)
	for _, rpc_user := range ur.Users {
		u := apimodel.PackFriendUser(rpc_user)
		u.IsFollow = true
		resp.UserList = append(resp.UserList, u)
		friend_map[u.UserID] = u
	}

	//先走缓存，从缓存中查看能不能得到friendlist
	frist_msg_list := cache.MC.GetFirstMessage(req.UserId, rpc_resp.FriendIds)
	missFriendId := make([]int64, 0)
	missFriend := make(map[int64]*apimodel.FriendUser, 0)
	// log.Println("cache中查到的fristlist:")
	for _, frist_msg := range frist_msg_list {
		if frist_msg.MsgType == -1 {
			//miss了
			missFriendId = append(missFriendId, frist_msg.FriendId)
			missFriend[frist_msg.FriendId] = friend_map[frist_msg.FriendId]
		} else {
			friend_map[frist_msg.FriendId].Message = frist_msg.Content
			friend_map[frist_msg.FriendId].MsgType = int64(frist_msg.MsgType)
			// resp.UserList[i].Message = frist_msg.Content
			// resp.UserList[i].MsgType = int64(frist_msg.MsgType)
		}
	}
	// log.Println("miss的friendId", missFriendId)
	if len(missFriend) != 0 {
		//调RPC方法走MySql数据库
		gfm_resp, err := rpc.GetFirstMessage(s.ctx, &message.GetFirstMessageRequest{
			Id:        req.UserId,
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
			if missFriend[message.FriendId].UserID == message.FriendId {
				if message.MsgType == -1 {
					missFriend[message.FriendId].Message = ""
				} else {
					missFriend[message.FriendId].Message = message.Message
				}
				missFriend[message.FriendId].MsgType = message.MsgType
			}
			//miss了就更新缓存
			// log.Println("miss了,更新redis缓存", missFriendId)
			if message.MsgType == 1 {
				cache.MC.SetFirstMessage(&apimodel.Message{
					ToUserId:   message.FriendId,
					FromUserId: req.UserId,
					Content:    message.Message,
				})
			} else if message.MsgType == 0 {
				cache.MC.SetFirstMessage(&apimodel.Message{
					ToUserId:   req.UserId,
					FromUserId: message.FriendId,
					Content:    message.Message,
				})
			} else {
				cache.MC.SetFirstMessage(&apimodel.Message{
					ToUserId:   req.UserId,
					FromUserId: message.FriendId,
					Content:    "",
				})
			}
		}
	}

	return resp, nil
}
