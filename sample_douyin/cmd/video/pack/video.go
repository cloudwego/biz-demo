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
	"github.com/cloudwego/biz-demo/sample_douyin/cmd/video/dal/db"
	"github.com/cloudwego/biz-demo/sample_douyin/kitex_gen/douyinvideo"
)

// Video pack video info
func Video(v *db.Video) *douyinvideo.Video {
	if v == nil {
		return nil
	}
	return &douyinvideo.Video{
		VideoId:       int64(v.ID),
		Author:        int64(v.Author),
		PlayUrl:       v.PlayUrl,
		CoverUrl:      v.CoverUrl,
		Title:         v.Title,
		FavoriteCount: int64(v.FavoriteCount),
		CommentCount:  int64(v.CommentCount),
		UploadTime:    v.CreatedAt.Format("20060102.150405"),
	}
}

func Videos(vs []*db.Video) []*douyinvideo.Video {
	videos := make([]*douyinvideo.Video, 0)
	for _, v := range vs {
		if temp := Video(v); temp != nil {
			videos = append(videos, temp)
		}
	}
	return videos
}
