package apimodel

import (
	"context"
	"mydouyin/cmd/api/biz/rpc"
	"mydouyin/kitex_gen/douyincomment"
	"mydouyin/kitex_gen/douyinfavorite"
	"mydouyin/kitex_gen/douyinuser"
	"mydouyin/kitex_gen/douyinvideo"
	"mydouyin/kitex_gen/message"
	"mydouyin/kitex_gen/relation"
	"mydouyin/pkg/consts"
	"mydouyin/pkg/errno"
)

type User struct {
	UserID          int64  `form:"user_id" json:"id" query:"user_id"`
	Username        string `form:"name" json:"name" query:"name"`
	FollowCount     int64  `form:"follow_count" json:"follow_count" query:"follow_count"`
	FollowerCount   int64  `form:"follower_count" json:"follower_count" query:"follower_count"`
	IsFollow        bool   `form:"is_follow" json:"is_follow" query:"is_follow"`
	Avatar          string `form:"avatar" json:"avatar" query:"avatar"`
	BackgroundImage string `form:"background_image" json:"background_image" query:"background_image"`
	Signature       string `form:"signature" json:"signature" query:"signature"`
	TotalFavorited  int64  `form:"total_favorited" json:"total_favorited" query:"total_favoried"`
	WorkCount       int64  `form:"work_count" json:"work_count" query:"work_count"`
	FavoriteCount   int64  `form:"favorite_count" json:"favorite_count" query:"favorite_count"`
}

func PackUser(douyin_user *douyinuser.User) *User {
	return &User{
		UserID:          douyin_user.UserId,
		Username:        douyin_user.Username,
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

func PackUserRelation(douyin_user *douyinuser.User, me int64) *User {
	user := PackUser(douyin_user)
	r, err := rpc.ValidIfFollowRequest(context.Background(), &relation.ValidIfFollowRequest{FollowId: user.UserID, FollowerId: me})
	if err != nil || r.BaseResp.StatusCode != 0 {
		return user
	}
	user.IsFollow = r.IfFollow
	return user
}

type Video struct {
	VideoID       int64  `form:"id" json:"id" query:"id"`
	Author        User   `form:"author" json:"author" query:"author"`
	PlayUrl       string `form:"play_url" json:"play_url" query:"play_url"`
	CoverUrl      string `form:"cover_url" json:"cover_url" query:"cover_url"`
	FavoriteCount int    `form:"favorite_count" json:"favorite_count" query:"favorite_count"`
	CommentCount  int    `form:"comment_count" json:"comment_count" query:"comment_count"`
	IsFavorite    bool   `form:"is_favorite" json:"is_favorite" query:"is_favorite"`
	Title         string `form:"title" json:"title" query:"title"`
	UploadTime    string `form:"upload_time" json:"upload_time" query:"upload_time"`
}

func PackVideo(douyin_video *douyinvideo.Video) *Video {
	return &Video{
		VideoID: douyin_video.VideoId,
		// Author:        douyin_video.Author,
		PlayUrl:       consts.CDNURL + douyin_video.PlayUrl,
		CoverUrl:      consts.CDNURL + douyin_video.CoverUrl,
		FavoriteCount: int(douyin_video.FavoriteCount),
		CommentCount:  int(douyin_video.CommentCount),
		IsFavorite:    douyin_video.IsFavorite,
		Title:         douyin_video.Title,
		UploadTime:    douyin_video.UploadTime,
	}
}

func PackVideos(douyin_videos []*douyinvideo.Video) []*Video {
	res := make([]*Video, 0, 30)
	for _, douyin_video := range douyin_videos {
		res = append(res, PackVideo(douyin_video))
	}
	return res
}

type FriendUser struct {
	User
	Message string `form:"message" json:"message" query:"message"`
	MsgType int64  `form:"msgType" json:"msgType" query:"msgType"`
}

func PackFriendUser(douyin_user *douyinuser.User) *FriendUser {
	return &FriendUser{
		*PackUser(douyin_user),
		"有新消息",
		1,
	}
}

func PackFriendUsers(douyin_users []*douyinuser.User, me int64) ([]*FriendUser, error) {
	res := make([]*FriendUser, 0, len(douyin_users))
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
		if res[i].UserID == message.FriendId {
			res[i].Message = message.Message
			res[i].MsgType = message.MsgType
		}
	}
	return res, nil
}

type Comment struct {
	CommentID  int64  `form:"id" json:"id" query:"id"`
	Commentor  User   `form:"user" json:"user" query:"user"`
	Content    string `form:"content" json:"content" query:"content"`
	CreateDate string `form:"create_date" json:"create_date" query:"create_date"`
}

func PackComment(douyin_comment *douyincomment.Comment) *Comment {
	return &Comment{
		CommentID: douyin_comment.CommentId,
		// User: douyin_comment.User
		Content:    douyin_comment.Content,
		CreateDate: douyin_comment.CreateDate,
	}
}

type Favorite struct {
	FavoriteID int64 `form:"id" json:"id" query:"id"`
	UserID     int64 `form:"user_id" json:"user_id" query:"user_id"`
	VideoID    int64 `form:"video_id" json:"video_id" query:"video_id"`
}

func PackVFavorite(douyin_favorite *douyinfavorite.Favorite) *Favorite {
	return &Favorite{
		FavoriteID: douyin_favorite.FavoriteId,
		UserID:     douyin_favorite.UserId,
		VideoID:    douyin_favorite.VideoId,
	}
}

type Message struct {
	ID         int64  `form:"id" json:"id" query:"id"`
	ToUserId   int64  `form:"to_user_id" json:"to_user_id" query:"to_user_id"`
	FromUserId int64  `form:"from_user_id" json:"from_user_id" query:"from_user_id"`
	Content    string `form:"content" json:"content" query:"content"`
	CreateTime int64  `form:"create_time" json:"create_time" query:"create_time"`
}

type FristMessage struct {
	FriendId int64
	MsgType  int //0表示接收的 1表示发送的 -1表示为空
	Content  string
}

func PackMessages(rpc_message []*message.Message) []*Message {
	res := make([]*Message, 0, 50)
	for _, res_msg := range rpc_message {
		res = append(res, &Message{
			ID:         res_msg.Id,
			ToUserId:   res_msg.ToUserId,
			FromUserId: res_msg.FromUserId,
			Content:    res_msg.Content,
			CreateTime: res_msg.CreateTime,
		})
	}
	return res
}

type MessageSorter []*Message

func (s MessageSorter) Len() int { return len(s) }

func (s MessageSorter) Less(i, j int) bool { return s[i].ID < s[j].ID }

func (s MessageSorter) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}
