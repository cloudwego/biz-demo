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
	"github.com/cloudwego/biz-demo/sample_douyin/cmd/favorite/dal/db"
	"github.com/cloudwego/biz-demo/sample_douyin/kitex_gen/douyinfavorite"
)

// Favorite pack video info
func Favorite(f *douyinfavorite.Favorite) *db.Favorite {
	if f == nil {
		return nil
	}
	return &db.Favorite{
		UserId:  int64(f.UserId),
		VideoId: int64(f.VideoId),
	}
}

func Favorites(fs []*douyinfavorite.Favorite) []*db.Favorite {
	favorites := make([]*db.Favorite, 0)
	for _, f := range fs {
		if temp := Favorite(f); temp != nil {
			favorites = append(favorites, temp)
		}
	}
	return favorites
}

func FavoriteToVideoids(favorites []*db.Favorite) []int64 {
	vids := make([]int64, 0)
	for i := 0; i < len(favorites); i++ {
		if temp := int64(favorites[i].VideoId); temp != 0 {
			vids = append(vids, temp)
		}
	}
	return vids
}
