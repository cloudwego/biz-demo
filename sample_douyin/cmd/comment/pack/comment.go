package pack

import (
	"mydouyin/cmd/comment/dal/db"
	"mydouyin/kitex_gen/douyincomment"
)

// object : change DB pattern to RPC pattern
func Comment(c *db.Comment) *douyincomment.Comment {
	if c == nil {
		return nil
	}
	return &douyincomment.Comment{
		CommentId: int64(c.ID),
		Video: c.Video,
		User: c.User,
		Content: c.Content,
		CreateDate: c.Date,
	}
} 

// list : change DB pattern to RPC pattern
func Comments(c []*db.Comment) []*douyincomment.Comment {
	comments := make([]*douyincomment.Comment, 0)
	for _, v := range c {
		if temp := Comment(v); temp != nil {
			comments = append(comments, temp)
		}
	}
	return comments
}