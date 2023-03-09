package service

import (
	"context"
	"mydouyin/cmd/api/biz/apimodel"
	"mydouyin/cmd/api/biz/rpc"
	videohandel "mydouyin/cmd/api/biz/videoHandel"
	"mydouyin/kitex_gen/douyinuser"
	"mydouyin/kitex_gen/douyinvideo"
	"mydouyin/pkg/errno"
	"strconv"
)

type PublishService struct {
	ctx context.Context
}

func NewPublishService(ctx context.Context) *PublishService {
	return &PublishService{
		ctx: ctx,
	}
}

func (s *PublishService) PublishVideo(req apimodel.PublishVideoRequest, user *apimodel.User) (*apimodel.PublishVideoResponse, error) {
	resp := new(apimodel.PublishVideoResponse)
	//err := videohandel.VH.UpLoadVideoV0(req.Data, user.UserID, req.Title)
    videoName, err := videohandel.VH.UpLoadVideo(req.Data)
	if err != nil {
		return resp, err
	}

	go videohandel.VH.CommitCommand(videoName, user.UserID, req.Title)

	return resp, err
}

func (s *PublishService) GetPublishList(req apimodel.GetPublishListRequest, user *apimodel.User) (*apimodel.GetPublishListResponse, error) {
	resp := new(apimodel.GetPublishListResponse)
	userId, err := strconv.Atoi(req.UserId)
	if err != nil {
		return resp, err
	}
	rpc_resp, err := rpc.GetList(s.ctx, &douyinvideo.GetListRequest{
		UserId: int64(userId),
	})
	if err != nil {
		return nil, err
	}
	if rpc_resp.BaseResp.StatusCode != 0 {
		return nil, errno.NewErrNo(rpc_resp.BaseResp.StatusCode, rpc_resp.BaseResp.StatusMessage)
	}
	resp.VideoList = make([]apimodel.Video, 0, 50)
	for _, rpc_video := range rpc_resp.VideoList {
		r, err := rpc.MGetUser(s.ctx, &douyinuser.MGetUserRequest{UserIds: []int64{rpc_video.Author}})
		if err != nil || r.BaseResp.StatusCode != 0 || len(r.Users) < 1 {
			continue
		}
		author := apimodel.PackUser(r.Users[0])
		video := apimodel.PackVideo(rpc_video)
		video.Author = *author
		resp.VideoList = append(resp.VideoList, *video)
	}
	return resp, nil
}
