package request

type OrderRequest struct {
	ProductOrders []OrderObject `json:"orders"`
	TypeOrder     int           `json:"order_type"`
}

type OrderObject struct {
	ProductId int `json:"product_id"`
	Quantity  int `json:"quantity"`
}
