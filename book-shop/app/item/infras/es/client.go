// Copyright 2023 CloudWeGo Authors
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

package es

import (
	"context"
	"strconv"
	"sync"

	"github.com/bytedance/sonic"
	"github.com/cloudwego/biz-demo/book-shop/app/item/common/entity"
	"github.com/cloudwego/biz-demo/book-shop/pkg/conf"
	"github.com/olivere/elastic/v7"
)

// ES client
var (
	esOnce sync.Once
	esCli  *elastic.Client
)

// GetESClient get ES client
func GetESClient() *elastic.Client {
	if esCli != nil {
		return esCli
	}

	esOnce.Do(func() {
		cli, err := elastic.NewSimpleClient(
			elastic.SetURL(conf.ESAddress),
		)
		if err != nil {
			panic("new es client failed, err=" + err.Error())
		}
		esCli = cli
	})
	return esCli
}

func UpsertProductES(ctx context.Context, productId int64, product *entity.ProductEntity) error {
	doc := getDocFromEntity(product)
	_, err := GetESClient().Update().Index(conf.ProductESIndex).Id(strconv.FormatInt(productId, 10)).Doc(doc).Upsert(doc).Refresh("true").Do(ctx)
	return err
}

func BatchGetProductById(ctx context.Context, productIds []int64) ([]*entity.ProductEntity, error) {
	mgetSvc := GetESClient().MultiGet()
	for _, id := range productIds {
		mgetSvc.Add(elastic.NewMultiGetItem().
			Index(conf.ProductESIndex).
			Id(strconv.FormatInt(id, 10)))
	}
	rsp, err := mgetSvc.Do(ctx)
	if err != nil {
		return nil, err
	}
	entities := make([]*entity.ProductEntity, 0)
	for _, doc := range rsp.Docs {
		entities = append(entities, getEntityFromSource(string(doc.Source)))
	}
	return entities, nil
}

func SearchProduct(ctx context.Context, filter map[string]interface{}) ([]*entity.ProductEntity, error) {
	boolQuery := elastic.NewBoolQuery()
	for k, v := range filter {
		boolQuery.Must(elastic.NewMatchQuery(k, v))
	}
	searchResult, err := GetESClient().Search().
		Index(conf.ProductESIndex).
		Query(boolQuery).
		Size(10000).
		From(0).
		Do(ctx)
	if err != nil {
		return nil, err
	}
	if searchResult.Hits.TotalHits.Value <= 0 || len(searchResult.Hits.Hits) <= 0 {
		return []*entity.ProductEntity{}, nil
	}
	ret := make([]*entity.ProductEntity, 0)
	for _, hit := range searchResult.Hits.Hits {
		ret = append(ret, getEntityFromSource(string(hit.Source)))
	}
	return ret, nil
}

func getEntityFromSource(source string) *entity.ProductEntity {
	sourceMap := make(map[string]interface{})
	_ = sonic.UnmarshalString(source, &sourceMap)
	ret := &entity.ProductEntity{
		ProductId:   int64(sourceMap["product_id"].(float64)),
		Name:        sourceMap["name"].(string),
		Pic:         sourceMap["pic"].(string),
		Description: sourceMap["description"].(string),
		Property: &entity.PropertyEntity{
			ISBN:     sourceMap["isbn"].(string),
			SpuName:  sourceMap["spu_name"].(string),
			SpuPrice: int64(sourceMap["spu_price"].(float64)),
		},
		Price:  int64(sourceMap["price"].(float64)),
		Stock:  int64(sourceMap["stock"].(float64)),
		Status: int64(sourceMap["status"].(float64)),
	}
	return ret
}

func getDocFromEntity(e *entity.ProductEntity) map[string]interface{} {
	ret := map[string]interface{}{
		"product_id":  e.ProductId,
		"name":        e.Name,
		"pic":         e.Pic,
		"description": e.Description,
		"price":       e.Price,
		"stock":       e.Stock,
		"status":      e.Status,
	}
	if e.Property != nil {
		ret["isbn"] = e.Property.ISBN
		ret["spu_name"] = e.Property.SpuName
		ret["spu_price"] = e.Property.SpuPrice
	}
	return ret
}
