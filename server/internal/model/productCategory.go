package model

type ProductCategory struct {
	ProductCategoryID string `json:"id" gorm:"primaryKey" gorm:"column:product_category_id"`
	CategoryName      string `json:"name" gorm:"column:category_name"`
}

func (ProductCategory) TableName() string {
	return "product_categories"
}
