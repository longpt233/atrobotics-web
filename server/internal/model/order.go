package model

import "time"

//Order --> Model to entity Order
type Order struct {
	OrderId        string    `json:"id" gorm:"primaryKey" json:"order_id"`
	UserId         string    `json:"userId" gorm:"column:user_id"`
	OrderDetail    string    `json:"detail" gorm:"column:order_detail"`
	OrderPrice     float32   `json:"price" gorm:"column:order_price"`
	OrderCreatedAt time.Time `json:"createdAt" gorm:"column:order_created_at"`
	OrderStatus    int       `json:"status" gorm:"column:order_status"` // status : 1: created, 2: accepted, 3:done, 4: paid?
}

//TableName --> Table for Order Model
func (Order) TableName() string {
	return "orders"
}
