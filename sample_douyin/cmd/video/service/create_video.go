package service

import (
	"context"
	"mydouyin/cmd/video/dal/db"
	"mydouyin/kitex_gen/douyinvideo"
)

type CreateVideoService struct {
	ctx context.Context
}

// NewCreateVideoService new CreateVideoService
func NewCreateVideoService(ctx context.Context) *CreateVideoService {
	return &CreateVideoService{ctx: ctx}
}

// CreateVideo create video info.
func (s *CreateVideoService) CreateVideo(req *douyinvideo.CreateVideoRequest) ([]int64,error) {
	return db.CreateVideo(s.ctx, []*db.Video{{
		Author: req.Author,
		PlayUrl: req.PlayUrl,
		CoverUrl: req.CoverUrl,
		Title: req.Title,
	}})
}
