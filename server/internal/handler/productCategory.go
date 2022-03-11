package handler

import (
	"atro/internal/helper"
	"atro/internal/model"
	"atro/internal/repository"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type ProductCategoryHandler interface {
	GetProductCategory(* gin.Context)
	GetAllProductCategories(* gin.Context)
	AddProductCategory(* gin.Context)
	UpdateProductCategory(* gin.Context)
	DeleteProductCategory(* gin.Context)
}

type productCategoryHandler struct{
	repo repository.ProductCategoryRepository
}

func NewProductCategoryHandler() ProductCategoryHandler{
	return &productCategoryHandler{
		repo: repository.NewProductCategoryRepository(),
	}
}

func (h *productCategoryHandler) GetProductCategory(ctx *gin.Context){
	id := ctx.Param("id")
	intID, err := strconv.Atoi(id)
	if err!= nil {
		ctx.JSON(http.StatusBadRequest, helper.BuildResponse(-1, "invalid id input", err.Error()))
		return
	}
	category, err := h.repo.GetProductCategory(intID)
	if err!= nil {
		ctx.JSON(http.StatusInternalServerError, helper.BuildResponse(-1, "error when find category", err.Error()))
		return
	}
	ctx.JSON(http.StatusOK,helper.BuildResponse(1,"get product category successfully!", category))

}
func (h *productCategoryHandler) GetAllProductCategories(ctx *gin.Context){
	categories, err := h.repo.GetAllProductCategories()
	if err!= nil {
		ctx.JSON(http.StatusInternalServerError, helper.BuildResponse(-1, "error when get list categories", err.Error()))
		return 
	}
	ctx.JSON(http.StatusOK, helper.BuildResponse(1,"get list categories successfully!", categories))
}
func (h *productCategoryHandler) UpdateProductCategory(ctx *gin.Context){
	id := ctx.Param("id")
	intID, err := strconv.Atoi(id)
	if err!= nil {
		ctx.JSON(http.StatusBadRequest, helper.BuildResponse(-1, "invalid id input", err.Error()))
		return
	}
	var category model.ProductCategory
	if err:= ctx.ShouldBindJSON(&category); err != nil {
		ctx.JSON(http.StatusBadRequest, helper.BuildResponse(-1,"invalid input format", err.Error()))
		return
	}
	category.ProductCategoryID = intID
	newCategory, err := h.repo.UpdateProductCategory(category)
	if err!= nil {
		ctx.JSON(http.StatusInternalServerError, helper.BuildResponse(-1, "error when find category", err.Error()))
		return
	}
	ctx.JSON(http.StatusOK, helper.BuildResponse(1,"update category successfully!",newCategory))

}
func (h *productCategoryHandler) DeleteProductCategory(ctx *gin.Context){
	id := ctx.Param("id")
	intID, err := strconv.Atoi(id)
	if err!= nil {
		ctx.JSON(http.StatusBadRequest, helper.BuildResponse(-1, "invalid id input", err.Error()))
		return
	}
	category, err := h.repo.DeleteProductCategory(intID)
	if err!= nil {
		ctx.JSON(http.StatusInternalServerError, helper.BuildResponse(-1, "error when find category", err.Error()))
		return
	}
	ctx.JSON(http.StatusOK, helper.BuildResponse(1,"delete category successfully!",category))
}
func (h *productCategoryHandler) AddProductCategory(ctx *gin.Context){
	var category model.ProductCategory
	if err:= ctx.ShouldBindJSON(&category); err != nil {
		ctx.JSON(http.StatusBadRequest, helper.BuildResponse(-1,"invalid input format", err.Error()))
		return
	}
	newCategory, err := h.repo.AddProductCategory(category)
	if err!= nil {
		ctx.JSON(http.StatusInternalServerError, helper.BuildResponse(-1, "error when add new category", err.Error()))
		return
	}
	ctx.JSON(http.StatusOK, helper.BuildResponse(1, "add product category successfully!", newCategory))
}