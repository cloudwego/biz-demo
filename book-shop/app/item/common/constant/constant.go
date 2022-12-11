package constant

type ProductStatus = int64

const (
	ProductStatusOnline  ProductStatus = 0 // 商品上架
	ProductStatusOffline ProductStatus = 1 // 商品下架
	ProductStatusDelete  ProductStatus = 2 // 商品删除
)

var ProductStatusDescMap = map[ProductStatus]string{
	ProductStatusOnline:  "上架",
	ProductStatusOffline: "下架",
	ProductStatusDelete:  "删除",
}

type StateOperationType = int64

const (
	StateOperationTypeAdd     StateOperationType = 1 //新建商品
	StateOperationTypeSave    StateOperationType = 2 //保存商品
	StateOperationTypeDel     StateOperationType = 3 //删除商品
	StateOperationTypeOffline StateOperationType = 4 //商品下架
	StateOperationTypeOnline  StateOperationType = 5 //商品上架
)
