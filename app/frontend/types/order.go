package types

type Order struct {
	OrderId    string
	CreatedAt  int32
	OrderState string
	Cost       float32
	Items      []OrderItem
}

type OrderItem struct {
	OrderId     string
	ProductId   uint32
	ProductName string
	Qty         uint32
	Cost        float32
}
