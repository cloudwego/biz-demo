// Copyright 2024 CloudWeGo Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package cart

import (
	"context"
	cart "github.com/cloudwego/biz-demo/gomall/app/cart/kitex_gen/cart"
	"github.com/cloudwego/kitex/client/callopt"
	"github.com/cloudwego/kitex/pkg/klog"
)

func AddItem(ctx context.Context, req *cart.AddItemRequest, callOptions ...callopt.Option) (resp *cart.Empty, err error) {
	resp, err = defaultClient.AddItem(ctx, req, callOptions...)
	if err != nil {
		klog.CtxErrorf(ctx, "AddItem call failed,err =%+v", err)
		return nil, err
	}
	return resp, nil
}

func GetCart(ctx context.Context, req *cart.GetCartRequest, callOptions ...callopt.Option) (resp *cart.Cart, err error) {
	resp, err = defaultClient.GetCart(ctx, req, callOptions...)
	if err != nil {
		klog.CtxErrorf(ctx, "GetCart call failed,err =%+v", err)
		return nil, err
	}
	return resp, nil
}

func EmptyCart(ctx context.Context, req *cart.EmptyCartRequest, callOptions ...callopt.Option) (resp *cart.Empty, err error) {
	resp, err = defaultClient.EmptyCart(ctx, req, callOptions...)
	if err != nil {
		klog.CtxErrorf(ctx, "EmptyCart call failed,err =%+v", err)
		return nil, err
	}
	return resp, nil
}
