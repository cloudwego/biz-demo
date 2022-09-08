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
