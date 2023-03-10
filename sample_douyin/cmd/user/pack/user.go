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

package pack

import (
	"github.com/cloudwego/biz-demo/sample_douyin/cmd/user/dal/db"
	"github.com/cloudwego/biz-demo/sample_douyin/kitex_gen/douyinuser"
)

// User pack user info
func User(u *db.User) *douyinuser.User {
	if u == nil {
		return nil
	}

	return &douyinuser.User{
		UserId:          int64(u.ID),
		Username:        u.Username,
		FollowCount:     int64(u.FollowCount),
		FollowerCount:   int64(u.FollowerCount),
		TotalFavorited:  int64(u.TotalFavorited),
		FavoriteCount:   int64(u.FavoriteCount),
		WorkCount:       int64(u.WorkCount),
		Avatar:          u.Avatar,
		Signature:       u.Signature,
		BackgroundImage: u.BackgroundImage,
	}
}

func Users(us []*db.User) []*douyinuser.User {
	users := make([]*douyinuser.User, 0)
	for _, u := range us {
		if temp := User(u); temp != nil {
			users = append(users, temp)
		}
	}
	return users
}
