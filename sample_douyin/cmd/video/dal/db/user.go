package db

import (
	"mydouyin/pkg/consts"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Username        string `json:"username"`
	Password        string `json:"password"`
	FollowCount     int    `json:"follow_count"`
	FollowerCount   int    `json:"follower_count"`
	FavoriteCount   int    `json:"favorite_count"`
	WorkCount       int    `json:"work_count"`
	TotalFavorited  int    `json:"total_favorited"`
	BackgroundImage string `json:"background_image"`
	Avatar          string `json:"avatar"`
	Signature       string `json:"signature"`
}

func (u *User) TableName() string {
	return consts.UserTableName
}
