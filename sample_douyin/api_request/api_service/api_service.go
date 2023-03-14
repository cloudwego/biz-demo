// Code generated by hertz generator.

package api_service

import (
	"context"
	"fmt"

	douyinapi "github.com/cloudwego/biz-demo/sample_douyin/hertz_gen/douyinapi"
	"github.com/cloudwego/hertz/pkg/common/config"
	"github.com/cloudwego/hertz/pkg/protocol"
)

// unused protection
var (
	_ = fmt.Formatter(nil)
)

type Client interface {
	RegistUser(context context.Context, req *douyinapi.RegistUserRequest, reqOpt ...config.RequestOption) (resp *douyinapi.RegistUserResponse, rawResponse *protocol.Response, err error)

	CheckUser(context context.Context, req *douyinapi.CheckUserRequest, reqOpt ...config.RequestOption) (resp *douyinapi.CheckUserResponse, rawResponse *protocol.Response, err error)

	GetUser(context context.Context, req *douyinapi.GetUserRequest, reqOpt ...config.RequestOption) (resp *douyinapi.GetUserResponse, rawResponse *protocol.Response, err error)

	GetFeed(context context.Context, req *douyinapi.GetFeedRequest, reqOpt ...config.RequestOption) (resp *douyinapi.GetFeedResponse, rawResponse *protocol.Response, err error)

	GetPublishList(context context.Context, req *douyinapi.GetPublishListRequest, reqOpt ...config.RequestOption) (resp *douyinapi.GetPublishListResponse, rawResponse *protocol.Response, err error)

	PublishVideo(context context.Context, req *douyinapi.PublishVideoRequest, reqOpt ...config.RequestOption) (resp *douyinapi.PublishVideoResponse, rawResponse *protocol.Response, err error)

	FavoriteAction(context context.Context, req *douyinapi.FavoriteActionRequest, reqOpt ...config.RequestOption) (resp *douyinapi.FavoriteActionResponse, rawResponse *protocol.Response, err error)

	GetFavoriteList(context context.Context, req *douyinapi.GetFavoriteListRequest, reqOpt ...config.RequestOption) (resp *douyinapi.GetFavoriteListResponse, rawResponse *protocol.Response, err error)

	CommentAction(context context.Context, req *douyinapi.CommentActionRequest, reqOpt ...config.RequestOption) (resp *douyinapi.CommentActionResponse, rawResponse *protocol.Response, err error)

	CommentList(context context.Context, req *douyinapi.CommentListRequest, reqOpt ...config.RequestOption) (resp *douyinapi.CommentListResponse, rawResponse *protocol.Response, err error)

	RelationAction(context context.Context, req *douyinapi.RelationActionRequest, reqOpt ...config.RequestOption) (resp *douyinapi.RelationActionResponse, rawResponse *protocol.Response, err error)

	FollowList(context context.Context, req *douyinapi.FollowAndFollowerListRequest, reqOpt ...config.RequestOption) (resp *douyinapi.FollowAndFollowerListResponse, rawResponse *protocol.Response, err error)

	FollowerList(context context.Context, req *douyinapi.FollowAndFollowerListRequest, reqOpt ...config.RequestOption) (resp *douyinapi.FollowAndFollowerListResponse, rawResponse *protocol.Response, err error)

	FriendList(context context.Context, req *douyinapi.FriendListRequest, reqOpt ...config.RequestOption) (resp *douyinapi.FriendListResponse, rawResponse *protocol.Response, err error)

	MessageChat(context context.Context, req *douyinapi.MessageChatRequest, reqOpt ...config.RequestOption) (resp *douyinapi.MessageChatResponse, rawResponse *protocol.Response, err error)

	MessageAction(context context.Context, req *douyinapi.MessageActionRequest, reqOpt ...config.RequestOption) (resp *douyinapi.MessageActionResponse, rawResponse *protocol.Response, err error)
}

type ApiServiceClient struct {
	client *cli
}

func NewApiServiceClient(hostUrl string, ops ...Option) (Client, error) {
	opts := getOptions(append(ops, withHostUrl(hostUrl))...)
	cli, err := newClient(opts)
	if err != nil {
		return nil, err
	}
	return &ApiServiceClient{
		client: cli,
	}, nil
}

