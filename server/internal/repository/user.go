package repository

import (
	"atro/internal/model"

	"github.com/jinzhu/gorm"
)

type UserRepository interface {
	GetUserByEmail(string) (model.User, error)
	GetUser(int) (model.User, error)
	AddUser(user model.User) (model.User, error)
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

func (db *userRepository) GetUser(id int) (user model.User, err error) {
	return user, db.connection.First(&user,"user_id=?", id).Error
}

func (db *userRepository) AddUser(user model.User) (model.User, error){
	return user, db.connection.Create(&user).Error
}

func (db *userRepository) GetUserByEmail(email string) (user model.User, err error){
	return user, db.connection.First(&user, "user_email=?", email).Error
}