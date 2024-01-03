package service

import (
	"context"
	"strconv"
	"time"

	"github.com/baiyutang/gomall/app/payment/biz/dal/mysql"
	"github.com/baiyutang/gomall/app/payment/biz/model"
	payment "github.com/baiyutang/gomall/app/payment/kitex_gen/payment"
	"github.com/cloudwego/kitex/pkg/kerrors"
	creditcard "github.com/durango/go-credit-card"
	"github.com/google/uuid"
)

type ChargeService struct {
	ctx context.Context
} // NewChargeService new ChargeService
func NewChargeService(ctx context.Context) *ChargeService {
	return &ChargeService{ctx: ctx}
}

// Run create note info
func (s *ChargeService) Run(req *payment.ChargeRequest) (resp *payment.ChargeResponse, err error) {
	card := creditcard.Card{
		Number: req.CreditCard.CreditCardNumber,
		Cvv:    strconv.Itoa(int(req.CreditCard.CreditCardCvv)),
		Month:  strconv.Itoa(int(req.CreditCard.CreditCardExpirationMonth)),
		Year:   strconv.Itoa(int(req.CreditCard.CreditCardExpirationYear)),
	}

	err = card.Validate(true)

	if err != nil {
		return nil, kerrors.NewBizStatusError(400, err.Error())
	}

	translationId, err := uuid.NewRandom()
	err = model.CreatePaymentLog(mysql.DB, s.ctx, &model.PaymentLog{
		UserId:        req.UserId,
		OrderId:       req.OrderId,
		TransactionId: translationId.String(),
		Amount:        req.Amount,
		PayAt:         time.Now(),
	})
	if err != nil {
		return nil, err
	}
	if err != nil {
		return nil, err
	}
	return &payment.ChargeResponse{TransactionId: translationId.String()}, nil
}
