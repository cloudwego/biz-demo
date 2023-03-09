package service

import (
	"context"
	"mydouyin/cmd/video/dal/db"
	"mydouyin/kitex_gen/douyinvideo"
	"mydouyin/cmd/video/pack"
)

type MGetVideoService struct {
	ctx context.Context
}

// MGetVideoService new MGetVideoService
func NewMGetVideoService(ctx context.Context) *MGetVideoService {
	return &MGetVideoService{ctx: ctx}
}

// MGetVideo.
func (s *MGetVideoService) MGetVideo(req *douyinvideo.MGetVideoRequest) ([]*douyinvideo.Video, error) {
	videos, err := db.MGetVideos(s.ctx, req.VideoIds)
	if err != nil {
		return nil, err
	} 
	return pack.Videos(videos) ,nil
}