func (s *ApiServiceClient) RegistUser(context context.Context, req *douyinapi.RegistUserRequest, reqOpt ...config.RequestOption) (resp *douyinapi.RegistUserResponse, rawResponse *protocol.Response, err error) {
	httpResp := &douyinapi.RegistUserResponse{}
	ret, err := s.client.r().
		setContext(context).
		setQueryParams(map[string]interface{}{
			"username": req.GetUsername(),
			"password": req.GetPassword(),
		}).
		setPathParams(map[string]string{}).
		setHeaders(map[string]string{}).
		setFormParams(map[string]string{}).
		setFormFileParams(map[string]string{}).
		setBodyParam(req).
		setRequestOption(reqOpt...).
		setResult(httpResp).
		execute("POST", "/douyin/user/register/")
	if err != nil {
		return nil, nil, err
	}

	resp = httpResp
	rawResponse = ret.rawResponse
	return resp, rawResponse, nil
}

func (s *ApiServiceClient) CheckUser(context context.Context, req *douyinapi.CheckUserRequest, reqOpt ...config.RequestOption) (resp *douyinapi.CheckUserResponse, rawResponse *protocol.Response, err error) {
	httpResp := &douyinapi.CheckUserResponse{}
	ret, err := s.client.r().
		setContext(context).
		setQueryParams(map[string]interface{}{
			"username": req.GetUsername(),
			"password": req.GetPassword(),
		}).
		setPathParams(map[string]string{}).
		setHeaders(map[string]string{}).
		setFormParams(map[string]string{}).
		setFormFileParams(map[string]string{}).
		setBodyParam(req).
		setRequestOption(reqOpt...).
		setResult(httpResp).
		execute("POST", "/douyin/user/login/")
	if err != nil {
		return nil, nil, err
	}

	resp = httpResp
	rawResponse = ret.rawResponse
	return resp, rawResponse, nil
}

func (s *ApiServiceClient) GetUser(context context.Context, req *douyinapi.GetUserRequest, reqOpt ...config.RequestOption) (resp *douyinapi.GetUserResponse, rawResponse *protocol.Response, err error) {
	httpResp := &douyinapi.GetUserResponse{}
	ret, err := s.client.r().
		setContext(context).
		setQueryParams(map[string]interface{}{
			"user_id": req.GetUserID(),
			"token":   req.GetToken(),
		}).
		setPathParams(map[string]string{}).
		setHeaders(map[string]string{}).
		setFormParams(map[string]string{}).
		setFormFileParams(map[string]string{}).
		setBodyParam(req).
		setRequestOption(reqOpt...).
		setResult(httpResp).
		execute("GET", "/douyin/user/")
	if err != nil {
		return nil, nil, err
	}

	resp = httpResp
	rawResponse = ret.rawResponse
	return resp, rawResponse, nil
}

func (s *ApiServiceClient) GetFeed(context context.Context, req *douyinapi.GetFeedRequest, reqOpt ...config.RequestOption) (resp *douyinapi.GetFeedResponse, rawResponse *protocol.Response, err error) {
	httpResp := &douyinapi.GetFeedResponse{}
	ret, err := s.client.r().
		setContext(context).
		setQueryParams(map[string]interface{}{
			"latest_time": req.GetLatestTime(),
			"token":       req.GetToken(),
		}).
		setPathParams(map[string]string{}).
		setHeaders(map[string]string{}).
		setFormParams(map[string]string{}).
		setFormFileParams(map[string]string{}).
		setBodyParam(req).
		setRequestOption(reqOpt...).
		setResult(httpResp).
		execute("GET", "/douyin/feed/")
	if err != nil {
		return nil, nil, err
	}

	resp = httpResp
	rawResponse = ret.rawResponse
	return resp, rawResponse, nil
}

