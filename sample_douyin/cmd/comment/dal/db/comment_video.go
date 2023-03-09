package db

import (
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
