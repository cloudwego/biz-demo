package types

type Consignee struct {
	Email string

	StreetAddress string
	City          string
	State         string
	Country       string
	ZipCode       int32
}

type Order struct {
	Consignee   Consignee
	OrderId     string
	CreatedDate string
	OrderState  string
	Cost        float32
	Items       []OrderItem
}

type OrderItem struct {
	ProductId   uint32
	ProductName string
	Picture     string
	Qty         uint32
	Cost        float32
}
