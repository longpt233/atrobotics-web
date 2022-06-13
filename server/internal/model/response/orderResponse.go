package response

import (
	"time"
)

type OrderResponse struct {
	OrderId        string      `json:"id" `
	UserId         string      `json:"userId"`
	OrderItems     interface{} `json:"details"`
	OrderCreatedAt time.Time   `json:"createdAt"`
	OrderStatus    int         `json:"status"` // status : 1: created, 2: accepted, 3:done, 4: paid?
	OrderCode      string      `json:"orderCode"`
	OrderAddress   string      `json:"orderAddress"`
}
