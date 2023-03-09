package service

import (
	"context"
	"mydouyin/cmd/comment/dal/db"
	"mydouyin/kitex_gen/douyincomment"
)

type DeleteCommentService struct {
	ctx context.Context
}

// get service
func NewDeleteCommentService(ctx context.Context) *DeleteCommentService {
	return &DeleteCommentService{
		ctx: ctx,
	}
}

// Delete the comment service
func (d *DeleteCommentService) DeleteComment(req *douyincomment.DeleteCommentRequest) error {
	err := db.DeleteComment(d.ctx, req.CommentId)
	return err
}
