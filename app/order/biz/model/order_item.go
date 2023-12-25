package model

type OrderItem struct {
	Base
	OrderID   int
	ProductId uint32
	Quantity  int32
	Cost      float32
}

func (oi OrderItem) TableName() string {
	return "order_item"
}