func (s *ApiServiceClient) GetPublishList(context context.Context, req *douyinapi.GetPublishListRequest, reqOpt ...config.RequestOption) (resp *douyinapi.GetPublishListResponse, rawResponse *protocol.Response, err error) {
	httpResp := &douyinapi.GetPublishListResponse{}
	ret, err := s.client.r().
		setContext(context).
		setQueryParams(map[string]interface{}{
			"token":   req.GetToken(),
			"user_id": req.GetUserID(),
		}).
		setPathParams(map[string]string{}).
		setHeaders(map[string]string{}).
		setFormParams(map[string]string{}).
		setFormFileParams(map[string]string{}).
		setBodyParam(req).
		setRequestOption(reqOpt...).
		setResult(httpResp).
		execute("GET", "/douyin/publish/list/")
	if err != nil {
		return nil, nil, err
	}

	resp = httpResp
	rawResponse = ret.rawResponse
	return resp, rawResponse, nil
}

func (s *ApiServiceClient) PublishVideo(context context.Context, req *douyinapi.PublishVideoRequest, reqOpt ...config.RequestOption) (resp *douyinapi.PublishVideoResponse, rawResponse *protocol.Response, err error) {
	httpResp := &douyinapi.PublishVideoResponse{}
	ret, err := s.client.r().
		setContext(context).
		setQueryParams(map[string]interface{}{}).
		setPathParams(map[string]string{}).
		setHeaders(map[string]string{}).
		setFormParams(map[string]string{
			"data":  fmt.Sprint(req.GetData()),
			"token": req.GetToken(),
			"title": req.GetTitle(),
		}).
		setFormFileParams(map[string]string{}).
		setRequestOption(reqOpt...).
		setResult(httpResp).
		execute("POST", "/douyin/publish/action/")
	if err != nil {
		return nil, nil, err
	}

	resp = httpResp
	rawResponse = ret.rawResponse
	return resp, rawResponse, nil
}

func (s *ApiServiceClient) FavoriteAction(context context.Context, req *douyinapi.FavoriteActionRequest, reqOpt ...config.RequestOption) (resp *douyinapi.FavoriteActionResponse, rawResponse *protocol.Response, err error) {
	httpResp := &douyinapi.FavoriteActionResponse{}
	ret, err := s.client.r().
		setContext(context).
		setQueryParams(map[string]interface{}{
			"token":       req.GetToken(),
			"video_id":    req.GetVideoID(),
			"action_type": req.GetActionType(),
		}).
		setPathParams(map[string]string{}).
		setHeaders(map[string]string{}).
		setFormParams(map[string]string{}).
		setFormFileParams(map[string]string{}).
		setBodyParam(req).
		setRequestOption(reqOpt...).
		setResult(httpResp).
		execute("POST", "/douyin/favorite/action/")
	if err != nil {
		return nil, nil, err
	}

	resp = httpResp
	rawResponse = ret.rawResponse
	return resp, rawResponse, nil
}

func (s *ApiServiceClient) GetFavoriteList(context context.Context, req *douyinapi.GetFavoriteListRequest, reqOpt ...config.RequestOption) (resp *douyinapi.GetFavoriteListResponse, rawResponse *protocol.Response, err error) {
	httpResp := &douyinapi.GetFavoriteListResponse{}
	ret, err := s.client.r().
		setContext(context).
		setQueryParams(map[string]interface{}{
			"token":   req.GetToken(),
			"user_id": req.GetUserID(),
		}).
		setPathParams(map[string]string{}).
		setHeaders(map[string]string{}).
		setFormParams(map[string]string{}).
		setFormFileParams(map[string]string{}).
		setBodyParam(req).
		setRequestOption(reqOpt...).
		setResult(httpResp).
		execute("GET", "/douyin/favorite/list/")
	if err != nil {
		return nil, nil, err
	}

	resp = httpResp
	rawResponse = ret.rawResponse
	return resp, rawResponse, nil
}

