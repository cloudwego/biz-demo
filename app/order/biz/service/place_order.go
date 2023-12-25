package service

import (
	"context"
	"fmt"

	"github.com/baiyutang/gomall/app/order/biz/dal/mysql"
	"github.com/baiyutang/gomall/app/order/biz/model"
	order "github.com/baiyutang/gomall/app/order/kitex_gen/order"
	"gorm.io/gorm"
)

type PlaceOrderService struct {
	ctx context.Context
} // NewPlaceOrderService new PlaceOrderService
func NewPlaceOrderService(ctx context.Context) *PlaceOrderService {
	return &PlaceOrderService{ctx: ctx}
}

// Run create note info
func (s *PlaceOrderService) Run(req *order.PlaceOrderRequest) (resp *order.PlaceOrderResponse, err error) {
	// Finish your business logic.
	if len(req.OrderItems) == 0 {
		err = fmt.Errorf("OrderItems empty")
		return
	}
	// 扣减库存
	mysql.DB.Transaction(func(tx *gorm.DB) error {
		o := &model.Order{
			UserId:       req.UserId,
			UserCurrency: req.UserCurrency,
			Consignee: model.Consignee{
				Email: req.Email,
			},
		}
		if err := tx.Create(o).Error; err != nil {
			return err
		}

		var itemList []*model.OrderItem
		for _, v := range req.OrderItems {
			itemList = append(itemList, &model.OrderItem{OrderID: o.ID, ProductId: v.Item.ProductId, Quantity: v.Item.Quantity})
		}
		if err := tx.Create(&itemList).Error; err != nil {
			return err
		}
		return nil
	})

	return
}
