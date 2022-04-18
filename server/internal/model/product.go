package model

import (
	"atro/internal/model/base"
	"time"
)

type Product struct {
	base.BaseProduct
	ProductID        string    `json:"id" gorm:"column:product_id" gorm:"primaryKey" `
	ProductImages    string    `json:"images" gorm:"column:product_images"`
	ProductUpdatedAt time.Time `json:"updatedAt" gorm:"column:product_updated_at"`
	ProductColor     string    `json:"color" gorm:"column:product_color"`
	ProductCreatedAt time.Time `json:"createdAt" gorm:"column:product_created_at"`
}

func (Product) TableName() string {
	return "products"
}
