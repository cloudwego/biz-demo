package model

type Consignee struct {
	Email string

	StreetAddress string
	City          string
	State         string
	Country       string
	ZipCode       int32
}

type Order struct {
	Base
	UserId       uint32
	UserCurrency string
	Consignee    Consignee `gorm:"embedded"`
}

func (o Order) TableName() string {
	return "order"
}

type OrderState string

const (
	OrderStatePlaced    OrderState = "placed"
	OrderStatePayed     OrderState = "payed"
	OrderStateCanceled  OrderState = "canceled"
	OrderStateDelivered OrderState = "delivered"
	OrderStateReceived  OrderState = "received"
)
