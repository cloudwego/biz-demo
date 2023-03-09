package pack

import (
	"mydouyin/cmd/video/dal/db"
	"mydouyin/kitex_gen/douyinvideo"
)

// Video pack video info
func Video(v *db.Video) *douyinvideo.Video {
	if v == nil {
		return nil
	}
	return &douyinvideo.Video{
		VideoId:       int64(v.ID),
		Author:        int64(v.Author),
		PlayUrl:       v.PlayUrl,
		CoverUrl:      v.CoverUrl,
		Title:         v.Title,
		FavoriteCount: int64(v.FavoriteCount),
		CommentCount:  int64(v.CommentCount),
		UploadTime:    v.CreatedAt.Format("20060102.150405"),
	}
}
func Videos(vs []*db.Video) []*douyinvideo.Video {
	videos := make([]*douyinvideo.Video, 0)
	for _, v := range vs {
		if temp := Video(v); temp != nil {
			videos = append(videos, temp)
		}
	}
	return videos
}
