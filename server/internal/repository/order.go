package repository

import (
	"atro/internal/model"

	"github.com/jinzhu/gorm"
)

//OrderRepository --> Repository for Order Model
type OrderRepository interface {
	OrderProduct(model.Order) (model.Order, error)
	GetAllOrder() ([]model.Order, error)
	GetOrderById(string, string) (model.Order, error)
	UpdateOrderStatus(string, int) (model.Order, error)
	GetAllOrderOptions(userId string, limit int, offset int, sortBy string) ([]model.Order, error)
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

func (db *orderRepository) OrderProduct(order model.Order) (model.Order, error) {
	return order, db.connection.Create(&order).Error
}

func (db *orderRepository) GetAllOrder() (orders []model.Order, err error) {
	return orders, db.connection.Find(&orders).Error
}

func (db *orderRepository) GetOrderById(userId string, orderId string) (order model.Order, err error) {
	return order, db.connection.Find(&order, "user_id=? and order_id=?", userId, orderId).Error
}

func (db *orderRepository) UpdateOrderStatus(orderId string, status int) (model.Order, error) {
	var checkOrder model.Order
	if err := db.connection.First(&checkOrder,"order_id=?",orderId).Error; err != nil {
		return checkOrder, err
	}
	return checkOrder, db.connection.Model(&checkOrder).Where("order_id=?",orderId).Update("order_status",status).Error
}

func (db *orderRepository) GetAllOrderOptions(userId string, limit int, offset int, sortBy string) (orders []model.Order, err error) {
	return orders, db.connection.Where("user_id=?",userId).Limit(limit).Offset(offset).Order("order_created_at "+sortBy).Find(&orders).Error
}
