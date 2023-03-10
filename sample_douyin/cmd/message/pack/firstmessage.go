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
	"github.com/cloudwego/biz-demo/sample_douyin/cmd/message/dal/db"
	"github.com/cloudwego/biz-demo/sample_douyin/kitex_gen/message"
)

// User pack user info
func FirstMessage(u *db.FirstMessage) *message.FirstMessage {
	if u == nil {
		return nil
	}

	return &message.FirstMessage{
		Message:  u.Message,
		MsgType:  int64(u.MsgType),
		FriendId: int64(u.FirendID),
	}
}

func FirstMessages(us []*db.FirstMessage) []*message.FirstMessage {
	firstmessages := make([]*message.FirstMessage, 0)
	for _, m := range us {
		if temp := FirstMessage(m); temp != nil {
			firstmessages = append(firstmessages, temp)
		}
	}
	return firstmessages
}
