package repository

import (
	"atro/internal/model"

	"github.com/jinzhu/gorm"
)

type OtpRepository interface {
	AddOtp(model.Otp) (model.Otp, error)
	UpdateOtpStatus(model.Otp) (model.Otp, error)
	GetOtpByUserId(string) (model.Otp, error)
}

type otpRepository struct {
	connection *gorm.DB
}

func NewOtpRepository() OtpRepository {
	return &otpRepository{
		connection: DB(),
	}
}

func (db *otpRepository) AddOtp(otp model.Otp) (model.Otp, error){
	return otp, db.connection.Create(&otp).Error
}
func (db *otpRepository) UpdateOtpStatus(otp model.Otp) (model.Otp, error){
	var checkOtp model.Otp
	if err := db.connection.First(&checkOtp,"otp_id=?",otp.OtpId).Error; err != nil {
		return checkOtp, err
	}
	return otp, db.connection.Model(&otp).Where(model.Otp{OtpId: otp.OtpId}).Updates(&otp).Error
}

func (db *otpRepository) GetOtpByUserId(userId string) (otp model.Otp, err error){
	return otp, db.connection.Find(&otp,"user_id=? AND otp_used_ok=0", userId).Error
}