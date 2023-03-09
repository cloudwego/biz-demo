package main

import (
	"context"
	"mydouyin/cmd/video/pack"
	"mydouyin/cmd/video/service"
	douyinvideo "mydouyin/kitex_gen/douyinvideo"
	"mydouyin/pkg/errno"
)

// VideoServiceImpl implements the last service interface defined in the IDL.
type VideoServiceImpl struct{}

// CreateVideo implements the VideoServiceImpl interface.
func (s *VideoServiceImpl) CreateVideo(ctx context.Context, req *douyinvideo.CreateVideoRequest) (resp *douyinvideo.CreateVideoResponse, err error) {
	// TODO: Your code here...
	resp = new(douyinvideo.CreateVideoResponse)
	if err = req.IsValid(); err != nil {
		resp.BaseResp = pack.BuildBaseResp(errno.ParamErr)
		return resp, nil
	}
	var idList []int64
	idList, err = service.NewCreateVideoService(ctx).CreateVideo(req)
	if err != nil {
		resp.BaseResp = pack.BuildBaseResp(err)
		return resp, nil
	}
	resp.BaseResp = pack.BuildBaseResp(errno.Success)
	resp.VideoIds = idList
	return resp, nil
}

// GetFeed implements the VideoServiceImpl interface.
func (s *VideoServiceImpl) GetFeed(ctx context.Context, req *douyinvideo.GetFeedRequest) (resp *douyinvideo.GetFeedResponse, err error) {
	resp = new(douyinvideo.GetFeedResponse)
	var videos []*douyinvideo.Video
	var nextTime int64
	nextTime, videos, err = service.NewGetFeedService(ctx).GetFeed(req)
	if err != nil {
		resp.BaseResp = pack.BuildBaseResp(err)
		return resp, nil
	}
	resp.NextTime = nextTime
	resp.VideoList = videos
	resp.BaseResp = pack.BuildBaseResp(errno.Success)
	return resp, nil
}

// GetList implements the VideoServiceImpl interface.
func (s *VideoServiceImpl) GetList(ctx context.Context, req *douyinvideo.GetListRequest) (resp *douyinvideo.GetListResponse, err error) {
	resp = new(douyinvideo.GetListResponse)
	var videos []*douyinvideo.Video
	videos, err = service.NewGetListService(ctx).GetList(req)
	if err != nil {
		resp.BaseResp = pack.BuildBaseResp(err)
		return resp, nil
	}
	resp.VideoList = videos
	resp.BaseResp = pack.BuildBaseResp(errno.Success)
	return resp, nil
}

// MGetVideoUser implements the VideoServiceImpl interface.
func (s *VideoServiceImpl) MGetVideoUser(ctx context.Context, req *douyinvideo.MGetVideoRequest) (resp *douyinvideo.MGetVideoResponse, err error) {
	resp = new(douyinvideo.MGetVideoResponse)
	var videos []*douyinvideo.Video
	videos, err = service.NewMGetVideoService(ctx).MGetVideo(req)
	if err != nil {
		resp.BaseResp = pack.BuildBaseResp(err)
		return resp, nil
	}
	resp.Videos = videos
	resp.BaseResp = pack.BuildBaseResp(errno.Success)
	return resp, nil
}

// DeleteVideo implements the VideoServiceImpl interface.
func (s *VideoServiceImpl) DeleteVideo(ctx context.Context, req *douyinvideo.DeleteVideoRequest) (resp *douyinvideo.DeleteVideoResponse, err error) {
	// TODO: Your code here...
	return
}

// GetTimeVideos implements the VideoServiceImpl interface.
func (s *VideoServiceImpl) GetTimeVideos(ctx context.Context, req *douyinvideo.GetTimeVideosRequest) (resp *douyinvideo.GetTimeVideosResponse, err error) {
	resp = new(douyinvideo.GetTimeVideosResponse)
	var videos []*douyinvideo.Video
	videos, err = service.NewGetTimeVideosService(ctx).GetTimeVideos(req)
	if err != nil {
		resp.BaseResp = pack.BuildBaseResp(err)
		return resp, nil
	}
	resp.VideoList = videos
	resp.BaseResp = pack.BuildBaseResp(errno.Success)
	return resp, nil
}
