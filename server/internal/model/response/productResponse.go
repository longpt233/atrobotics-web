package response

import (
	"atro/internal/model"
	"atro/internal/model/base"
	"encoding/json"
	"time"
)

type ProductResponse struct {
	base.BaseProduct
	ProductID        string    `json:"product_id" gorm:"primaryKey"`
	ProductImages    []string  `json:"product_images"`
	ProductUpdatedAt time.Time `json:"product_updated_at"`
	ProductColor     []string  `json:"product_color"`
	ProductCreatedAt time.Time `json:"product_created_at"`
}

func (p *ProductResponse) ProductToProductResponse(product model.Product) (ProductResponse, error) {
	var productImg []string
	var productCol []string
	if err := json.Unmarshal([]byte(product.ProductImages), &productImg); err != nil {
		return ProductResponse{}, err
	}
	if err := json.Unmarshal([]byte(product.ProductColor), &productCol); err != nil {
		return ProductResponse{}, err
	}
	*p = ProductResponse{
		ProductID:        product.ProductID,
		ProductImages:    productImg,
		ProductColor:     productCol,
		ProductUpdatedAt: product.ProductUpdatedAt,
		ProductCreatedAt: product.ProductCreatedAt,
		BaseProduct: base.BaseProduct{
			ProductName:       product.ProductName,
			ProductPrice:      product.ProductPrice,
			ProductShortDesc:  product.ProductShortDesc,
			ProductLongDesc:   product.ProductLongDesc,
			ProductCategoryID: product.ProductCategoryID,
			ProductAvailable:  product.ProductAvailable,
			ProductBrand:      product.ProductBrand,
			ProductSold:       product.ProductSold,
		},
	}
	return *p, nil
}