func (s *ApiServiceClient) CommentAction(context context.Context, req *douyinapi.CommentActionRequest, reqOpt ...config.RequestOption) (resp *douyinapi.CommentActionResponse, rawResponse *protocol.Response, err error) {
	httpResp := &douyinapi.CommentActionResponse{}
	ret, err := s.client.r().
		setContext(context).
		setQueryParams(map[string]interface{}{
			"token":        req.GetToken(),
			"video_id":     req.GetVideoID(),
			"action_type":  req.GetActionType(),
			"comment_text": req.GetCommentText(),
			"comment_id":   req.GetCommentID(),
		}).
		setPathParams(map[string]string{}).
		setHeaders(map[string]string{}).
		setFormParams(map[string]string{}).
		setFormFileParams(map[string]string{}).
		setBodyParam(req).
		setRequestOption(reqOpt...).
		setResult(httpResp).
		execute("POST", "/douyin/comment/action/")
	if err != nil {
		return nil, nil, err
	}

	resp = httpResp
	rawResponse = ret.rawResponse
	return resp, rawResponse, nil
}

func (s *ApiServiceClient) CommentList(context context.Context, req *douyinapi.CommentListRequest, reqOpt ...config.RequestOption) (resp *douyinapi.CommentListResponse, rawResponse *protocol.Response, err error) {
	httpResp := &douyinapi.CommentListResponse{}
	ret, err := s.client.r().
		setContext(context).
		setQueryParams(map[string]interface{}{
			"token":    req.GetToken(),
			"video_id": req.GetVideoID(),
		}).
		setPathParams(map[string]string{}).
		setHeaders(map[string]string{}).
		setFormParams(map[string]string{}).
		setFormFileParams(map[string]string{}).
		setBodyParam(req).
		setRequestOption(reqOpt...).
		setResult(httpResp).
		execute("GET", "/douyin/comment/list/")
	if err != nil {
		return nil, nil, err
	}

	resp = httpResp
	rawResponse = ret.rawResponse
	return resp, rawResponse, nil
}

func (s *ApiServiceClient) RelationAction(context context.Context, req *douyinapi.RelationActionRequest, reqOpt ...config.RequestOption) (resp *douyinapi.RelationActionResponse, rawResponse *protocol.Response, err error) {
	httpResp := &douyinapi.RelationActionResponse{}
	ret, err := s.client.r().
		setContext(context).
		setQueryParams(map[string]interface{}{
			"token":       req.GetToken(),
			"to_user_id":  req.GetToUserID(),
			"action_type": req.GetActionType(),
		}).
		setPathParams(map[string]string{}).
		setHeaders(map[string]string{}).
		setFormParams(map[string]string{}).
		setFormFileParams(map[string]string{}).
		setBodyParam(req).
		setRequestOption(reqOpt...).
		setResult(httpResp).
		execute("POST", "/douyin/relation/action/")
	if err != nil {
		return nil, nil, err
	}

	resp = httpResp
	rawResponse = ret.rawResponse
	return resp, rawResponse, nil
}

func (s *ApiServiceClient) FollowList(context context.Context, req *douyinapi.FollowAndFollowerListRequest, reqOpt ...config.RequestOption) (resp *douyinapi.FollowAndFollowerListResponse, rawResponse *protocol.Response, err error) {
	httpResp := &douyinapi.FollowAndFollowerListResponse{}
	ret, err := s.client.r().
		setContext(context).
		setQueryParams(map[string]interface{}{
			"user_id": req.GetUserID(),
			"token":   req.GetToken(),
		}).
		setPathParams(map[string]string{}).
		setHeaders(map[string]string{}).
		setFormParams(map[string]string{}).
		setFormFileParams(map[string]string{}).
		setBodyParam(req).
		setRequestOption(reqOpt...).
		setResult(httpResp).
		execute("GET", "/douyin/relation/follow/list/")
	if err != nil {
		return nil, nil, err
	}

	resp = httpResp
	rawResponse = ret.rawResponse
	return resp, rawResponse, nil
}

func (s *ApiServiceClient) FollowerList(context context.Context, req *douyinapi.FollowAndFollowerListRequest, reqOpt ...config.RequestOption) (resp *douyinapi.FollowAndFollowerListResponse, rawResponse *protocol.Response, err error) {
	httpResp := &douyinapi.FollowAndFollowerListResponse{}
	ret, err := s.client.r().
		setContext(context).
		setQueryParams(map[string]interface{}{
			"user_id": req.GetUserID(),
			"token":   req.GetToken(),
		}).
		setPathParams(map[string]string{}).
		setHeaders(map[string]string{}).
		setFormParams(map[string]string{}).
		setFormFileParams(map[string]string{}).
		setBodyParam(req).
		setRequestOption(reqOpt...).
		setResult(httpResp).
		execute("GET", "/douyin/relation/follower/list/")
	if err != nil {
		return nil, nil, err
	}

	resp = httpResp
	rawResponse = ret.rawResponse
	return resp, rawResponse, nil
}

