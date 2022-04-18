package model

import "time"

type Banner struct {
	BannerId        string    `json:"id" gorm:"primaryKey" gorm:"column:banner_id"`
	BannerProductId string    `json:"productId" gorm:"column:banner_product_id"`
	BannerImage     string    `json:"image" gorm:"column:banner_image"`
	BannerUpdateAt  time.Time `json:"updateAt" gorm:"column:banner_update_at"`
	BannerCreateAt  time.Time `json:"createAt" gorm:"column:banner_create_at"`
}
