package repository

import (
	"atro/internal/model"

	"github.com/jinzhu/gorm"
)

type ProductRepository interface {
	AddProduct(model.Product) (model.Product, error)
	GetProduct(string) (model.Product, error)
	CountProduct() (int, error)
	GetAllProductWithOptions(filter map[string]interface{}, limit int, offset int, query string, searchPattern string) ([]model.Product, error)
	UpdateProduct(model.Product) (model.Product, error)
	DeleteProduct(string) (model.Product, error)
	GetAllProductBrand() ([]model.Product, error)
	SearchByShortDescription(pattern string) ([]model.Product, error)
	GetProductByCategory(string) ([]model.Product, error)
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

func (db *productRepository) CountProduct() (count int, err error) {
	product := model.Product{}
	return count, db.connection.Model(&product).Count(&count).Error
}

func (db *productRepository) GetAllProductWithOptions(filter map[string]interface{}, limit int, offset int, order string, searchPattern string) (products []model.Product, err error) {
	return products, db.connection.Where(filter).Where("product_short_desc LIKE ?", "%"+searchPattern+"%").Limit(limit).Offset(offset).Order(order).Find(&products).Error
}

func (db *productRepository) GetAllProductBrand() ([]model.Product, error) {
	var listBrand []model.Product
	return listBrand, db.connection.Raw("SELECT DISTINCT product_brand FROM products").Scan(&listBrand).Error
}

func (db *productRepository) SearchByShortDescription(pattern string) ([]model.Product, error) {
	var listProduct []model.Product

	return listProduct, db.connection.Raw("SELECT * FROM products WHERE product_short_desc LIKE '%" + pattern + "%'").Scan(&listProduct).Error
}
func (db *productRepository) GetProductByCategory(categoryId string) (listProduct []model.Product, err error) {
	return listProduct, db.connection.Find(&listProduct, "product_category_id=?", categoryId).Limit(4).Error
}