func (s *ApiServiceClient) FriendList(context context.Context, req *douyinapi.FriendListRequest, reqOpt ...config.RequestOption) (resp *douyinapi.FriendListResponse, rawResponse *protocol.Response, err error) {
	httpResp := &douyinapi.FriendListResponse{}
	ret, err := s.client.r().
		setContext(context).
		setQueryParams(map[string]interface{}{
			"user_id": req.GetUserID(),
			"token":   req.GetToken(),
		}).
		setPathParams(map[string]string{}).
		setHeaders(map[string]string{}).
		setFormParams(map[string]string{}).
		setFormFileParams(map[string]string{}).
		setBodyParam(req).
		setRequestOption(reqOpt...).
		setResult(httpResp).
		execute("GET", "/douyin/relation/friend/list/")
	if err != nil {
		return nil, nil, err
	}

	resp = httpResp
	rawResponse = ret.rawResponse
	return resp, rawResponse, nil
}

func (s *ApiServiceClient) MessageChat(context context.Context, req *douyinapi.MessageChatRequest, reqOpt ...config.RequestOption) (resp *douyinapi.MessageChatResponse, rawResponse *protocol.Response, err error) {
	httpResp := &douyinapi.MessageChatResponse{}
	ret, err := s.client.r().
		setContext(context).
		setQueryParams(map[string]interface{}{
			"token":        req.GetToken(),
			"to_user_id":   req.GetToUserID(),
			"pre_msg_time": req.GetPreMsgTime(),
		}).
		setPathParams(map[string]string{}).
		setHeaders(map[string]string{}).
		setFormParams(map[string]string{}).
		setFormFileParams(map[string]string{}).
		setBodyParam(req).
		setRequestOption(reqOpt...).
		setResult(httpResp).
		execute("GET", "/douyin/message/chat/")
	if err != nil {
		return nil, nil, err
	}

	resp = httpResp
	rawResponse = ret.rawResponse
	return resp, rawResponse, nil
}

func (s *ApiServiceClient) MessageAction(context context.Context, req *douyinapi.MessageActionRequest, reqOpt ...config.RequestOption) (resp *douyinapi.MessageActionResponse, rawResponse *protocol.Response, err error) {
	httpResp := &douyinapi.MessageActionResponse{}
	ret, err := s.client.r().
		setContext(context).
		setQueryParams(map[string]interface{}{
			"token":       req.GetToken(),
			"to_user_id":  req.GetToUserID(),
			"action_type": req.GetActionType(),
			"content":     req.GetContent(),
		}).
		setPathParams(map[string]string{}).
		setHeaders(map[string]string{}).
		setFormParams(map[string]string{}).
		setFormFileParams(map[string]string{}).
		setBodyParam(req).
		setRequestOption(reqOpt...).
		setResult(httpResp).
		execute("POST", "/douyin/message/action/")
	if err != nil {
		return nil, nil, err
	}

	resp = httpResp
	rawResponse = ret.rawResponse
	return resp, rawResponse, nil
}

var defaultClient, _ = NewApiServiceClient("http://127.0.0.1:8080")

func ConfigDefaultClient(ops ...Option) (err error) {
	defaultClient, err = NewApiServiceClient("http://127.0.0.1:8080", ops...)
	return
}

func RegistUser(context context.Context, req *douyinapi.RegistUserRequest, reqOpt ...config.RequestOption) (resp *douyinapi.RegistUserResponse, rawResponse *protocol.Response, err error) {
	return defaultClient.RegistUser(context, req, reqOpt...)
}

