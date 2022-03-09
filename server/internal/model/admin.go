package model

//User -> model for users table
type Admin struct {
	ID       uint   `gorm:"primary_key"`
	Username string `json:"" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func (Admin) TableName() string {
	return "admin"
}
