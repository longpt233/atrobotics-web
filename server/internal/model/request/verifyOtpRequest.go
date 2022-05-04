package request

type VerifyOtpRequest struct {
	UserId   string `json:"userId" binding:"required"`
	OtpValue string `json:"otpValue" binding:"required"`
}
