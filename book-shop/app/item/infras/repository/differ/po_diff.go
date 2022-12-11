package differ

import (
	"github.com/cloudwego/biz-demo/book-shop/app/item/common/po"
	"github.com/r3labs/diff/v2"
)

type productPODiffer struct {
}

var ProductPODiffer *productPODiffer

func (differ *productPODiffer) GetChangedMap(origin *po.Product, target *po.Product) map[string]interface{} {
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
