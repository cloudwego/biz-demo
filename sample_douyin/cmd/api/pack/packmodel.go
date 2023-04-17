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

package pack

import (
	"context"

	"github.com/cloudwego/biz-demo/sample_douyin/cmd/api/rpc"
	douyinapi "github.com/cloudwego/biz-demo/sample_douyin/hertz_gen/douyinapi"
	"github.com/cloudwego/biz-demo/sample_douyin/kitex_gen/douyincomment"
	"github.com/cloudwego/biz-demo/sample_douyin/kitex_gen/douyinfavorite"
	"github.com/cloudwego/biz-demo/sample_douyin/kitex_gen/douyinuser"
	"github.com/cloudwego/biz-demo/sample_douyin/kitex_gen/douyinvideo"
	"github.com/cloudwego/biz-demo/sample_douyin/kitex_gen/message"
	"github.com/cloudwego/biz-demo/sample_douyin/kitex_gen/relation"
	"github.com/cloudwego/biz-demo/sample_douyin/pkg/consts"
	"github.com/cloudwego/biz-demo/sample_douyin/pkg/errno"
)

func PackUser(douyin_user *douyinuser.User) *douyinapi.User {
	return &douyinapi.User{
		ID:              douyin_user.UserId,
		Name:            douyin_user.Username,
		FollowCount:     douyin_user.FollowCount,
		FollowerCount:   douyin_user.FollowerCount,
		Avatar:          douyin_user.Avatar,
		BackgroundImage: douyin_user.BackgroundImage,
		Signature:       douyin_user.Signature,
		TotalFavorited:  douyin_user.TotalFavorited,
		WorkCount:       douyin_user.WorkCount,
		FavoriteCount:   douyin_user.FavoriteCount,
		IsFollow:        false,
	}
}

func PackUserRelation(douyin_user *douyinuser.User, me int64) *douyinapi.User {
	user := PackUser(douyin_user)
	r, err := rpc.ValidIfFollowRequest(context.Background(), &relation.ValidIfFollowRequest{FollowId: user.ID, FollowerId: me})
	if err != nil || r.BaseResp.StatusCode != 0 {
		return user
	}
	user.IsFollow = r.IfFollow
	return user
}

func PackVideo(douyin_video *douyinvideo.Video) *douyinapi.Video {
	return &douyinapi.Video{
		ID:            douyin_video.VideoId,
		PlayURL:       consts.CDNURL + douyin_video.PlayUrl,
		CoverURL:      consts.CDNURL + douyin_video.CoverUrl,
		FavoriteCount: douyin_video.FavoriteCount,
		CommentCount:  douyin_video.CommentCount,
		IsFavorite:    douyin_video.IsFavorite,
		Title:         douyin_video.Title,
		UploadTime:    douyin_video.UploadTime,
	}
}

func PackVideos(douyin_videos []*douyinvideo.Video) []*douyinapi.Video {
	res := make([]*douyinapi.Video, 0, 30)
	for _, douyin_video := range douyin_videos {
		res = append(res, PackVideo(douyin_video))
	}
	return res
}

func PackFriendUser(douyin_user *douyinuser.User) *douyinapi.FriendUser {
	return &douyinapi.FriendUser{
		ID:              douyin_user.UserId,
		Name:            douyin_user.Username,
		FollowCount:     douyin_user.FollowCount,
		FollowerCount:   douyin_user.FollowerCount,
		Avatar:          douyin_user.Avatar,
		BackgroundImage: douyin_user.BackgroundImage,
		Signature:       douyin_user.Signature,
		TotalFavorited:  douyin_user.TotalFavorited,
		WorkCount:       douyin_user.WorkCount,
		FavoriteCount:   douyin_user.FavoriteCount,
		IsFollow:        false,
		Message:         "以成为好友,快来聊天吧",
		MsgType:         1,
	}
}

func PackFriendUsers(douyin_users []*douyinuser.User, me int64) ([]*douyinapi.FriendUser, error) {
	res := make([]*douyinapi.FriendUser, 0, len(douyin_users))
	friendIds := make([]int64, 0, len(douyin_users))
	for _, douyin_user := range douyin_users {
		res = append(res, PackFriendUser(douyin_user))
		friendIds = append(friendIds, douyin_user.UserId)
	}
	rpc_resp, err := rpc.GetFirstMessage(context.Background(), &message.GetFirstMessageRequest{
		Id:        me,
		FriendIds: friendIds,
	})
	if err != nil {
		return nil, err
	}
	if rpc_resp.BaseResp.StatusCode != 0 {
		return nil, errno.NewErrNo(rpc_resp.BaseResp.StatusCode, rpc_resp.BaseResp.StatusMessage)
	}
	if len(rpc_resp.FirstMessageList) != len(res) {
		return nil, errno.QueryErr
	}
	for i, message := range rpc_resp.FirstMessageList {
		if res[i].ID == message.FriendId {
			res[i].Message = message.Message
			res[i].MsgType = message.MsgType
		}
	}
	return res, nil
}

func PackComment(douyin_comment *douyincomment.Comment) *douyinapi.Comment {
	return &douyinapi.Comment{
		ID:         douyin_comment.CommentId,
		Content:    douyin_comment.Content,
		CreateDate: douyin_comment.CreateDate,
	}
}

func PackVFavorite(douyin_favorite *douyinfavorite.Favorite) *douyinapi.Favorite {
	return &douyinapi.Favorite{
		ID:      douyin_favorite.FavoriteId,
		UserID:  douyin_favorite.UserId,
		VideoID: douyin_favorite.VideoId,
	}
}

type FirstMessage struct {
	FriendId int64
	MsgType  int // 0表示接收的 1表示发送的 -1表示为空
	Content  string
}

func PackMessages(rpc_message []*message.Message) []*douyinapi.Message {
	res := make([]*douyinapi.Message, 0, 50)
	for _, res_msg := range rpc_message {
		res = append(res, &douyinapi.Message{
			ID:         res_msg.Id,
			ToUserID:   res_msg.ToUserId,
			FromUserID: res_msg.FromUserId,
			Content:    res_msg.Content,
			CreateTime: res_msg.CreateTime,
		})
	}
	return res
}

type MessageSorter []*douyinapi.Message

func (s MessageSorter) Len() int { return len(s) }

func (s MessageSorter) Less(i, j int) bool { return s[i].ID < s[j].ID }

func (s MessageSorter) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}
