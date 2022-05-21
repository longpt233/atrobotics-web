package model

import "time"

type CardItems struct {
	CartId          string    `json:"id" gorm:"column:cart_id" gorm:"primaryKey"`
	CartUserId      string    `json:"userId" gorm:"column:cart_user_id"`
	CartProductId   string    `json:"productId" gorm:"column:cart_product_id"`
	CartProductData Product   `json:"productData" gorm:"foreignKey:ProductID"`
	CartQuantity    int       `json:"quantity" gorm:"column:cart_quantity"`
	CartCreatedAt   time.Time `json:"createdAt" gorm:"column:cart_created_at"`
	CartUpdatedAt   time.Time `json:"updateAt" gorm:"column:cart_updated_at"`
	CartColor       string    `json:"color" gorm:"column:cart_color"`
}

func (CardItems) TableName() string {
	return "cart_items"
}
