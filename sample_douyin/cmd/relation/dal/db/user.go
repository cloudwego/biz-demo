package db

import (
	"mydouyin/pkg/consts"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Username      string `json:"username"`
	Password      string `json:"password"`
	FollowCount   int    `json:"follow_count"`
	FollowerCount int    `json:"follower_count"`
}

func (u *User) TableName() string {
	return consts.UserTableName
}
