package request

type NewProductForm struct {
	ProductName       string   `json:"product_name" binding:"required"`
	ProductPrice      float64  `json:"product_price" binding:"required"`
	ProductShortDesc  string   `json:"product_short_desc" binding:"required"`
	ProductLongDesc   string   `json:"product_long_desc"`
	ProductImages     []string `json:"product_images"`
	ProductCategoryID int      `json:"product_category_id" binding:"required"`
	ProductAvailable  int      `json:"product_available"`
	ProductColor      []string `json:"product_color"`
	ProductBrand      string   `json:"product_brand"`
	ProductSold       int      `json:"product_sold"`
}
