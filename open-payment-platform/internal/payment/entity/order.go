package entity

type Order struct {
	ID              int
	MerchantID      string
	Channel         string
	PayWay          string
	OrderStatus     int8
	OutOrderNo      string
	TotalAmount     uint64
	Body            string
	AuthCode        string
	WxAppid         string
	SubOpenid       string
	JumpURL         string
	NotifyURL       string
	ClientIP        string
	Attach          string
	OrderExpiration string
	ExtendParams    string
}

func NewOrder() *Order {
	return &Order{}
}
