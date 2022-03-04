package repository

import (
	"atro/internal/model"
	"github.com/jinzhu/gorm"
)

type UserRepository interface {
	GetUser(int) (model.User, error)
}

type userRepository struct {
	connection *gorm.DB
}

//NewUserRepository --> returns new user repository
func NewUserRepository() UserRepository {
	return &userRepository{
		connection: DB(),
	}
}
