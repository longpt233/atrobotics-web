package handler

import (
	"atro/internal/helper"
	"atro/internal/model"
	"atro/internal/model/base"
	"atro/internal/model/request"
	"atro/internal/model/response"
	"atro/internal/repository"
	"encoding/json"
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
	var productImg []string
	var productCol []string
	if err := json.Unmarshal([]byte(product.ProductImages), &productImg); err != nil {
		ctx.JSON(http.StatusBadRequest, helper.BuildResponse(-1, "Cant convert json to []string", ""))
		return
	}
	if err := json.Unmarshal([]byte(product.ProductColor), &productCol); err != nil {
		ctx.JSON(http.StatusBadRequest, helper.BuildResponse(-1, "Cant convert json to []string", ""))
		return
	}
	resProduct := response.ProductResponse{
		ProductID:        intID,
		ProductImages:    productImg,
		ProductColor:     productCol,
		ProductUpdatedAt: product.ProductUpdatedAt,
		ProductCreatedAt: product.ProductCreatedAt,
		BaseProduct: base.BaseProduct{
			ProductName:       product.ProductName,
			ProductPrice:      product.ProductPrice,
			ProductShortDesc:  product.ProductShortDesc,
			ProductLongDesc:   product.ProductLongDesc,
			ProductCategoryID: product.ProductCategoryID,
			ProductAvailable:  product.ProductAvailable,
			ProductBrand:      product.ProductBrand,
			ProductSold:       product.ProductSold,
		},
	}
	ctx.JSON(http.StatusOK, helper.BuildResponse(1, "get product successfully!", resProduct))
}

func (h *productHandler) AddProduct(ctx *gin.Context) {
	var newProduct request.NewProductForm
	if err := ctx.ShouldBindJSON(&newProduct); err != nil {
		ctx.JSON(http.StatusBadRequest, helper.BuildResponse(-1, "invalid id input", err.Error()))
		return
	}
	productImg, err := json.Marshal(newProduct.ProductImages)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, helper.BuildResponse(-1, "Cant convert array to json", ""))
		return
	}
	productCol, err := json.Marshal(newProduct.ProductColor)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, helper.BuildResponse(-1, "Cant convert array to json", ""))
		return
	}
	rsProduct := model.Product{
		ProductImages:    string(productImg),
		ProductColor:     string(productCol),
		ProductCreatedAt: time.Now(),
		ProductUpdatedAt: time.Now(),
		BaseProduct: base.BaseProduct{
			ProductName:       newProduct.ProductName,
			ProductPrice:      newProduct.ProductPrice,
			ProductShortDesc:  newProduct.ProductShortDesc,
			ProductLongDesc:   newProduct.ProductLongDesc,
			ProductCategoryID: newProduct.ProductCategoryID,
			ProductAvailable:  newProduct.ProductAvailable,
			ProductSold:       newProduct.ProductSold,
			ProductBrand:      newProduct.ProductBrand,
		},
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
	productImg, err := json.Marshal(newProduct.ProductImages)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, helper.BuildResponse(-1, "Cant convert array to json", ""))
		return
	}
	productCol, err := json.Marshal(newProduct.ProductColor)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, helper.BuildResponse(-1, "Cant convert array to json", ""))
		return
	}
	rsProduct := model.Product{
		ProductID:        intID,
		ProductImages:    string(productImg),
		ProductColor:     string(productCol),
		ProductCreatedAt: time.Now(),
		ProductUpdatedAt: time.Now(),
		BaseProduct: base.BaseProduct{
			ProductName:       newProduct.ProductName,
			ProductPrice:      newProduct.ProductPrice,
			ProductShortDesc:  newProduct.ProductShortDesc,
			ProductLongDesc:   newProduct.ProductLongDesc,
			ProductCategoryID: newProduct.ProductCategoryID,
			ProductAvailable:  newProduct.ProductAvailable,
			ProductSold:       newProduct.ProductSold,
			ProductBrand:      newProduct.ProductBrand,
		},
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
