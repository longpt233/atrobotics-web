package repository

import (
	"atro/internal/model"

	"github.com/jinzhu/gorm"
)

type DeliveryAddressRepository interface {
	AddDeliveryAddress(model.DeliveryAddress) (model.DeliveryAddress, error)
	GetDeliveryAddressByUser(string) ([]model.DeliveryAddress, error)
	DeleteDeliveryAddress(string) (model.DeliveryAddress, error)
	UpdateDeliveryAddress(model.DeliveryAddress) (model.DeliveryAddress, error)
	GetDeliveryAddressById(string) (model.DeliveryAddress, error)
}

type deliveryAddressRepository struct {
	connection *gorm.DB
}

func NewDeliveryAddressRepository() DeliveryAddressRepository{
	 
	myclient := &MySQLClient{}
	return &deliveryAddressRepository{
		connection:myclient.GetConn(),
	}
}

func (db *deliveryAddressRepository) AddDeliveryAddress(newAddress model.DeliveryAddress) (model.DeliveryAddress, error){
	return newAddress, db.connection.Create(&newAddress).Error
}
func (db *deliveryAddressRepository) GetDeliveryAddressByUser(userId string) (listAddress []model.DeliveryAddress, err error){
	return listAddress, db.connection.Find(&listAddress,"user_id=?",userId).Error
}
func (db *deliveryAddressRepository) DeleteDeliveryAddress(addressId string) (model.DeliveryAddress, error){
	var address model.DeliveryAddress
	if err := db.connection.First(&address,"delivery_address_id=?",addressId).Error; err != nil {
		return address, err
	}

	return address, db.connection.Delete(&address,"delivery_address_id=?",addressId).Error
}
func (db *deliveryAddressRepository) UpdateDeliveryAddress(modifyAddress model.DeliveryAddress) (model.DeliveryAddress, error){
	var checkAddress model.DeliveryAddress
	if err := db.connection.First(&checkAddress,"delivery_address_id=?", modifyAddress.DeliveryAddressId).Error; err != nil {
		return checkAddress, err
	}
	modifyAddress.CreatedAt = checkAddress.CreatedAt
	return modifyAddress, db.connection.Model(&modifyAddress).Where(model.DeliveryAddress{DeliveryAddressId: modifyAddress.DeliveryAddressId}).Updates(&modifyAddress).Error
}
func (db *deliveryAddressRepository) GetDeliveryAddressById(id string) (address model.DeliveryAddress, err error){
	return address, db.connection.First(&address, "delivery_address_id=?",id).Error
}