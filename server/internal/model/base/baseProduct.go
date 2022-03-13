package base

type BaseProduct struct{
	ProductName       string    `json:"product_name"`
	ProductPrice      float64   `json:"product_price"`
	ProductShortDesc  string    `json:"product_short_desc"`
	ProductLongDesc   string    `json:"product_long_desc"`
	ProductCategoryID int       `json:"product_category_id"`
	ProductAvailable  int       `json:"product_available"`
	ProductBrand      string    `json:"product_brand"`
	ProductSold       int       `json:"product_sold"`
}