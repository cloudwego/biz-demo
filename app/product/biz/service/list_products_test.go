package service

import (
	"context"
	"testing"
	product "github.com/baiyutang/gomall/app/product/kitex_gen/product"
)

func TestListProducts_Run(t *testing.T) {
	ctx := context.Background()
	s := NewListProductsService(ctx)
	// init req and assert value

	req := &product.Empty{}
	resp, err := s.Run(req)
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}
	if resp == nil {
		t.Errorf("unexpected nil response")
	}
	// todo: edit your unit test

}