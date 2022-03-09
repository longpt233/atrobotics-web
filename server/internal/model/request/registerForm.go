package request

type RegisterForm struct{
	Email string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
	FirstName string `json:"firstname"`
	LastName string `json:"lastname"`
	Phone string `json:"phone"`
	Address string `json:"address"`
}