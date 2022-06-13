package request

type DeliveryAddressRequest struct {
	Fullname          string    `json:"fullname" `
	Phone             string    `json:"phone" `
	City              string    `json:"city"`
	District          string    `json:"district" `
	Ward              string    `json:"ward" `
	DetailAddress     string    `json:"detailAddress"`
}
