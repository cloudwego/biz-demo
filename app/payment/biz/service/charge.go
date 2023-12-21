package service

import (
	"context"
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
	// Finish your business logic.
	card := creditcard.Card{
		Number: req.CreditCard.CreditCardNumber,
		Cvv:    string(req.CreditCard.CreditCardCvv),
		Month:  string(req.CreditCard.CreditCardExpirationMonth),
		Year:   string(req.CreditCard.CreditCardExpirationYear),
	}

	err = card.Validate(true)

	if err != nil {
		return nil, kerrors.NewBizStatusError(400, err.Error())
	}

	translationId, err := uuid.NewRandom()
	if err != nil {
		return nil, err
	}
	return &payment.ChargeResponse{TransactionId: translationId.String()}, nil
}
