package model

import "time"

//Order --> Model to entity Order
type Order struct {
	OrderId        int       `json:"order_id" gorm:"primaryKey"`
	UserId         int       `json:"user_id"`
	OrderDetail    string    `json:"order_detail"`
	OrderPrice     float32   `json:"order_price"`
	OrderCreatedAt time.Time `json:"order_created_at"`
	OrderStatus    int       `json:"order_status"`       // status : 1: created, 2: accepted, 3:done, 4: paid? 
}

//TableName --> Table for Order Model
func (Order) TableName() string {
	return "orders"
}
