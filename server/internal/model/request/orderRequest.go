package request

type OrderRequest struct {
	ProductOrders []OrderObject `json:"orders"`
	TypeOrder     int           `json:"orderType"`
}

type OrderObject struct {
	ProductId string `json:"productId"`
	Quantity  int    `json:"quantity"`
}
