package apimodel

import "mime/multipart"

type RegistUserRequest struct {
	Username string `json:"username" query:"username" vd:"len($) > 0"`
	Password string `json:"password" query:"password" vd:"len($) > 0"`
}

type CheckUserRequest struct {
	Username string `json:"username" query:"username" vd:"len($) > 0"`
	Password string `json:"password" query:"password" vd:"len($) > 0"`
}

type GetUserRequest struct {
	UserID string `json:"user_id" query:"user_id"`
	Token  string `json:"token" query:"token"`
}

type GetFeedRequest struct {
	LatestTime string `json:"latest_time" query:"latest_time"`
	Token      string `json:"token" query:"token"`
}

type PublishVideoRequest struct {
	Data  *multipart.FileHeader `json:"data" form:"data"`
	Token string                `json:"token" form:"token"`
	Title string                `json:"title" form:"title"`
}

type GetPublishListRequest struct {
	Token  string `json:"token" query:"token"`
	UserId string `json:"user_id" query:"user_id"`
}

type RelationActionRequest struct {
	Token      string `json:"token" query:"token"`
	ToUserId   string `json:"to_user_id" query:"to_user_id"`
	ActionType string `json:"action_type" query:"action_type"`
}

type FollowAndFollowerListRequest struct {
	UserId string `json:"user_id" query:"user_id"`
	Token  string `json:"token" query:"token"`
}

type FriendListRequest struct {
	UserId int64  `json:"user_id" query:"user_id"`
	Token  string `json:"token" query:"token"`
}

type FavoriteActionRequest struct {
	Token      string `json:"query" query:"token"`
	VideoID    string `json:"video_id" query:"video_id"`
	ActionType string `json:"action_type" query:"action_type"`
}

type GetFavoriteListRequest struct {
	Token  string `json:"query" query:"token"`
	UserId string `json:"user_id" query:"user_id"`
}

type CommentActionRequest struct {
	Token       string `json:"token" query:"token"`
	VideoId     string `json:"video_id" query:"video_id"`
	ActionType  string `json:"action_type" query:"action_type"`
	CommentText string `json:"comment_text" query:"comment_text"`
	CommentId   string `json:"comment_id" query:"comment_id"`
}

type CommentListRequest struct {
	Token   string `json:"token" query:"token"`
	VideoId string `json:"video_id" query:"video_id"`
}

type MessageChatRequest struct {
	Token      string `json:"token" query:"token"`
	ToUserId   int64  `json:"to_user_id" query:"to_user_id"`
	PreMsgTime int64  `json:"pre_msg_time" query:"pre_msg_time"`
}

type MessageActionRequest struct {
	Token      string `json:"token" query:"token"`
	ToUserId   int64  `json:"to_user_id" query:"to_user_id"`
	ActionType int32  `json:"action_type" query:"action_type"`
	Content    string `json:"content" query:"content"`
}
