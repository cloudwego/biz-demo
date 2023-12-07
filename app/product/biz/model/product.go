package model

import "gorm.io/gorm"

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

func GetProductsByCategories(db *gorm.DB, categoryNames []string) (p []Product, err error) {
	var categories []Category
	for _, categoryName := range categoryNames {
		categories = append(categories, Category{Name: categoryName})
	}
	err = db.Model(&Product{Categories: categories}).Find(&p).Error
	return p, err
}
