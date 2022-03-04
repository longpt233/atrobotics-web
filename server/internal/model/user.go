package model

import "github.com/jinzhu/gorm"

//User -> model for users table
type User struct {
	gorm.Model
	Id       int    `json:"id" binding:"required"`   
	Name     string `json:"name" binding:"required"` // col name
	Email    string `json:"email" binding:"required,email" gorm:"unique"`
	Password string `json:"password" binding:"required"`
}

//TableName --> Table for Product Model
func (User) TableName() string {
	return "user"
}
