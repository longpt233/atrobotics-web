package model

import "time"

type CardItems struct {
	CartId        string    `json:"id" gorm:"column:cart_id" gorm:"primaryKey"`
	CartUserId    string    `json:"userId" gorm:"column:cart_user_id"`
	CartProductId string    `json:"productId" gorm:"column:cart_product_id"`
	CartQuantity  int       `json:"quantity" gorm:"column:cart_quantity"`
	CartCreatedAt time.Time `json:"createdAt" gorm:"column:cart_created_at"`
	CartUpdatedAt time.Time `json:"updateAt" gorm:"column:cart_updated_at"`
	CartState     int       `json:"state" gorm:"column:cart_state"` 
	//State:  0: in cart, 1: wait for confirm
			// 2: Delivering, 3: Delivered 
			// 4: Cancelled
}

func (CardItems) TableName() string {
	return "cart_items"
}
