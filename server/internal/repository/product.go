package repository

import (
	"github.com/jinzhu/gorm"
)

// chưa implement Tuấn đỉnh cao làm nha 
type ProductRepository interface { 
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
