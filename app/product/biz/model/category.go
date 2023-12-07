package model

import "github.com/baiyutang/gomall/app/product/biz/dal/mysql"

type Category struct {
	Base
	Name        string    `json:"name"`
	Description string    `json:"description"`
	Product     []Product `json:"product" gorm:"many2many:product_category"`
}

func (c Category) TableName() string {
	return "category"
}

func (c Category) List() {

}

func (c Category) GetByName(name string) (category *Category, err error) {
	err = mysql.DB.Model(&Category{Name: name}).Preload("Product").First(&category).Error
	return category, err
}
