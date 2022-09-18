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

type impl struct {
}

func New() details.DetailsService {
	return &impl{}
}

func (i *impl) GetProduct(ctx context.Context, req *details.GetProductReq) (r *details.GetProductResp, err error) {
	klog.CtxInfof(ctx, "get product details %s", req.ID)
	bags := baggage.FromContext(ctx)
	env := bags.Member("env")
	klog.CtxInfof(ctx, "env from baggage: %s", env.String())

	return &details.GetProductResp{
		Product: &details.Product{
			ID:          req.GetID(),
			Title:       "《查拉图斯特拉如是说》",
			Author:      "尼采",
			Description: "《查拉图斯特拉如是说》是哲学家、思想家弗里德里希·威廉·尼采创作的散文诗体哲学著作。\n全书共分四个部分。第一部分包括1篇至关重要的序言和2篇演讲，演讲地点主要集中在一个叫做“斑牛镇”的城镇里。第二部分同样由22篇演讲构成，演讲地点主要在幸福岛及其周围。第三部分的16篇演讲是在查拉图斯特拉渡海返回他的山洞的途中所作。第四部分包含20篇演讲，在查拉图斯特拉的洞府及其周围完成。 [1] \n综观全书，尼采要否定的，是以信仰和服从为准则的旧价值体系，他要肯定的，是以生命和人的意志为准则的新价值体系，前者的立足点是上帝，后者的立足点是大地。尼采将即要到来的时代比喻为人类开始觉醒的时代。人类精神仿佛处在由忍辱负重的骆驼向雄武强健的狮子变形的过程之中，它预示着超越人类的赤子新生命的到来。 [2] ",
		},
	}, nil
}
