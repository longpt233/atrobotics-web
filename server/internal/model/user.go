package model

//User -> model for users table
type User struct {
	UserID        int    `json:"user_id" gorm:"primaryKey"`
	UserRoleID    int    `json:"user_role_id"` // col name
	UserEmail     string `json:"user_email" binding:"required,email" gorm:"unique"`
	UserPassword  string `json:"user_password" binding:"required"`
	UserFirstName string `json:"user_firstname"`
	UserLastName  string `json:"user_lastname"`
	UserPhone     string `json:"user_phone"`
	UserAddress   string `json:"user_address"`
}

//TableName --> Table for Product Model
func (User) TableName() string {
	return "users"
}
