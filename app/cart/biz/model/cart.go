package model

import (
	"context"

	"gorm.io/gorm"
)

type CartItem struct {
	Base
	UserId    uint32
	ProductId uint32
	Quantity  int32
}

func (c CartItem) TableName() string {
	return "cart_item"
}

func Add(db *gorm.DB, ctx context.Context, item *CartItem) (err error) {
	return db.WithContext(ctx).Create(item).Error
}

func Empty(db *gorm.DB, ctx context.Context, userId uint32) (err error) {
	return db.WithContext(ctx).Where("user_id = ?", userId).Delete(&CartItem{}).Error
}

func GetCartList(db *gorm.DB, ctx context.Context, userId uint32) (list []*CartItem, err error) {
	err = db.WithContext(ctx).Where("user_id = ?", userId).Find(&list).Error
	return
}
