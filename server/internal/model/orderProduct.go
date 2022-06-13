package model

type OrderProduct struct {
	ProductId        string `json:"productId"`
	Quantity         int `json:"quantity"`
	ProductName      string `json:"productName"`
	CurrentPrice     float64 `json:"currentPrice"`
	ProductImage     string `json:"productImage"`
	ShortDescription string `json:"shortDescription"`
	Colors           string `json:"colors"`
}
