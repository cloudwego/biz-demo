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

	"github.com/cloudwego/biz-demo/sample_douyin/pkg/consts"

	"gorm.io/gorm"
)

type Comment struct {
	gorm.Model
	Video   int64  `json:"video"`
	User    int64  `json:"user"`
	Content string `json:"content"`
	Date    string `json:"date"`
}

func (v *Comment) TableName() string {
	return consts.CommentTableName
}

// create a comment
func CreateComment(ctx context.Context, comment *Comment) (int64, error) {
	result := DB.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		// 在事务中执行一些 db 操作（从这里开始，您应该使用 'tx' 而不是 'db'）
		if err := tx.Create(comment).Error; err != nil {
			// 返回任何错误都会回滚事务
			return err
		}
		// update the comment number of the video
		if err := tx.Model(&Video{}).Where("id = ?", comment.Video).Update("comment_count", gorm.Expr("comment_count + ?", 1)).Error; err != nil {
			return err
		}
		// 返回 nil 提交事务
		return nil
	})
	return int64(comment.ID), result
}

// delete a comment
func DeleteComment(ctx context.Context, comment_id int64) error {
	return DB.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		var comment Comment
		if err := tx.First(&comment, comment_id).Error; err != nil {
			return err
		}
		// 在事务中执行一些 db 操作（从这里开始，您应该使用 'tx' 而不是 'db'）
		if err := tx.Delete(&Comment{}, comment_id).Error; err != nil {
			return err
		}
		if err := tx.Model(&Video{}).Where("id = ?", comment.Video).Update("comment_count", gorm.Expr("comment_count - ?", 1)).Error; err != nil {
			return err
		}
		// 返回 nil 提交事务
		return nil
	})
}

// get a comment list of a video
func GetVideoComments(ctx context.Context, video_id int64) ([]*Comment, error) {
	res := make([]*Comment, 0)
	// it will search the context that didn't be deleted
	if err := DB.WithContext(ctx).Where("video = ?", video_id).Order("id desc").Find(&res).Error; err != nil {
		return nil, err
	}
	return res, nil
}
