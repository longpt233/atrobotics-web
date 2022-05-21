package repository

import (
	"atro/internal/model"

	"github.com/jinzhu/gorm"
)

type ProductCategoryRepository interface {
	AddProductCategory(model.ProductCategory) (model.ProductCategory, error)
	GetProductCategory(string) (model.ProductCategory, error)
	GetAllProductCategories() ([]model.ProductCategory, error)
	DeleteProductCategory(string) (model.ProductCategory, error)
	UpdateProductCategory(model.ProductCategory) (model.ProductCategory, error)
	GetProductCategoryByName(string) (model.ProductCategory, error)
}

type productCategoryRepository struct {
	connection *gorm.DB
}

func NewProductCategoryRepository() ProductCategoryRepository {
	return &productCategoryRepository{
		connection: DB(),
	}
}

func (db *productCategoryRepository) GetProductCategory(id string) (productCategory model.ProductCategory, err error) {
	return productCategory, db.connection.First(&productCategory, "product_category_id=?", id).Error
}

func (db *productCategoryRepository) GetAllProductCategories() (productCategories []model.ProductCategory, err error) {
	return productCategories, db.connection.Find(&productCategories).Error
}

func (db *productCategoryRepository) AddProductCategory(productCategory model.ProductCategory) (model.ProductCategory, error) {
	return productCategory, db.connection.Create(&productCategory).Error
}

func (db *productCategoryRepository) DeleteProductCategory(id string) (model.ProductCategory, error) {
	var category model.ProductCategory
	if err := db.connection.First(&category, "product_category_id=?", id).Error; err != nil {
		return category, err
	}
	return category, db.connection.Delete(&category, "product_category_id=?", category.ProductCategoryID).Error
}

func (db *productCategoryRepository) UpdateProductCategory(category model.ProductCategory) (model.ProductCategory, error) {
	var checkCategory model.ProductCategory
	if err := db.connection.First(&checkCategory, "product_category_id=?", category.ProductCategoryID).Error; err != nil {
		return checkCategory, err
	}
	return category, db.connection.Model(&category).Where(model.ProductCategory{ProductCategoryID: category.ProductCategoryID}).Updates(&category).Error
}
func (db *productCategoryRepository) GetProductCategoryByName(categoryName string) (category model.ProductCategory, err error){
	return category, db.connection.First(&category,"category_name=?",categoryName).Error
}
