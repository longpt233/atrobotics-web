package model

import "time"

type Otp struct {
	OtpId        string    `json:"id" gorm:"primaryKey" gorm:"column:otp_id"`
	UserId       string    `json:"userId" gorm:"column:user_id"`
	OtpValue     string    `json:"otpValue" gorm:"column:otp_value"`
	OtpCreateAt  time.Time `json:"createAt" gorm:"column:otp_create_at"`
	OtpUpdateAt  time.Time `json:"updateAt" gorm:"column:otp_update_at"`
	OtpTimeStart time.Time `json:"timeStart" gorm:"column:otp_time_start"`
	OtpTimeEnd   time.Time `json:"timeEnd" gorm:"column:otp_time_end"`
	OtpUsedOk    int       `json:"usedOk" gorm:"column:otp_used_ok"`
}

//TableName --> Table for Order Model
func (Otp) TableName() string {
	return "otps"
}