func CheckUser(context context.Context, req *douyinapi.CheckUserRequest, reqOpt ...config.RequestOption) (resp *douyinapi.CheckUserResponse, rawResponse *protocol.Response, err error) {
	return defaultClient.CheckUser(context, req, reqOpt...)
}

func GetUser(context context.Context, req *douyinapi.GetUserRequest, reqOpt ...config.RequestOption) (resp *douyinapi.GetUserResponse, rawResponse *protocol.Response, err error) {
	return defaultClient.GetUser(context, req, reqOpt...)
}

func GetFeed(context context.Context, req *douyinapi.GetFeedRequest, reqOpt ...config.RequestOption) (resp *douyinapi.GetFeedResponse, rawResponse *protocol.Response, err error) {
	return defaultClient.GetFeed(context, req, reqOpt...)
}

func GetPublishList(context context.Context, req *douyinapi.GetPublishListRequest, reqOpt ...config.RequestOption) (resp *douyinapi.GetPublishListResponse, rawResponse *protocol.Response, err error) {
	return defaultClient.GetPublishList(context, req, reqOpt...)
}

func PublishVideo(context context.Context, req *douyinapi.PublishVideoRequest, reqOpt ...config.RequestOption) (resp *douyinapi.PublishVideoResponse, rawResponse *protocol.Response, err error) {
	return defaultClient.PublishVideo(context, req, reqOpt...)
}

func FavoriteAction(context context.Context, req *douyinapi.FavoriteActionRequest, reqOpt ...config.RequestOption) (resp *douyinapi.FavoriteActionResponse, rawResponse *protocol.Response, err error) {
	return defaultClient.FavoriteAction(context, req, reqOpt...)
}

func GetFavoriteList(context context.Context, req *douyinapi.GetFavoriteListRequest, reqOpt ...config.RequestOption) (resp *douyinapi.GetFavoriteListResponse, rawResponse *protocol.Response, err error) {
	return defaultClient.GetFavoriteList(context, req, reqOpt...)
}

func CommentAction(context context.Context, req *douyinapi.CommentActionRequest, reqOpt ...config.RequestOption) (resp *douyinapi.CommentActionResponse, rawResponse *protocol.Response, err error) {
	return defaultClient.CommentAction(context, req, reqOpt...)
}

func CommentList(context context.Context, req *douyinapi.CommentListRequest, reqOpt ...config.RequestOption) (resp *douyinapi.CommentListResponse, rawResponse *protocol.Response, err error) {
	return defaultClient.CommentList(context, req, reqOpt...)
}

func RelationAction(context context.Context, req *douyinapi.RelationActionRequest, reqOpt ...config.RequestOption) (resp *douyinapi.RelationActionResponse, rawResponse *protocol.Response, err error) {
	return defaultClient.RelationAction(context, req, reqOpt...)
}

func FollowList(context context.Context, req *douyinapi.FollowAndFollowerListRequest, reqOpt ...config.RequestOption) (resp *douyinapi.FollowAndFollowerListResponse, rawResponse *protocol.Response, err error) {
	return defaultClient.FollowList(context, req, reqOpt...)
}

func FollowerList(context context.Context, req *douyinapi.FollowAndFollowerListRequest, reqOpt ...config.RequestOption) (resp *douyinapi.FollowAndFollowerListResponse, rawResponse *protocol.Response, err error) {
	return defaultClient.FollowerList(context, req, reqOpt...)
}

func FriendList(context context.Context, req *douyinapi.FriendListRequest, reqOpt ...config.RequestOption) (resp *douyinapi.FriendListResponse, rawResponse *protocol.Response, err error) {
	return defaultClient.FriendList(context, req, reqOpt...)
}

func MessageChat(context context.Context, req *douyinapi.MessageChatRequest, reqOpt ...config.RequestOption) (resp *douyinapi.MessageChatResponse, rawResponse *protocol.Response, err error) {
	return defaultClient.MessageChat(context, req, reqOpt...)
}

func MessageAction(context context.Context, req *douyinapi.MessageActionRequest, reqOpt ...config.RequestOption) (resp *douyinapi.MessageActionResponse, rawResponse *protocol.Response, err error) {
	return defaultClient.MessageAction(context, req, reqOpt...)
}