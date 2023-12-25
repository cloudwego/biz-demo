package types

type AddCartReq struct {
	ProductId uint32 `json:"product_id" form:"product_id"`
	Quantity  int32  `json:"quantity" form:"quantity"`
}

type CartItem struct {
	ProductId   uint32
	ProductName string
	ProductImg  string
	Quantity    int32
}
