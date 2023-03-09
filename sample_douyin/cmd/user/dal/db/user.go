package db

import (
	"context"
	"mydouyin/pkg/consts"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Username      string `json:"username"`
	Password      string `json:"password"`
	FollowCount   int    `json:"follow_count"`
	FollowerCount int    `json:"follower_count"`
	FavoriteCount   int  `json:"favorite_count"`
	WorkCount       int  `json:"work_count"`
	TotalFavorited  int  `json:"total_favorited"`
	BackgroundImage string `json:"background_image"`
	Avatar          string `json:"avatar"`
	Signature       string `json:"signature"`

}

func (u *User) TableName() string {
	return consts.UserTableName
}

// CreateUser create user info
func CreateUser(ctx context.Context, users []*User) error {
	return DB.WithContext(ctx).Create(users).Error
}

// QueryUser query list of user info
func QueryUser(ctx context.Context, userName string) ([]*User, error) {
	res := make([]*User, 0)
	if err := DB.WithContext(ctx).Where("username = ?", userName).Find(&res).Error; err != nil {
		return nil, err
	}
	return res, nil
}

// MGetUsers multiple get list of user info
func MGetUSers(ctx context.Context, userIDs []int64) ([]*User, error) {
	res := make([]*User, 0)
	if len(userIDs) == 0 {
		return res, nil
	}

	if err := DB.WithContext(ctx).Where("id in ?", userIDs).Find(&res).Error; err != nil {
		return nil, err
	}
	return res, nil
}
