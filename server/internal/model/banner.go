package model

import "time"

type Banner struct {
	BannerId        string    `json:"banner_id" gorm:"primaryKey"`
	BannerProductId string    `json:"banner_product_id"`
	BannerImage     string    `json:"banner_image"`
	BannerUpdateAt  time.Time   `json:"banner_update_at"`
	BannerCreateAt time.Time `json:"banner_create_at"`
}
