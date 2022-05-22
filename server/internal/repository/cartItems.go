package repository

import (
	"atro/internal/model"

	"github.com/jinzhu/gorm"
)

type CartItemsRepository interface {
	GetCartItemsByUserId(string) ([]model.CartItems, error)
	AddCartItems(model.CartItems) (model.CartItems, error)
	DeleteCartItems(string) (model.CartItems, error)
	UpdateCartItems(model.CartItems) (model.CartItems, error)
	GetCartItemsByUserIdAndProductId(string, string) (model.CartItems, error)
}

type cartItemsRepository struct {
	connection *gorm.DB
}

func NewCartItemsRepository() CartItemsRepository {

	myclient := &MySQLClient{}
	return &cartItemsRepository{
		connection: myclient.GetConn(),
	}
}

func (db *cartItemsRepository) GetCartItemsByUserId(userId string) (listCartItems []model.CartItems, err error) {

	// var  cartItem model.CartItems
	return listCartItems, db.connection.
		Select("cart_items.*, products.*").
		Where("cart_user_id=?", userId).
		Joins("join products as products on products.product_id = cart_items.cart_product_id").
		// Preload("CartItems").
		Find(&listCartItems).Error
}

func (db *cartItemsRepository) AddCartItems(cartItems model.CartItems) (model.CartItems, error) {
	return cartItems, db.connection.Create(&cartItems).Error
}

func (db *cartItemsRepository) DeleteCartItems(cartId string) (model.CartItems, error) {
	var cartItems model.CartItems
	if err := db.connection.First(&cartItems, "cart_id=?", cartId).Error; err != nil {
		return cartItems, err
	}

	return cartItems, db.connection.Delete(&cartItems, "cart_id=?", cartId).Error
}

func (db *cartItemsRepository) UpdateCartItems(modifyCartItems model.CartItems) (model.CartItems, error) {
	var checkCartItems model.CartItems
	if err := db.connection.First(&checkCartItems, "cart_id=?", modifyCartItems.CartId).Error; err != nil {
		return checkCartItems, err
	}
	modifyCartItems.CartCreatedAt = checkCartItems.CartCreatedAt
	return modifyCartItems, db.connection.Model(&modifyCartItems).Where(model.CartItems{CartId: modifyCartItems.CartId}).Updates(&modifyCartItems).Error
}

func (db *cartItemsRepository) GetCartItemsByUserIdAndProductId(userId string, productId string) (cartItem model.CartItems, err error) {
	return cartItem, db.connection.Find(&cartItem, "cart_user_id=? and cart_product_id=?", userId, productId).Error
}
