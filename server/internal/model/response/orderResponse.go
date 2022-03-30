package response

import "atro/internal/model"

type OrderResponse struct {
	Orders       []model.Order `json:"orders"`
	OrdersLength int           `json:"size"`
}
