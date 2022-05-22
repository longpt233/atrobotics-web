package repository

import (
	"atro/internal/model"

	"github.com/jinzhu/gorm"
)

type ProductRepository interface {
	AddProduct(model.Product) (model.Product, error)
	GetProduct(string) (model.Product, error)
	GetAllProductWithOptions(filter map[string]interface{}, limit int, offset int, query string) ([]model.Product, error)
	UpdateProduct(model.Product) (model.Product, error)
	DeleteProduct(string) (model.Product, error)
}

type productRepository struct {
	connection *gorm.DB
}

//NewProductRepository --> returns new product repository
func NewProductRepository() ProductRepository {

	myclient := &MySQLClient{}
	return &productRepository{
		connection:myclient.GetConn(),
	}
}

func (db *productRepository) GetProduct(id string) (product model.Product, err error) {
	return product, db.connection.First(&product, "product_id=?", id).Error
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

func (db *productRepository) GetAllProductWithOptions(filter map[string]interface{}, limit int, offset int, query string) (products []model.Product, err error) {
	return products, db.connection.Where(filter).Limit(limit).Offset(offset).Order(query).Find(&products).Error
}
