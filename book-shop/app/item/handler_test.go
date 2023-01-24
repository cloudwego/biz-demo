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
	"github.com/cloudwego/biz-demo/book-shop/kitex_gen/cwg/bookshop/item"
	"github.com/cloudwego/hertz/pkg/common/test/assert"
	"testing"
)

func TestItemServiceImpl(t *testing.T) {
	Init()

	t.Log("===Add Test Begin===")
	addReq := &item.AddReq{
		Name:        "Java核心技术卷1-双11大促",
		Pic:         "http://www.wld5.com/uploads/allimg/191115/030F13017-0.png",
		Description: "Java核心技术卷1",
		Property: &item.BookProperty{
			Isbn:     "9787115188335",
			SpuName:  "Java核心技术卷1",
			SpuPrice: 5000,
		},
		Price: 5000,
		Stock: 100,
	}
	serviceImpl := &ItemServiceImpl{}
	addResp, err := serviceImpl.Add(context.TODO(), addReq)
	t.Logf("resp: %v\n, err: %v", addResp, err)
	assert.Nil(t, err)
	t.Log("===Add Test Pass===")

	t.Log("===Edit Test Begin===")
	editName := "Java核心技术卷1-年货节大促"
	editReq := &item.EditReq{
		ProductId: addResp.ProductId,
		Name:      &editName,
	}
	editResp, err := serviceImpl.Edit(context.TODO(), editReq)
	t.Logf("resp: %v\n, err: %v", editResp, err)
	assert.Nil(t, err)
	t.Log("===Edit Test Pass===")

	t.Log("===Offline Test Begin===")
	offlineReq := &item.OfflineReq{ProductId: addResp.ProductId}
	offlineResp, err := serviceImpl.Offline(context.TODO(), offlineReq)
	t.Logf("resp: %v\n, err: %v", offlineResp, err)
	assert.Nil(t, err)
	t.Log("===Offline Test Pass===")

	t.Log("===Online Test Begin===")
	onlineReq := &item.OnlineReq{ProductId: addResp.ProductId}
	onlineResp, err := serviceImpl.Online(context.TODO(), onlineReq)
	t.Logf("resp: %v\n, err: %v", onlineResp, err)
	assert.Nil(t, err)
	t.Log("===Online Test Pass===")

	t.Log("===List Test Begin===")
	spuName := "Java核心技术卷1"
	status := item.Status_Online
	listReq := &item.ListReq{
		SpuName: &spuName,
		Status:  &status,
	}
	listResp, err := serviceImpl.List(context.TODO(), listReq)
	t.Logf("resp: %v\n, err: %v", listResp, err)
	assert.Nil(t, err)
	t.Log("===List Test Pass===")

	t.Log("===Get Test Begin===")
	getReq := &item.GetReq{
		ProductId: addResp.ProductId,
	}
	getResp, err := serviceImpl.Get(context.TODO(), getReq)
	t.Logf("resp: %v\n, err: %v", getResp, err)
	assert.Nil(t, err)
	t.Log("===Get Test Pass===")

	t.Log("===DecrStock Test Begin===")
	decrStockReq := &item.DecrStockReq{
		ProductId: addResp.ProductId,
		StockNum:  2,
	}
	decrStockResp, err := serviceImpl.DecrStock(context.TODO(), decrStockReq)
	t.Logf("resp: %v\n, err: %v", decrStockResp, err)
	assert.Nil(t, err)
	t.Log("===DecrStock Test Pass===")

	t.Log("===DecrStockRevert Test Begin===")
	decrStockRevertReq := &item.DecrStockReq{
		ProductId: addResp.ProductId,
		StockNum:  2,
	}
	decrStockRevertResp, err := serviceImpl.DecrStockRevert(context.TODO(), decrStockRevertReq)
	t.Logf("resp: %v\n, err: %v", decrStockRevertResp, err)
	assert.Nil(t, err)
	t.Log("===DecrStockRevert Test Pass===")

	t.Log("===Delete Test Begin===")
	deleteReq := &item.DeleteReq{ProductId: addResp.ProductId}
	deleteResp, err := serviceImpl.Delete(context.TODO(), deleteReq)
	t.Logf("resp: %v\n, err: %v", deleteResp, err)
	assert.Nil(t, err)
	t.Log("===Delete Test Pass===")

	t.Log("===Get After Delete Test Begin===")
	getReq = &item.GetReq{
		ProductId: addResp.ProductId,
	}
	getResp, err = serviceImpl.Get(context.TODO(), getReq)
	t.Logf("resp: %v\n, err: %v", getResp, err)
	assert.Nil(t, err)
	t.Log("===Get After Delete Test Pass===")

	t.Log("===Search Test Begin===")
	searchName := "Java"
	searchReq := &item.SearchReq{
		Name: &searchName,
	}
	searchResp, err := serviceImpl.Search(context.TODO(), searchReq)
	t.Logf("resp: %v\n, err: %v", searchResp, err)
	assert.Nil(t, err)
	t.Log("===Search Test Pass===")

	t.Log("===Get2C Test Begin===")
	get2CReq := &item.MGet2CReq{
		ProductIds: []int64{addResp.ProductId},
	}
	get2CResp, err := serviceImpl.MGet2C(context.TODO(), get2CReq)
	t.Logf("resp: %v\n, err: %v", get2CResp, err)
	assert.Nil(t, err)
	t.Log("===Get2C Test Pass===")

}
