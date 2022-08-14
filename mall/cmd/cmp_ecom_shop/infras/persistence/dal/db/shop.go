package db

import (
	"context"
	"github.com/cloudwego/biz-demo/mall/pkg/conf"
	"gorm.io/gorm"
)

type ShopPO struct {
	gorm.Model
	ShopId   int64  `json:"shop_id"`
	ShopName string `json:"shop_name"`
	UserId   uint   `json:"user_id"`
}

func (shop *ShopPO) TableName() string {
	return conf.ShopTableName
}

func GetShopInfoByUserId(ctx context.Context, userId int64) (*ShopPO, error) {
	ret := &ShopPO{}
	if err := DB.WithContext(ctx).Where("user_id = ?", uint(userId)).Find(ret).Error; err != nil {
		return nil, err
	}
	return ret, nil
}

func CreateShop(ctx context.Context, shop *ShopPO) error {
	return DB.WithContext(ctx).Create(shop).Error
}
