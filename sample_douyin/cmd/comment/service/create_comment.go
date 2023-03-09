package service

import (
	"context"
	"mydouyin/cmd/comment/dal/db"
	"mydouyin/kitex_gen/douyincomment"
)

type CreateCommentService struct {
	ctx context.Context
}

func NewCreateCommentService(ctx context.Context) *CreateCommentService {
	return &CreateCommentService{
		ctx: ctx,
	}
}

// create user
func (c *CreateCommentService) CreateComment(req *douyincomment.CreateCommentRequest) (int64, error) {
	return db.CreateComment(c.ctx, &db.Comment{
		Video: req.Video,
		User: req.User,
		Content: req.Content,
		Date: req.CreateDate,
	})
}
