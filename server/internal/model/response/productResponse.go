package response

import (
	"atro/internal/model/base"
	"time"
)

type ProductResponse struct {
	base.BaseProduct
	ProductID         int       `json:"product_id" gorm:"primaryKey"`
	ProductImages     []string    `json:"product_images"`
	ProductUpdatedAt  time.Time `json:"product_updated_at"`
	ProductColor      []string    `json:"product_color"`
	ProductCreatedAt  time.Time `json:"product_created_at"`
}