package request

type CartItemsRequest struct {
	CartProductId string `json:"productId"`
	CartQuantity  int    `json:"quantity"`
	CartColor     string `json:"color"`
}
