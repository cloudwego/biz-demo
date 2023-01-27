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

package differ

import (
	"github.com/cloudwego/biz-demo/book-shop/app/item/common/po"
	"github.com/r3labs/diff/v2"
)

type productPODiffer struct{}

var ProductPODiffer *productPODiffer

func (differ *productPODiffer) GetChangedMap(origin, target *po.Product) map[string]interface{} {
	d, _ := diff.NewDiffer(diff.TagName("json"))
	changedMap := make(map[string]interface{})
	changeLog, _ := d.Diff(origin, target)
	for _, change := range changeLog {
		if depth := len(change.Path); depth != 1 {
			continue
		}
		if change.Type == diff.UPDATE {
			changedMap[change.Path[0]] = change.To
		}
	}
	return changedMap
}
