package repository

import (
	"atro/internal/model"

	"github.com/jinzhu/gorm"
)

//OrderRepository --> Repository for Order Model
type OrderRepository interface {
	OrderProduct(string, string, int) error
}

type orderRepository struct {
	connection *gorm.DB
}

//NewOrderRepository --> returns new order repository
func NewOrderRepository() OrderRepository {
	return &orderRepository{
		connection: DB(),
	}
}

func (db *orderRepository) OrderProduct(userID string, productID string, quantity int) error {
	return db.connection.Create(&model.Order{
		ProductID: productID,
		UserID:    userID,
		Quantity:  quantity,
	}).Error
}
