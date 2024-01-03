// Copyright 2022 CloudWeGo Authors
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
	"github.com/cloudwego/biz-demo/easy_note/cmd/user/dal/db"
	"github.com/cloudwego/biz-demo/easy_note/kitex_gen/demouser"
)

// User pack user info
func User(u *db.User) *demouser.User {
	if u == nil {
		return nil
	}

	return &demouser.User{UserId: int64(u.ID), Username: u.Username, Avatar: "test"}
}

// Users pack list of user info
func Users(us []*db.User) []*demouser.User {
	users := make([]*demouser.User, 0)
	for _, u := range us {
		if temp := User(u); temp != nil {
			users = append(users, temp)
		}
	}
	return users
}
