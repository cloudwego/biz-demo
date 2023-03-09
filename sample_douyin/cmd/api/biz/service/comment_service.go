package service

import (
	"context"
	"mydouyin/cmd/api/biz/apimodel"
	"mydouyin/cmd/api/biz/rpc"
	"mydouyin/kitex_gen/douyincomment"
	"mydouyin/kitex_gen/douyinuser"
	"mydouyin/pkg/errno"
	"strconv"
	"strings"
	"time"
)

type CommentService struct {
	ctx context.Context
}

func NewCommentService(ctx context.Context) *CommentService {
	return &CommentService{
		ctx: ctx,
	}
}

func (s *CommentService) CommentAction(req apimodel.CommentActionRequest, user *apimodel.User) (*apimodel.CommentActionResponse, error) {
	resp := new(apimodel.CommentActionResponse)
	actionType, err := strconv.Atoi(req.ActionType)
	if err != nil {
		return resp, err
	}
	switch actionType {
	case 1:
		// create the date
		getMonth := time.Now().Format("01")
		getDay := time.Now().Format("02")
		var build strings.Builder
		build.WriteString(getMonth)
		build.WriteString("-")
		build.WriteString(getDay)
		date := build.String()
		// create the VideoID
		videoID, err := strconv.ParseInt(req.VideoId, 10, 64)
		if err != nil {
			return resp, err
		}

		rpc_resp, err := rpc.CreateComment(s.ctx, &douyincomment.CreateCommentRequest{
			Video:      videoID,
			User:       user.UserID,
			Content:    req.CommentText,
			CreateDate: date,
		})
		if err != nil {
			return resp, err
		}

		if rpc_resp.BaseResp.StatusCode != 0 {
			return resp, errno.NewErrNo(rpc_resp.BaseResp.StatusCode, rpc_resp.BaseResp.StatusMessage)
		}

		resp.Comment = apimodel.Comment{
			CommentID:  rpc_resp.CommentId,
			Commentor:  *user,
			Content:    req.CommentText,
			CreateDate: date,
		}
	case 2:
		// delete the date
		commentID, err := strconv.ParseInt(req.CommentId, 10, 64)
		if err != nil {
			return resp, err
		}
		rpc_resp, err := rpc.DeleteComment(s.ctx, &douyincomment.DeleteCommentRequest{
			CommentId: commentID,
		})
		if err != nil {
			return resp, err
		}
		if rpc_resp.BaseResp.StatusCode != 0 {
			return resp, errno.NewErrNo(rpc_resp.BaseResp.StatusCode, rpc_resp.BaseResp.StatusMessage)
		}
	default:
		return nil, errno.ParamErr
	}
	return resp, nil
}

func (s *CommentService) CommentList(req apimodel.CommentListRequest) (*apimodel.CommentListResponse, error) {
	resp := new(apimodel.CommentListResponse)
	// get the VideoID
	videoID, err := strconv.ParseInt(req.VideoId, 10, 64)
	if err != nil {
		return resp, err
	}

	rpc_resp, err := rpc.GetVideoComments(s.ctx, &douyincomment.GetVideoCommentsRequest{
		Video: videoID,
	})
	if err != nil {
		return resp, err
	}
	if rpc_resp.BaseResp.StatusCode != 0 {
		return resp, errno.NewErrNo(rpc_resp.BaseResp.StatusCode, rpc_resp.BaseResp.StatusMessage)
	}
	resp.CommentList = make([]apimodel.Comment, 0, 50)
	for _, rpc_comment := range rpc_resp.Comments {
		r, err := rpc.MGetUser(s.ctx, &douyinuser.MGetUserRequest{UserIds: []int64{rpc_comment.User}})
		if err != nil || r.BaseResp.StatusCode != 0 || len(r.Users) < 1 {
			continue
		}
		user := apimodel.PackUser(r.Users[0])
		comment := apimodel.PackComment(rpc_comment)
		comment.Commentor = *user
		resp.CommentList = append(resp.CommentList, *comment)
	}
	return resp, nil
}
