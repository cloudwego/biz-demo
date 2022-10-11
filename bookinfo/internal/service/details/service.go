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

package details

import (
	"context"

	"github.com/cloudwego/biz-demo/bookinfo/kitex_gen/cwg/bookinfo/details"
	"github.com/cloudwego/kitex/pkg/klog"
	"go.opentelemetry.io/otel/baggage"
)

type impl struct{}

// New create service impl instance
func New() details.DetailsService {
	return &impl{}
}

// GetProduct get product detail info
func (i *impl) GetProduct(ctx context.Context, req *details.GetProductReq) (r *details.GetProductResp, err error) {
	klog.CtxInfof(ctx, "get product details %s", req.ID)
	klog.CtxDebugf(ctx, "baggage: %s", baggage.FromContext(ctx).String())

	return &details.GetProductResp{
		Product: &details.Product{
			ID:          req.GetID(),
			Title:       "《Also sprach Zarathustra》",
			Author:      "Friedrich Nietzsche",
			Description: `Thus Spoke Zarathustra: A Book for All and None, also translated as Thus Spake Zarathustra, is a work of philosophical fiction written by German philosopher Friedrich Nietzsche between 1883 and 1885.`,
		},
	}, nil
}
