package repository

import (
	"atro/internal/model"

	"github.com/jinzhu/gorm"
)

type UserRepository interface {
	GetUserByEmail(string) (model.User, error)
	GetUser(string) (model.User, error)
	AddUser(user model.User) (model.User, error)
	UpdateUser(user model.User) (model.User, error)
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

func (db *userRepository) GetUser(id string) (user model.User, err error) { // TODO tại sao k xóa dc cái user thường ở cái return này đi như hàm bên dưới ?
	return user, db.connection.First(&user, "user_id=?", id).Error
}

func (db *userRepository) AddUser(user model.User) (model.User, error) {
	return user, db.connection.Create(&user).Error
}

func (db *userRepository) GetUserByEmail(email string) (user model.User, err error) {
	return user, db.connection.First(&user, "user_email=?", email).Error
}
func (db *userRepository) UpdateUser(user model.User) (model.User, error){
	var checkUser model.User
	if err := db.connection.First(&checkUser, "user_id = ?",user.UserID).Error; err != nil {
		return checkUser, err
	}
	return user, db.connection.Model(&user).Where(model.User{UserID: user.UserID}).Updates(&user).Error
}
