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
	"github.com/cloudwego/biz-demo/sample_douyin/cmd/comment/dal/db"
	"github.com/cloudwego/biz-demo/sample_douyin/kitex_gen/douyincomment"
)

// object : change DB pattern to RPC pattern
func Comment(c *db.Comment) *douyincomment.Comment {
	if c == nil {
		return nil
	}
	return &douyincomment.Comment{
		CommentId: int64(c.ID),
		Video: c.Video,
		User: c.User,
		Content: c.Content,
		CreateDate: c.Date,
	}
} 

// list : change DB pattern to RPC pattern
func Comments(c []*db.Comment) []*douyincomment.Comment {
	comments := make([]*douyincomment.Comment, 0)
	for _, v := range c {
		if temp := Comment(v); temp != nil {
			comments = append(comments, temp)
		}
	}
	return comments
}