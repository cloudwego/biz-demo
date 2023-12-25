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
