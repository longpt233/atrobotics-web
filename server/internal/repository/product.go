package repository

import (
	"atro/internal/model"

	"github.com/jinzhu/gorm"
)

type ProductRepository interface {
	AddProduct(model.Product) (model.Product, error)
	GetProduct(string) (model.Product, error)
	GetAllProducts() ([]model.Product, error)
	UpdateProduct(model.Product) (model.Product, error)
	DeleteProduct(string) (model.Product, error)
}

type productRepository struct {
	connection *gorm.DB
}

//NewProductRepository --> returns new product repository
func NewProductRepository() ProductRepository {
	return &productRepository{
		connection: DB(),
	}
}

func (db *productRepository) GetProduct(id string) (product model.Product, err error) {
	return product, db.connection.First(&product, "product_id=?", id).Error
}
func (db *productRepository) GetAllProducts() (products []model.Product, err error) {
	return products, db.connection.Find(&products).Error
}
func (db *productRepository) AddProduct(product model.Product) (model.Product, error) {
	return product, db.connection.Create(&product).Error
}
func (db *productRepository) UpdateProduct(product model.Product) (model.Product, error) {
	var checkProduct model.Product
	if err := db.connection.First(&checkProduct, "product_id=?", product.ProductID).Error; err != nil {
		return checkProduct, err
	}
	product.ProductCreatedAt = checkProduct.ProductCreatedAt
	return product, db.connection.Model(&product).Where(model.Product{ProductID: product.ProductID}).Updates(&product).Error
}
func (db *productRepository) DeleteProduct(id string) (model.Product, error) {
	var product model.Product
	if err := db.connection.First(&product, "product_id=?", id).Error; err != nil {
		return product, err
	}
	return product, db.connection.Delete(&product, "product_id=?", product.ProductID).Error
}
