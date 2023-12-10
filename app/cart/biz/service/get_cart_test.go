package service

import (
	"context"
	"testing"
	cart "github.com/baiyutang/gomall/app/cart/kitex_gen/cart"
)

func TestGetCart_Run(t *testing.T) {
	ctx := context.Background()
	s := NewGetCartService(ctx)
	// init req and assert value

	req := &cart.GetCartRequest{}
	resp, err := s.Run(req)
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}
	if resp == nil {
		t.Errorf("unexpected nil response")
	}
	// todo: edit your unit test

}
