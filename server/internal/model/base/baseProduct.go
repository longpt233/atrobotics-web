package base

type BaseProduct struct {
	ProductName       string  `json:"name" gorm:"column:product_name"`
	ProductPrice      float64 `json:"price" gorm:"column:product_price"`
	ProductShortDesc  string  `json:"shortDesc" gorm:"column:product_short_desc"`
	ProductLongDesc   string  `json:"longDesc" gorm:"column:product_long_desc"`
	ProductCategoryID string  `json:"categoryId" gorm:"column:product_category_id"`
	ProductAvailable  int     `json:"available" gorm:"column:product_available"`
	ProductBrand      string  `json:"brand" gorm:"column:product_brand"`
	ProductSold       int     `json:"sold" gorm:"column:product_sold"`
}
