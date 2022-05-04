package model

import "time"

//Order --> Model to entity Order
type Order struct {
	OrderId        string    `json:"id" gorm:"primaryKey" gorm:"column:order_id"`
	UserId         string    `json:"userId" gorm:"column:user_id"`
	OrderItems     string    `json:"detail" gorm:"column:order_items"`
	OrderCreatedAt time.Time `json:"createdAt" gorm:"column:order_created_at"`
	OrderStatus    int       `json:"status" gorm:"column:order_status"` 
	//status: 0->CanCel 1->create(wait for accepting) 2->accepted 3->delivering 4->Success
	OrderCode      string    `json:"orderCode" gorm:"column:order_code"`
	OrderAddress string `json:"orderAddress" gorm:"column:order_address"`
}

//TableName --> Table for Order Model
func (Order) TableName() string {
	return "orders"
}
