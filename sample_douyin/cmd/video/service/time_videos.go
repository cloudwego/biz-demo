package service

import (
	"context"
	"mydouyin/cmd/video/dal/db"
	"mydouyin/cmd/video/pack"
	"mydouyin/kitex_gen/douyinvideo"
)

type GetTimeVideosService struct {
	ctx context.Context
}

func NewGetTimeVideosService(ctx context.Context) *GetTimeVideosService {
	return &GetTimeVideosService{ctx: ctx}
}

func (s *GetTimeVideosService) GetTimeVideos(req *douyinvideo.GetTimeVideosRequest) ([]*douyinvideo.Video, error) {
	videos, err := db.GetVideosFromTime(s.ctx, req.Start, req.End)
	if err != nil {
		return nil, err
	}
	return pack.Videos(videos), nil
}
