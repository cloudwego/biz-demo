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

package productpage

import (
	"context"
	"net/http"

	"github.com/cloudwego/biz-demo/bookinfo/kitex_gen/base"
	"github.com/cloudwego/biz-demo/bookinfo/kitex_gen/cwg/bookinfo/details"
	"github.com/cloudwego/biz-demo/bookinfo/kitex_gen/cwg/bookinfo/details/detailsservice"
	"github.com/cloudwego/biz-demo/bookinfo/kitex_gen/cwg/bookinfo/product"
	"github.com/cloudwego/biz-demo/bookinfo/kitex_gen/cwg/bookinfo/reviews"
	"github.com/cloudwego/biz-demo/bookinfo/kitex_gen/cwg/bookinfo/reviews/reviewsservice"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/kitex/pkg/klog"
	"go.opentelemetry.io/otel/baggage"
)

type Handler struct {
	reviewsClient reviewsservice.Client
	detailsClient detailsservice.Client
}

func New(reviewsClient reviewsservice.Client, detailsClient detailsservice.Client) *Handler {
	return &Handler{reviewsClient: reviewsClient, detailsClient: detailsClient}
}

func (h *Handler) GetProduct(ctx context.Context, c *app.RequestContext) {
	productID := c.Param("productID")

	bags := baggage.FromContext(ctx)
	env := bags.Member("env")
	klog.CtxInfof(ctx, "env from baggage: %s", env.String())

	reviewsResp, err := h.reviewsClient.ReviewProduct(ctx, &reviews.ReviewReq{ProductID: productID})
	if err != nil {
		klog.CtxErrorf(ctx, "call reviews error: %s", err.Error())
		c.JSON(http.StatusInternalServerError, &base.BaseResp{
			StatusMessage: "internal error",
			StatusCode:    http.StatusInternalServerError,
			Extra:         nil,
		})
		return
	}

	detailsResp, err := h.detailsClient.GetProduct(ctx, &details.GetProductReq{ID: productID})
	if err != nil {
		klog.CtxErrorf(ctx, "call details error: %s", err.Error())
		return
	}
	p := detailsResp.GetProduct()

	resp := &product.Product{
		ID:          productID,
		Title:       p.GetTitle(),
		Author:      p.GetAuthor(),
		Description: p.GetDescription(),
		Rating:      reviewsResp.GetReview().GetRating(),
	}

	c.JSON(http.StatusOK, resp)
}
