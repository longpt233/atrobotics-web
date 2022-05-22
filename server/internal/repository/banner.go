package repository

import (
	"atro/internal/model"

	"github.com/jinzhu/gorm"
)

type BannerRepository interface {
	GetBanner(string) (model.Banner, error)
	AddBanner(model.Banner) (model.Banner, error)
	UpdateBanner(model.Banner) (model.Banner, error)
	DeleteBanner(string) (model.Banner, error)
	GetTop3NewestBanner()([]model.Banner, error)
}

type bannerRepository struct{
	connection *gorm.DB
}

func NewBannerRepository() BannerRepository{
	myclient := &MySQLClient{}
	return &bannerRepository{
		connection:myclient.GetConn(),
	}
}

func (db *bannerRepository) GetBanner(id string) (banner model.Banner, err error){
	return banner, db.connection.First(&banner, "banner_id=?", id).Error
}
func (db *bannerRepository) AddBanner(banner model.Banner) (model.Banner, error){
	return banner, db.connection.Create(&banner).Error
}
func (db *bannerRepository) UpdateBanner(banner model.Banner) (model.Banner, error){
	var checkBanner model.Banner
	if err := db.connection.First(&checkBanner,"banner_id=?",banner.BannerId).Error; err != nil{
		return checkBanner, err
	}
	banner.BannerCreateAt = checkBanner.BannerCreateAt
	return banner, db.connection.Model(&banner).Where(model.Banner{BannerId: banner.BannerId}).Updates(&banner).Error
}
func (db *bannerRepository) DeleteBanner(id string) (model.Banner, error){
	var banner model.Banner
	if err := db.connection.First(&banner,"banner_id=?",id).Error; err != nil{
		return banner, err
	}
	return banner, db.connection.Delete(&banner, "banner_id=?", banner.BannerId).Error

}
func (db *bannerRepository) GetTop3NewestBanner()(listBanner []model.Banner, err error){
	return listBanner, db.connection.Order("banner_create_at DESC").Limit(3).Find(&listBanner).Error
}
