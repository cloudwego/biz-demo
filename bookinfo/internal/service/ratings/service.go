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

package ratings

import (
	"context"

	"github.com/cloudwego/biz-demo/bookinfo/kitex_gen/cwg/bookinfo/ratings"
)

type impl struct {
}

func New() ratings.RatingService {
	return &impl{}
}

func (i *impl) Ratings(ctx context.Context, req *ratings.RatingReq) (r *ratings.RatingResp, err error) {
	return &ratings.RatingResp{
		Rating: 4,
	}, nil
}
