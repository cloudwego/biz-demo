
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
	"github.com/cloudwego/biz-demo/sample_douyin/pkg/consts"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Username        string `json:"username"`
	Password        string `json:"password"`
	FollowCount     int    `json:"follow_count"`
	FollowerCount   int    `json:"follower_count"`
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
