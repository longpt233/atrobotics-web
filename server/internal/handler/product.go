package handler

import (
	"atro/internal/helper"
	"atro/internal/model"
	"atro/internal/model/request"
	"atro/internal/repository"
	"net/http"
	"strconv"
	"time"

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
	products, err := h.repo.GetAllProducts()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, helper.BuildResponse(-1, "error when get list products", err.Error()))
		return
	}
	ctx.JSON(http.StatusOK, helper.BuildResponse(1, "get list products successfully!", products))
}

func (h *productHandler) GetProduct(ctx *gin.Context) {
	id := ctx.Param("id")
	intID, err := strconv.Atoi(id)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, helper.BuildResponse(-1, "invalid id input", err.Error()))
		return
	}
	product, err := h.repo.GetProduct(intID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, helper.BuildResponse(-1, "error when find product", err.Error()))
		return
	}
	ctx.JSON(http.StatusOK, helper.BuildResponse(1, "get product successfully!", product))
}

func (h *productHandler) AddProduct(ctx *gin.Context) {
	var newProduct request.NewProductForm
	if err := ctx.ShouldBindJSON(&newProduct); err != nil {
		ctx.JSON(http.StatusBadRequest, helper.BuildResponse(-1, "invalid id input", err.Error()))
		return
	}
	rsProduct := model.Product{
		ProductName:       newProduct.ProductName,
		ProductPrice:      newProduct.ProductPrice,
		ProductShortDesc:  newProduct.ProductShortDesc,
		ProductLongDesc:   newProduct.ProductLongDesc,
		ProductImages:     newProduct.ProductImages,
		ProductCategoryID: newProduct.ProductCategoryID,
		ProductAvailable:  newProduct.ProductAvailable,
		ProductColor:      newProduct.ProductColor,
		ProductBrand:      newProduct.ProductBrand,
		ProductSold:       newProduct.ProductSold,
		ProductCreatedAt:  time.Now(),
		ProductUpdatedAt:  time.Now(),
	}
	product, err := h.repo.AddProduct(rsProduct)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, helper.BuildResponse(-1, "error when add product", err.Error()))
		return
	}
	ctx.JSON(http.StatusOK, helper.BuildResponse(1, "add product successfully!", product))
}
func (h *productHandler) UpdateProduct(ctx *gin.Context) {
	var newProduct request.NewProductForm
	if err := ctx.ShouldBindJSON(&newProduct); err != nil {
		ctx.JSON(http.StatusBadRequest, helper.BuildResponse(-1, "invalid id input", err.Error()))
		return
	}
	id := ctx.Param("id")
	intID, err := strconv.Atoi(id)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, helper.BuildResponse(-1, "invalid id input", err.Error()))
		return
	}
	rsProduct := model.Product{
		ProductID:         intID,
		ProductName:       newProduct.ProductName,
		ProductPrice:      newProduct.ProductPrice,
		ProductShortDesc:  newProduct.ProductShortDesc,
		ProductLongDesc:   newProduct.ProductLongDesc,
		ProductImages:     newProduct.ProductImages,
		ProductCategoryID: newProduct.ProductCategoryID,
		ProductAvailable:  newProduct.ProductAvailable,
		ProductColor:      newProduct.ProductColor,
		ProductBrand:      newProduct.ProductBrand,
		ProductSold:       newProduct.ProductSold,
		ProductUpdatedAt:  time.Now(),
	}
	updateProduct, err := h.repo.UpdateProduct(rsProduct)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, helper.BuildResponse(-1, "error when find product", err.Error()))
		return
	}
	ctx.JSON(http.StatusOK, helper.BuildResponse(1, "update product successfully!", updateProduct))
}
func (h *productHandler) DeleteProduct(ctx *gin.Context) {
	id := ctx.Param("id")
	intID, err := strconv.Atoi(id)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, helper.BuildResponse(-1, "invalid id input", err.Error()))
		return
	}
	product, err := h.repo.DeleteProduct(intID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, helper.BuildResponse(-1, "error when find product", err.Error()))
		return
	}
	ctx.JSON(http.StatusOK, helper.BuildResponse(1, "delete product successfully!", product))

}
