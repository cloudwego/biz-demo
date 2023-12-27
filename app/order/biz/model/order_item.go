package model

type OrderItem struct {
	Base
	ProductId    uint32
	OrderIdRefer string `gorm:"size:256;index"`
	Quantity     int32
	Cost         float32
}

func (oi OrderItem) TableName() string {
	return "order_item"
}
