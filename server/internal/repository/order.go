package repository

import (
	"atro/internal/model"

	"github.com/jinzhu/gorm"
)

//OrderRepository --> Repository for Order Model
type OrderRepository interface {
	OrderProduct(model.Order) (model.Order,error)
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

func (db *orderRepository) OrderProduct(order model.Order) (model.Order,error) {
	return order, db.connection.Create(&order).Error
}
