package service

import (
	"context"
	"mydouyin/cmd/comment/dal/db"
	"mydouyin/cmd/comment/pack"
	"mydouyin/kitex_gen/douyincomment"
)

type GetVideoCommentsService struct {
	ctx context.Context
}

// 
func NewGetVideoCommentsService(ctx context.Context) *GetVideoCommentsService{
	return &GetVideoCommentsService{
		ctx: ctx,
	}
}

// Get the list of comments of a video
func (s *GetVideoCommentsService) GetVideoComments(req *douyincomment.GetVideoCommentsRequest) ([]*douyincomment.Comment, error) {
	comments, err := db.GetVideoComments(s.ctx, req.Video)
	if err != nil {
		return nil, err
	}
	return pack.Comments(comments), nil
}
