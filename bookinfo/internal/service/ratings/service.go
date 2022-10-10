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
	"os"
	"strconv"

	"github.com/cloudwego/biz-demo/bookinfo/kitex_gen/cwg/bookinfo/ratings"
	"github.com/cloudwego/biz-demo/bookinfo/pkg/constants"
	"github.com/cloudwego/kitex/pkg/klog"
	"go.opentelemetry.io/otel/baggage"
)

type impl struct{}

// New create service impl
func New() ratings.RatingService {
	return &impl{}
}

// Ratings get ratings
func (i *impl) Ratings(ctx context.Context, req *ratings.RatingReq) (r *ratings.RatingResp, err error) {
	klog.CtxInfof(ctx, "get product details %s", req.ProductID)
	klog.CtxDebugf(ctx, "baggage: %s", baggage.FromContext(ctx).String())

	ratingValue, err := strconv.ParseInt(os.Getenv(constants.RatingsValueEnvKey), 10, 8)
	if err != nil {
		ratingValue = 1
	}
	return &ratings.RatingResp{
		Rating: int8(ratingValue),
	}, nil
}
