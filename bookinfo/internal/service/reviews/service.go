package reviews

import (
	"context"
	"os"

	"github.com/cloudwego/biz-demo/bookinfo/kitex_gen/cwg/bookinfo/ratings"
	"github.com/cloudwego/biz-demo/bookinfo/kitex_gen/cwg/bookinfo/ratings/ratingservice"
	"github.com/cloudwego/biz-demo/bookinfo/kitex_gen/cwg/bookinfo/reviews"
	"github.com/cloudwego/biz-demo/bookinfo/pkg/constants"
)

type impl struct {
	ratingsClient ratingservice.Client
}

func New(ratingsClient ratingservice.Client) reviews.ReviewsService {
	return &impl{ratingsClient: ratingsClient}
}

func (i *impl) ReviewProduct(ctx context.Context, req *reviews.ReviewReq) (r *reviews.ReviewResp, err error) {
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
		return nil, err
	}

	return &reviews.ReviewResp{
		Review: &reviews.Review{
			Type:   reviews.ReviewType_Green,
			Rating: ratingResp.GetRating(),
		},
	}, nil
}
