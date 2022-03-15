package handler

import (
	"atro/internal/helper"
	"atro/internal/model/request"
	"atro/internal/model/response"
	"atro/internal/repository"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
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
	var rsProducts []response.ProductResponse
	for i := 0; i < len(products); i++ {
		var p response.ProductResponse
		p, err := p.ProductToProductResponse(products[i])
		if err != nil {
			ctx.JSON(http.StatusBadRequest, helper.BuildResponse(-1, "Cant convert json to array", err.Error()))
			return
		}
		rsProducts = append(rsProducts, p)
	}
	
	ctx.JSON(http.StatusOK, helper.BuildResponse(1, "get list products successfully!", rsProducts))
}

func (h *productHandler) GetProduct(ctx *gin.Context) {
	id := ctx.Param("id")
	product, err := h.repo.GetProduct(id)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, helper.BuildResponse(-1, "error when find product", err.Error()))
		return
	}
	var resProduct response.ProductResponse
	resProduct, err = resProduct.ProductToProductResponse(product)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, helper.BuildResponse(-1, "Cant convert json to array", err.Error()))
		return
	}
	ctx.JSON(http.StatusOK, helper.BuildResponse(1, "get product successfully!", resProduct))
}

func (h *productHandler) AddProduct(ctx *gin.Context) {
	var newProduct request.ProductRequest
	if err := ctx.ShouldBindJSON(&newProduct); err != nil {
		ctx.JSON(http.StatusBadRequest, helper.BuildResponse(-1, "invalid id input", err.Error()))
		return
	}
	rsProduct, err := newProduct.ProductRequestToProduct()
	
	if err != nil {
		ctx.JSON(http.StatusBadRequest, helper.BuildResponse(-1, "Cant convert array to json", err.Error()))
		return
	}
	rsProduct.ProductID = uuid.NewString()
	product, err := h.repo.AddProduct(rsProduct)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, helper.BuildResponse(-1, "error when add product", err.Error()))
		return
	}
	ctx.JSON(http.StatusOK, helper.BuildResponse(1, "add product successfully!", product))
}
func (h *productHandler) UpdateProduct(ctx *gin.Context) {
	var newProduct request.ProductRequest
	if err := ctx.ShouldBindJSON(&newProduct); err != nil {
		ctx.JSON(http.StatusBadRequest, helper.BuildResponse(-1, "invalid id input", err.Error()))
		return
	}
	id := ctx.Param("id")
	rsProduct, err := newProduct.ProductRequestToProduct()
	if err != nil {
		ctx.JSON(http.StatusBadRequest, helper.BuildResponse(-1, "Cant convert array to json", err.Error()))
		return
	}
	rsProduct.ProductID = id
	updateProduct, err := h.repo.UpdateProduct(rsProduct)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, helper.BuildResponse(-1, "error when find product", err.Error()))
		return
	}
	ctx.JSON(http.StatusOK, helper.BuildResponse(1, "update product successfully!", updateProduct))
}
func (h *productHandler) DeleteProduct(ctx *gin.Context) {
	id := ctx.Param("id")
	product, err := h.repo.DeleteProduct(id)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, helper.BuildResponse(-1, "error when find product", err.Error()))
		return
	}
	ctx.JSON(http.StatusOK, helper.BuildResponse(1, "delete product successfully!", product))

}
