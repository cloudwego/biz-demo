package model

import (
	"gorm.io/gorm"
)

type Product struct {
	Base
	Name        string     `json:"name"`
	Description string     `json:"description"`
	Picture     string     `json:"picture"`
	Price       float32    `json:"price"`
	Categories  []Category `json:"categories" gorm:"many2many:product_category"`
}

func (p Product) TableName() string {
	return "product"
}

func GetProductById(db *gorm.DB, productId int) (product Product, err error) {
	err = db.Model(&Product{}).Where(&Product{Base: Base{ID: productId}}).First(&product).Error
	return product, err
}

func SearchProduct(db *gorm.DB, q string) (product []*Product, err error) {
	err = db.Model(&Product{}).Find(&product, "name like ? or description like ?", "%"+q+"%", "%"+q+"%").Error
	return product, err
}
