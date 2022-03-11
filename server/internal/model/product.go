package model

import (
	"time"
)

type Product struct {
	ProductID         int       `json:"product_id" gorm:"primaryKey"`
	ProductName       string    `json:"product_name"`
	ProductPrice      float64    `json:"product_price"`
	ProductShortDesc  string    `json:"product_short_desc"`
	ProductLongDesc   string    `json:"product_long_desc"`
	ProductImages     string    `json:"product_images"`
	ProductCategoryID int       `json:"product_category_id"`
	ProductAvailable  int       `json:"product_available"`
	ProductUpdatedAt  time.Time `json:"product_updated_at"`
	ProductColor      string    `json:"product_color"`
	ProductBrand      string    `json:"product_brand"`
	ProductSold       int       `json:"product_sold"`
	ProductCreatedAt  time.Time `json:"product_created_at"`
}

func (Product) TableName() string {
	return "products"
}
