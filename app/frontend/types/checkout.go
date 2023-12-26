package types

type CheckoutForm struct {
	Email           string `json:"email" form:"email"`
	Firstname       string `json:"firstname" form:"firstname"`
	Lastname        string `json:"lastname" form:"lastname"`
	Street          string `json:"street" form:"street"`
	Zipcode         string `json:"zipcode" form:"city"`
	Province        string `json:"province" form:"province"`
	Country         string `json:"country" form:"country"`
	City            string `json:"city" form:"city"`
	CardNum         string `json:"cardNum" form:"cardNum"`
	ExpirationMonth int32  `json:"expirationMonth" form:"expirationMonth"`
	ExpirationYear  int32  `json:"expirationYear" form:"expirationYear"`
	Cvv             int32  `json:"cvv" form:"cvv"`
	Payment         string `json:"payment" form:"payment"`
}
