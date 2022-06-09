package handler

import (
	"atro/internal/helper"
	"atro/internal/model"
	"atro/internal/model/request"
	"atro/internal/model/response"
	"atro/internal/repository"
	"fmt"
	"net/http"
	"net/url"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type ProductHandler interface {
	GetProduct(*gin.Context)
	GetAllProduct(*gin.Context)
	AddProduct(*gin.Context)
	UpdateProduct(*gin.Context)
	DeleteProduct(*gin.Context)
	GetAllProductBrand(*gin.Context)
	SearchByShortDescription(*gin.Context)
	GetProductByCategory(*gin.Context)
	GetListProductForAllCategory(*gin.Context)
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

func (h *productHandler) GetAllProduct(ctx *gin.Context) {

	// tạo query sort
	sortBy := ctx.Query("sort-by")
	if sortBy == "" {
		sortBy = "product_id.asc" // sortBy is expected to look like field.orderdirection i. e. id.asc
	}
	sortQuery, err := helper.ValidateAndReturnSortQuery(model.Product{}, sortBy)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, helper.BuildResponse(-1, "invalid param sort", err.Error()))
		return
	}

	// tao query limit
	strLimit := ctx.Query("limit")
	fmt.Println("param limit", strLimit)
	limit := -1 // with a value as -1 for gorms Limit method, we'll get a request without limit as default
	if strLimit != "" {
		limit, err = strconv.Atoi(strLimit)
		if err != nil || limit < -1 {
			ctx.JSON(http.StatusBadRequest, helper.BuildResponse(-1, "limit query parameter is no valid number", err.Error()))
			return
		}
	}

	// tạo query offset
	strOffset := ctx.Query("offset")
	offset := -1
	if strOffset != "" {
		offset, err = strconv.Atoi(strOffset)
		if err != nil || offset < -1 {
			ctx.JSON(http.StatusBadRequest, helper.BuildResponse(-1, "offset query parameter is no valid number", err.Error()))
			return
		}
		offset = offset - 1
	}

	// tạo query filter
	filter := ctx.Query("filter")
	filterMap := map[string]interface{}{}
	if filter != "" {
		filterMap, err = helper.ValidateAndReturnFilterMap(model.Product{}, filter)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, helper.BuildResponse(-1, "invalid filter param ", err.Error()))
			return
		}
	}

	q, err := url.ParseQuery(ctx.Request.URL.RawQuery)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, helper.BuildResponse(-1, "cant not parse url", err.Error()))
		return
	}
	// tao query search
	pattern := q.Get("q") // default ""

	productsCountAll, err := h.repo.GetAllProductWithOptions(filterMap, -1, -1, sortQuery, pattern)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, helper.BuildResponse(-1, "cant not count all ?? ", err.Error()))
		return
	}

	offset = offset * int(len(productsCountAll)/limit)

	// gửi query
	products, err := h.repo.GetAllProductWithOptions(filterMap, limit, offset, sortQuery, pattern)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, helper.BuildResponse(-1, "not found !", err.Error()))
		return
	}

	// trả về thành công
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

	returnData := map[string]interface{}{
		"data":  rsProducts,
		"total": len(productsCountAll),
	}

	ctx.JSON(http.StatusOK, helper.BuildResponse(1, "get list products successfully!", returnData))

}

func (h *productHandler) GetAllProductBrand(ctx *gin.Context) {
	listBrand, err := h.repo.GetAllProductBrand()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, helper.BuildResponse(0, "error when get all brand", err))
		return
	}
	var brands []string
	for i := 0; i < len(listBrand); i++ {
		brands = append(brands, listBrand[i].ProductBrand)
	}
	ctx.JSON(http.StatusOK, helper.BuildResponse(1, "get list brand successfully", brands))
}
func (h *productHandler) SearchByShortDescription(ctx *gin.Context) {

	q, err := url.ParseQuery(ctx.Request.URL.RawQuery)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, helper.BuildResponse(-1, "cant not parse url", err.Error()))
		return
	}
	// tao query search
	pattern := q.Get("q")

	listProduct, err := h.repo.SearchByShortDescription(pattern)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, helper.BuildResponse(0, "error when search", err))
		return
	}
	var rsProducts []response.ProductResponse
	for i := 0; i < len(listProduct); i++ {
		var p response.ProductResponse
		p, err := p.ProductToProductResponse(listProduct[i])
		if err != nil {
			ctx.JSON(http.StatusBadRequest, helper.BuildResponse(-1, "Cant convert json to array", err.Error()))
			return
		}
		rsProducts = append(rsProducts, p)
	}
	ctx.JSON(http.StatusOK, helper.BuildResponse(1, "search successfully", rsProducts))
}
func (h *productHandler) GetProductByCategory(ctx *gin.Context) {
	categoryId := ctx.Query("categoryId")

	category, err := repository.NewProductCategoryRepository().GetProductCategory(categoryId)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, helper.BuildResponse(0, "Category Id not exist", err))
		return
	}
	listProduct, err := h.repo.GetProductByCategory(category.ProductCategoryID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, helper.BuildResponse(0, "error when get product list by category id", err))
		return
	}
	var rsProducts []response.ProductResponse
	for i := 0; i < len(listProduct); i++ {
		var p response.ProductResponse
		p, err := p.ProductToProductResponse(listProduct[i])
		if err != nil {
			ctx.JSON(http.StatusBadRequest, helper.BuildResponse(-1, "Cant convert json to array", err.Error()))
			return
		}
		rsProducts = append(rsProducts, p)
	}
	ctx.JSON(http.StatusOK, helper.BuildResponse(1, "get list product by category successfully", rsProducts))
}

func (h *productHandler) GetListProductForAllCategory(ctx *gin.Context) {
	listCategory, err := repository.NewProductCategoryRepository().GetAllProductCategories()

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, helper.BuildResponse(0, "error when get list category", err))
		return
	}
	var responseList response.ListProductAllCategoryResponse

	for i := 0; i < len(listCategory); i++ {
		responseList.CategoryName = listCategory[i].CategoryName

		listProduct, err := h.repo.GetProductByCategory(listCategory[i].ProductCategoryID)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, helper.BuildResponse(0, "error when get product list by category", err))
			return
		}
		for j := 0; j < len(listProduct); j++ {
			var p response.ProductResponse
			p, err := p.ProductToProductResponse(listProduct[j])
			if err != nil {
				ctx.JSON(http.StatusBadRequest, helper.BuildResponse(-1, "Cant convert json to array", err.Error()))
				return
			}
			responseList.ProductList = append(responseList.ProductList, p)
		}
	}
	ctx.JSON(http.StatusOK, helper.BuildResponse(1, "get list product for all category", responseList))
	return
}
