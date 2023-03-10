// Copyright 2023 CloudWeGo Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//

package db

import (
	"context"
	"github.com/cloudwego/biz-demo/sample_douyin/cmd/video/dal/db"
	"github.com/cloudwego/biz-demo/sample_douyin/pkg/consts"
	"github.com/cloudwego/biz-demo/sample_douyin/pkg/errno"
	"time"

	"gorm.io/gorm"
)

type Favorite struct {
	ID        uint `gorm:"primarykey"`
	CreatedAt time.Time
	UpdatedAt time.Time
	UserId    int64 `json:"user_id"`
	VideoId   int64 `json:"video_id"`
}

func (f *Favorite) TableName() string {
	return consts.FavoriteTableName
}

// CreateVideo create video info
func CreateFavorite(ctx context.Context, favorites []*Favorite) error {
	err := DB.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		if err := tx.WithContext(ctx).Create(favorites).Error; err != nil {
			return err
		}
		for _, f := range favorites {
			if err := tx.WithContext(ctx).Model(&Video{}).Where("id = ?", f.VideoId).Update("favorite_count", gorm.Expr("favorite_count + ?", 1)).Error; err != nil {
				return err
			}
			if err := tx.WithContext(ctx).Model(&User{}).Where("id = ?", f.UserId).Update("favorite_count", gorm.Expr("favorite_count + ?", 1)).Error; err != nil {
				return err
			}
			videos := make([]*db.Video, 0)
			if err := tx.WithContext(ctx).Model(&Video{}).Where("id = ?", f.VideoId).Find(&videos).Error; err != nil{
				return err
			}
			for _,v := range videos{
				if err := tx.WithContext(ctx).Model(&User{}).Where("id = ?", v.Author).Update("total_favorited", gorm.Expr("total_favorited + ?", 1)).Error; err != nil {
					return err
				}
			}
		}
		// 返回 nil 提交事务
		return nil
	})
	return err
}

func CancleFavorite(ctx context.Context, favorites []*Favorite) error {
	err := DB.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		for _, f := range favorites {
			var favorite Favorite
			if err := tx.Where("user_id = ? and video_id = ?", f.UserId, f.VideoId).Delete(&favorite).Error; err != nil {
				return err
			}
		}
		for _, f := range favorites {
			if err := tx.WithContext(ctx).Model(&Video{}).Where("id = ?", f.VideoId).Update("favorite_count", gorm.Expr("favorite_count - ?", 1)).Error; err != nil {
				return err
			}
			if err := tx.WithContext(ctx).Model(&User{}).Where("id = ?", f.UserId).Update("favorite_count", gorm.Expr("favorite_count - ?", 1)).Error; err != nil {
				return err
			}
			videos := make([]*db.Video, 0)
			if err := tx.WithContext(ctx).Model(&Video{}).Where("id = ?", f.VideoId).Find(&videos).Error; err != nil{
				return err
			}
			for _,v := range videos{
				if err := tx.WithContext(ctx).Model(&User{}).Where("id = ?", v.Author).Update("total_favorited", gorm.Expr("total_favorited - ?", 1)).Error; err != nil {
					return err
				}
			}
		}
		// 返回 nil 提交事务
		return nil
	})
	return err
}

func QueryFavoriteById(ctx context.Context, favorites []*Favorite) ([]bool, error) {
	res := make([]bool, 0)
	for _, favorite := range favorites {
		find := make([]*Favorite, 0)
		if err := DB.WithContext(ctx).Where("user_id = ? and video_id = ?", favorite.UserId, favorite.VideoId).Find(&find).Error; err != nil {
			return res, err
		}
		if len(find) > 0 {
			res = append(res, true)
		} else {
			res = append(res, false)
		}
	}
	if len(res) != len(favorites) {
		return res, errno.NewErrNo(0000000, "something wrong")
	}
	return res, nil
}

func GetFavoriteList(ctx context.Context, userID int64) ([]*Favorite, error) {
	res := make([]*Favorite, 0)
	if err := DB.WithContext(ctx).Where("user_id = ?", userID).Find(&res).Error; err != nil {
		return nil, err
	}
	return res, nil
}
