package model

//User -> model for users table
type User struct {
	UserID        string `json:"id" gorm:"primaryKey" gorm:"column:user_id"`
	UserRoleID    string `json:"roleId" gorm:"column:user_role_id"`                                      // col name
	UserEmail     string `json:"email" binding:"required,email" gorm:"unique" gorm:"column:user_email"` //TODO:  binding ?
	UserPassword  string `json:"password" binding:"required" gorm:"column:user_password"`
	UserFirstName string `json:"firstName" gorm:"column:user_first_name"`
	UserLastName  string `json:"lastName" gorm:"column:user_last_name"`
	UserPhone     string `json:"phone" gorm:"column:user_phone"`
	UserAddress   string `json:"address" gorm:"column:user_address"`
}

//TableName --> Table for Product Model
func (User) TableName() string {
	return "users"
}
