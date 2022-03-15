package request

type OrderRequest struct {
	ProductOrders []string `json:"orders"`
	StatusOrder   int      `json:"order_status"`
}
