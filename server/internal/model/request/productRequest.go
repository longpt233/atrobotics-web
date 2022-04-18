package request

import (
	"atro/internal/model"
	"atro/internal/model/base"
	"encoding/json"
	"time"
)

type ProductRequest struct {
	base.BaseProduct
	ProductImages []string `json:"productImages"`
	ProductColor  []string `json:"productColor"`
}

func (p *ProductRequest) ProductRequestToProduct() (model.Product, error) {
	productImg, err := json.Marshal(p.ProductImages)
	if err != nil {
		return model.Product{}, err
	}
	productCol, err := json.Marshal(p.ProductColor)
	if err != nil {
		return model.Product{}, err
	}
	rsProduct := model.Product{
		ProductImages:    string(productImg),
		ProductColor:     string(productCol),
		ProductCreatedAt: time.Now(),
		ProductUpdatedAt: time.Now(),
		BaseProduct: base.BaseProduct{
			ProductName:       p.ProductName,
			ProductPrice:      p.ProductPrice,
			ProductShortDesc:  p.ProductShortDesc,
			ProductLongDesc:   p.ProductLongDesc,
			ProductCategoryID: p.ProductCategoryID,
			ProductAvailable:  p.ProductAvailable,
			ProductSold:       p.ProductSold,
			ProductBrand:      p.ProductBrand,
		},
	}
	return rsProduct, nil
}
