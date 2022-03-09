package repository

import (
	"atro/internal/model"

	"github.com/jinzhu/gorm"
)

type AdminRepository interface {
	GetByUserName(string) (model.Admin, error)
	AddUser(admin model.Admin) (model.Admin, error)
}

type adminRepository struct {
	connection *gorm.DB
}

func NewAdminRepository() AdminRepository {
	return &adminRepository{
		connection: DB(),
	}
}

func (db *adminRepository) GetByUserName(name string) (admin model.Admin, err error) {   // TODO tại sao k xóa dc cái admin thường ở cái return này đi như hàm bên dưới ? 
	return admin, db.connection.First(&admin, "username=?", name).Error
}

func (db *adminRepository) AddUser(admin model.Admin) (model.Admin,error) {
	return admin, db.connection.Create(&admin).Error
}
