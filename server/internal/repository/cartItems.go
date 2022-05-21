package repository

import (
	"atro/internal/model"

	"github.com/jinzhu/gorm"
)

type CartItemsRepository interface {
	GetCartItemsByUserId(string) ([]model.CardItems, error)
	AddCartItems(model.CardItems) (model.CardItems, error)
	DeleteCartItems(string) (model.CardItems, error)
	UpdateCartItems(model.CardItems) (model.CardItems, error)
	GetCartItemsByUserIdAndProductId(string, string) (model.CardItems, error)
}

type cartItemsRepository struct {
	connection *gorm.DB
}

func NewCartItemsRepository() CartItemsRepository {
	return &cartItemsRepository{
		connection: DB(),
	}
}

func (db *cartItemsRepository) GetCartItemsByUserId(userId string) (listCartItems []model.CardItems, err error) {

	// var  cartItem model.CardItems
	return listCartItems, db.connection.Model(listCartItems).
		Where("cart_user_id=?", userId).
		Preload("CardItems").
		Find(&listCartItems).Error
}

func (db *cartItemsRepository) AddCartItems(cartItems model.CardItems) (model.CardItems, error) {
	return cartItems, db.connection.Create(&cartItems).Error
}

func (db *cartItemsRepository) DeleteCartItems(cartId string) (model.CardItems, error) {
	var cartItems model.CardItems
	if err := db.connection.First(&cartItems, "cart_id=?", cartId).Error; err != nil {
		return cartItems, err
	}

	return cartItems, db.connection.Delete(&cartItems, "cart_id=?", cartId).Error
}

func (db *cartItemsRepository) UpdateCartItems(modifyCartItems model.CardItems) (model.CardItems, error) {
	var checkCartItems model.CardItems
	if err := db.connection.First(&checkCartItems, "cart_id=?", modifyCartItems.CartId).Error; err != nil {
		return checkCartItems, err
	}
	modifyCartItems.CartCreatedAt = checkCartItems.CartCreatedAt
	return modifyCartItems, db.connection.Model(&modifyCartItems).Where(model.CardItems{CartId: modifyCartItems.CartId}).Updates(&modifyCartItems).Error
}

func (db *cartItemsRepository) GetCartItemsByUserIdAndProductId(userId string, productId string) (cartItem model.CardItems, err error) {
	return cartItem, db.connection.Find(&cartItem, "cart_user_id=? and cart_product_id=?", userId, productId).Error
}
