package request

import "atro/internal/model/base"

type NewProductForm struct {
	base.BaseProduct
	ProductImages []string `json:"product_images"`
	ProductColor  []string `json:"product_color"`
}
