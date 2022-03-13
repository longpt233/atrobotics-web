package model

type ProductCategory struct {
	ProductCategoryID int    `json:"product_category_id" gorm:"primaryKey"`
	CategoryName      string `json:"category_name"`
}

func (ProductCategory) TableName() string {
	return "product_categories"
}
