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

package main

import (
	"context"
	item "github.com/cloudwego/biz-demo/book-shop/kitex_gen/cwg/bookshop/item"
)

// ItemServiceImpl implements the last service interface defined in the IDL.
type ItemServiceImpl struct{}

// Add implements the ItemServiceImpl interface.
func (s *ItemServiceImpl) Add(ctx context.Context, req *item.AddReq) (resp *item.AddResp, err error) {
	// TODO: Your code here...
	return
}

// Edit implements the ItemServiceImpl interface.
func (s *ItemServiceImpl) Edit(ctx context.Context, req *item.EditReq) (resp *item.EditResp, err error) {
	// TODO: Your code here...
	return
}

// Delete implements the ItemServiceImpl interface.
func (s *ItemServiceImpl) Delete(ctx context.Context, req *item.DeleteReq) (resp *item.DeleteResp, err error) {
	// TODO: Your code here...
	return
}

// Online implements the ItemServiceImpl interface.
func (s *ItemServiceImpl) Online(ctx context.Context, req *item.OnlineReq) (resp *item.OnlineResp, err error) {
	// TODO: Your code here...
	return
}

// Offline implements the ItemServiceImpl interface.
func (s *ItemServiceImpl) Offline(ctx context.Context, req *item.OfflineReq) (resp *item.OfflineResp, err error) {
	// TODO: Your code here...
	return
}

// Get implements the ItemServiceImpl interface.
func (s *ItemServiceImpl) Get(ctx context.Context, req *item.GetReq) (resp *item.GetResp, err error) {
	// TODO: Your code here...
	return
}

// Search implements the ItemServiceImpl interface.
func (s *ItemServiceImpl) Search(ctx context.Context, req *item.SearchReq) (resp *item.SearchResp, err error) {
	// TODO: Your code here...
	return
}

// List implements the ItemServiceImpl interface.
func (s *ItemServiceImpl) List(ctx context.Context, req *item.ListReq) (resp *item.ListResp, err error) {
	// TODO: Your code here...
	return
}

// MGet2C implements the ItemServiceImpl interface.
func (s *ItemServiceImpl) MGet2C(ctx context.Context, req *item.MGet2CReq) (resp *item.MGet2CResp, err error) {
	// TODO: Your code here...
	return
}

// DecrStock implements the ItemServiceImpl interface.
func (s *ItemServiceImpl) DecrStock(ctx context.Context, req *item.DecrStockReq) (resp *item.DecrStockResp, err error) {
	// TODO: Your code here...
	return
}

// DecrStockRevert implements the ItemServiceImpl interface.
func (s *ItemServiceImpl) DecrStockRevert(ctx context.Context, req *item.DecrStockReq) (resp *item.DecrStockResp, err error) {
	// TODO: Your code here...
	return
}
