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

package reviews

import (
	"context"
	"os"

	"github.com/cloudwego/biz-demo/bookinfo/kitex_gen/cwg/bookinfo/ratings"
	"github.com/cloudwego/biz-demo/bookinfo/kitex_gen/cwg/bookinfo/ratings/ratingservice"
	"github.com/cloudwego/biz-demo/bookinfo/kitex_gen/cwg/bookinfo/reviews"
	"github.com/cloudwego/biz-demo/bookinfo/pkg/constants"
	"github.com/cloudwego/kitex/pkg/klog"
	"go.opentelemetry.io/otel/baggage"
)

type impl struct {
	ratingsClient ratingservice.Client
}

// New create service impl
func New(ratingsClient ratingservice.Client) reviews.ReviewsService {
	return &impl{ratingsClient: ratingsClient}
}

// ReviewProduct review product
func (i *impl) ReviewProduct(ctx context.Context, req *reviews.ReviewReq) (r *reviews.ReviewResp, err error) {
	klog.CtxDebugf(ctx, "baggage: %s", baggage.FromContext(ctx).String())

	if os.Getenv(constants.EnableRatingsEnvKey) == constants.Disable {
		return &reviews.ReviewResp{
			Review: &reviews.Review{
				Type:   reviews.ReviewType_Local,
				Rating: 0,
			},
		}, err
	}

	ratingResp, err := i.ratingsClient.Ratings(ctx, &ratings.RatingReq{ProductID: req.GetProductID()})
	if err != nil {
		klog.CtxErrorf(ctx, "call ratings error: %s", err.Error())
		return nil, err
	}

	return &reviews.ReviewResp{
		Review: &reviews.Review{
			Type:   reviews.ReviewType_Green,
			Rating: ratingResp.GetRating(),
		},
	}, nil
}
