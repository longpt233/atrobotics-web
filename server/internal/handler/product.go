package handler

import (
	"atro/internal/helper"
	"atro/internal/repository"
	"net/http"

	"github.com/gin-gonic/gin"
)

type ProductHandler interface {
	GetProduct(*gin.Context)
	GetAllProduct(*gin.Context)
	AddProduct(*gin.Context)
	UpdateProduct(*gin.Context)
	DeleteProduct(*gin.Context)
}

type productHandler struct {
	repo repository.ProductRepository
}

//NewProductHandler --> returns new handler for product entity
func NewProductHandler() ProductHandler {
	return &productHandler{
		repo: repository.NewProductRepository(),
	}
}

func (h *productHandler) GetAllProduct(ctx *gin.Context) {

}

func (h *productHandler) GetProduct(ctx *gin.Context) {

}

func (h *productHandler) AddProduct(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, helper.BuildResponse(1, "add product oke nha", ""))

}
func (h *productHandler) UpdateProduct(ctx *gin.Context) {

}
func (h *productHandler) DeleteProduct(ctx *gin.Context) {

}
