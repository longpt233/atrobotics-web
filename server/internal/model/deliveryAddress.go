package model

import "time"

type DeliveryAddress struct {
	DeliveryAddressId string    `json:"id" gorm:"primaryKey" gorm:"column:delivery_address_id"`
	UserId            string    `json:"userId" gorm:"column:user_id"`
	Fullname          string    `json:"fullname" gorm:"column:fullname"`
	Phone             string    `json:"phone" gorm:"column:phone"`
	City              string    `json:"city" gorm:"column:city"`
	District          string    `json:"district" gorm:"column:district"`
	Ward              string    `json:"ward" gorm:"column:ward"`
	DetailAddress     string    `json:"detailAddress" gorm:"column:detail_address"`
	CreatedAt         time.Time `json:"createdAt" gorm:"column:created_at"`
	UpdatedAt         time.Time `json:"updatedAt" gorm:"column:updated_at"`
}

func (DeliveryAddress) TableName() string {
	return "delivery_address"
}