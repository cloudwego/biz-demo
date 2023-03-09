package db

import (
	"context"
	"mydouyin/pkg/consts"

	"gorm.io/gorm"
)

type Video struct {
	gorm.Model
	Author        int64  `json:"author"`
	PlayUrl       string `json:"play_url"`
	CoverUrl      string `json:"cover_url"`
	Title         string `json:"title"`
	FavoriteCount int    `json:"favorite_count"`
	CommentCount  int    `json:"comment_count"`
}

func (v *Video) TableName() string {
	return consts.VideoTableName
}

// CreateVideo create video info
func CreateVideo(ctx context.Context, videos []*Video) ([]int64, error) {
	idList := make([]int64, 0)
	err := DB.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		if err := tx.WithContext(ctx).Create(videos).Error; err != nil{
			return err
		}
		for _, i := range videos {
			idList = append(idList, int64(i.ID))
			if err := tx.WithContext(ctx).Model(&User{}).Where("id = ?", i.Author).Update("work_count", gorm.Expr("work_count + ?", 1)).Error; err != nil {
				return err
			}
		}
		// 返回 nil 提交事务
		return nil	
	})
	return idList, err
}

// MGetVideos multiple get list of video info
func MGetVideos(ctx context.Context, videoIDs []int64) ([]*Video, error) {
	res := make([]*Video, 0)
	if len(videoIDs) == 0 {
		return res, nil
	}
	if err := DB.WithContext(ctx).Where("id in ?", videoIDs).Find(&res).Error; err != nil {
		return nil, err
	}
	return res, nil
}

// GetFeed multiple get list of video info
func GetFeed(ctx context.Context, latest_time string) ([]*Video, error) {
	res := make([]*Video, 0)
	if err := DB.WithContext(ctx).Where("created_at < ?", latest_time).Limit(30).Order("id desc").Find(&res).Error; err != nil {
		return nil, err
	}
	return res, nil
}

// GetVideosFromTime
func GetVideosFromTime(ctx context.Context, start, end string) ([]*Video, error) {
	res := make([]*Video, 0)
	if err := DB.WithContext(ctx).Where("created_at > ? AND created_at < ?", start, end).Limit(30).Order("id desc").Find(&res).Error; err != nil {
		return nil, err
	}
	return res, nil
}

// MGetVideos multiple get list of video info
func MGetVideosbyAuthor(ctx context.Context, authorID int64) ([]*Video, error) {
	res := make([]*Video, 0)
	if err := DB.WithContext(ctx).Where("author = ?", authorID).Find(&res).Error; err != nil {
		return nil, err
	}
	return res, nil
}

// Delete the video
func DeleteVideo(ctx context.Context, video_id int64) error {
	if err := DB.WithContext(ctx).Delete(&Video{}, video_id).Error; err != nil {
		return nil
	}
	return nil
}
