package service

import (
	"context"
	"fmt"
	"strconv"

	"github.com/baiyutang/gomall/app/checkout/infra/rpc"
	"github.com/baiyutang/gomall/app/checkout/kitex_gen/cart"
	checkout "github.com/baiyutang/gomall/app/checkout/kitex_gen/checkout"
	"github.com/baiyutang/gomall/app/checkout/kitex_gen/order"
	"github.com/baiyutang/gomall/app/checkout/kitex_gen/payment"
	"github.com/baiyutang/gomall/app/checkout/kitex_gen/product"
	"github.com/cloudwego/kitex/pkg/klog"
)

type CheckoutService struct {
	ctx context.Context
} // NewCheckoutService new CheckoutService
func NewCheckoutService(ctx context.Context) *CheckoutService {
	return &CheckoutService{ctx: ctx}
}

/*
	Run

1. get cart
2. calculate cart
3. create order
4. empty cart
5. pay
6. change order result
7. finish
*/
func (s *CheckoutService) Run(req *checkout.CheckoutReq) (resp *checkout.CheckoutRes, err error) {
	// Finish your business logic.
	// Idempotent
	// get cart
	cartResult, err := rpc.CartClient.GetCart(s.ctx, &cart.GetCartRequest{UserId: req.UserId})
	if err != nil {
		klog.Error(err)
	}
	var oi []*order.OrderItem
	for _, cartItem := range cartResult.Items {
		p, err := rpc.ProductClient.GetProduct(s.ctx, &product.GetProductRequest{Id: cartItem.ProductId})
		if err != nil {
			klog.Error(err)
		}
		oi = append(oi, &order.OrderItem{Item: &cart.CartItem{ProductId: cartItem.ProductId, Quantity: cartItem.Quantity}, Cost: p.Price * float32(cartItem.Quantity)})
	}
	// create order
	orderReq := &order.PlaceOrderRequest{
		UserId: req.UserId,
	}
	if req.Address != nil {
		addr := req.Address
		zipCodeInt, _ := strconv.Atoi(addr.ZipCode)
		orderReq.Address = &order.Address{
			StreetAddress: addr.StreetAddress,
			City:          addr.City,
			Country:       addr.Country,
			State:         addr.State,
			ZipCode:       int32(zipCodeInt),
		}
	}
	orderResult, err := rpc.OrderClient.PlaceOrder(s.ctx, orderReq)
	if err != nil {
		klog.Error(err)
	}
	// empty cart
	emptyResult, err := rpc.CartClient.EmptyCart(s.ctx, &cart.EmptyCartRequest{UserId: req.UserId})
	if err != nil {
		klog.Error(err)
	}
	fmt.Println(emptyResult)
	// charge
	paymentResult, err := rpc.PaymentClient.Charge(s.ctx, &payment.ChargeRequest{UserId: req.UserId, OrderId: orderResult.Order.OrderId, Amount: 1, CreditCard: &payment.CreditCardInfo{
		CreditCardNumber:          req.CreditCard.CreditCardNumber,
		CreditCardExpirationYear:  req.CreditCard.CreditCardExpirationYear,
		CreditCardExpirationMonth: req.CreditCard.CreditCardExpirationMonth,
		CreditCardCvv:             req.CreditCard.CreditCardCvv,
	}})
	if err != nil {
		klog.Error(err)
	}

	fmt.Println(paymentResult)
	// change order state
	fmt.Println(orderResult)

	return
}
