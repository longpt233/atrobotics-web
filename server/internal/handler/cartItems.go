package handler

import (
	"atro/internal/helper"
	"atro/internal/model"
	"atro/internal/model/request"
	"atro/internal/model/response"
	"atro/internal/repository"
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type CartItemsHandler interface {
	GetCartItemsByUserId(*gin.Context)
	AddCartItems(*gin.Context)
	UpdateCartItems(*gin.Context)
	DeleteCartItems(*gin.Context)
}

type cartItemsHandler struct {
	repo repository.CartItemsRepository
}

func NewCartItemsHandler() CartItemsHandler {
	return &cartItemsHandler{
		repo: repository.NewCartItemsRepository(),
	}
}

func (h *cartItemsHandler) AddCartItems(ctx *gin.Context) {
	var newCartItems request.CartItemsRequest
	if err := ctx.ShouldBindJSON(&newCartItems); err != nil {
		ctx.JSON(http.StatusBadRequest, helper.BuildResponse(-1, "invalid request body", err.Error()))
		return
	}
	if(newCartItems.CartQuantity <= 0){
		ctx.JSON(http.StatusBadRequest, helper.BuildResponse(-1, "Số lượng sản phẩm không hợp lệ", ""))
		return
	}
	userId, isExist := ctx.Get("userID")
	if !isExist {
		ctx.JSON(http.StatusBadRequest, helper.BuildResponse(-1, "unauthorized", "Invalid token"))
		return
	}
	strUserId := fmt.Sprint(userId)
	_, errProduct := repository.NewProductRepository().GetProduct(newCartItems.CartProductId)
	if errProduct != nil {
		ctx.JSON(http.StatusBadRequest, helper.BuildResponse(-1, "product id not found", errProduct.Error()))
		return
	}
	checkExistCart, err := repository.NewCartItemsRepository().GetCartItemsByUserIdAndProductId(strUserId, newCartItems.CartProductId)
	var rsCartItems model.CartItems
	if err != nil{
		rsCartItems = model.CartItems{
			CartId:        uuid.NewString(),
			CartUserId:    strUserId,
			CartProductId: newCartItems.CartProductId,
			CartQuantity:  newCartItems.CartQuantity,
			CartCreatedAt: time.Now(),
			CartUpdatedAt: time.Now(),
			CartColor:     newCartItems.CartColor, //init state is in cart
		}
		cartItems, err := h.repo.AddCartItems(rsCartItems)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, helper.BuildResponse(-1, "error when add cart items", err.Error()))
			return
		}
		ctx.JSON(http.StatusOK, helper.BuildResponse(1, "add cart items successfully!", cartItems))
	} else {
		var cartItems model.CartItems
		if checkExistCart.CartColor != newCartItems.CartColor {
			rsCartItems = model.CartItems{
				CartId:        uuid.NewString(),
				CartUserId:    strUserId,
				CartProductId: newCartItems.CartProductId,
				CartQuantity:  newCartItems.CartQuantity,
				CartCreatedAt: time.Now(),
				CartUpdatedAt: time.Now(),
				CartColor:     newCartItems.CartColor, //init state is in cart
			}
			cartItems, err := h.repo.AddCartItems(rsCartItems)
			if err != nil {
				ctx.JSON(http.StatusInternalServerError, helper.BuildResponse(-1, "error when add cart items", err.Error()))
				return
			}
			ctx.JSON(http.StatusOK, helper.BuildResponse(1, "add cart items successfully!", cartItems))
		}else {
			rsCartItems = model.CartItems{
				CartId:        checkExistCart.CartId,
				CartUserId:    strUserId,
				CartProductId: newCartItems.CartProductId,
				CartQuantity:  newCartItems.CartQuantity + checkExistCart.CartQuantity,
				CartCreatedAt: checkExistCart.CartCreatedAt,
				CartUpdatedAt: time.Now(),
				CartColor:     newCartItems.CartColor, //init state is in cart
			}
			cartItems, err = h.repo.UpdateCartItems(rsCartItems)
			if err != nil {
				ctx.JSON(http.StatusInternalServerError, helper.BuildResponse(-1, "error when update cart items", err.Error()))
				return
			}
			ctx.JSON(http.StatusOK, helper.BuildResponse(1, "update cart items successfully!", cartItems))
		}
	}

}

func (h *cartItemsHandler) DeleteCartItems(ctx *gin.Context) {
	id := ctx.Param("id")
	cartItems, err := h.repo.DeleteCartItems(id)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, helper.BuildResponse(-1, "error when find cart items", err.Error()))
		return
	}
	ctx.JSON(http.StatusOK, helper.BuildResponse(1, "delete cart items successfully!", cartItems))
}

func (h *cartItemsHandler) GetCartItemsByUserId(ctx *gin.Context) {
	userId, isExist := ctx.Get("userID")
	if !isExist {
		ctx.JSON(http.StatusBadRequest, helper.BuildResponse(-1, "unauthorized", "Invalid token"))
		return
	}
	listCart, err := h.repo.GetCartItemsByUserId(fmt.Sprint(userId))
	for i := 0; i<len(listCart) ; i++ {
		product, err := repository.NewProductRepository().GetProduct(listCart[i].CartProductId)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, helper.BuildResponse(-1, "error when find product with id - "+ listCart[i].CartProductId, err.Error()))
			return
		}
		var p response.ProductResponse
		p, err = p.ProductToProductResponse(product)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, helper.BuildResponse(-1, "Cant convert json to array", err.Error()))
			return
		}
		listCart[i].CartProductData = p
	}
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, helper.BuildResponse(-1, "error when find cart items", err.Error()))
		return
	}
	ctx.JSON(http.StatusOK, helper.BuildResponse(1, "get list cartItems successfully!", listCart))
}

func (h *cartItemsHandler) UpdateCartItems(ctx *gin.Context) {
	userId, isExist := ctx.Get("userID")
	if !isExist {
		ctx.JSON(http.StatusBadRequest, helper.BuildResponse(-1, "unauthorized", "Invalid token"))
		return
	}
	id := ctx.Param("id")
	quantity := ctx.Query("quantity")
	intQuantity, err := strconv.Atoi(quantity)
	if err != nil || intQuantity <= 0{
		ctx.JSON(http.StatusBadRequest, helper.BuildResponse(-1, "Số lượng sản phẩm không hợp lệ", ""))
		return
	}
	modifyCart := model.CartItems{
		CartId:        id,
		CartUserId:    fmt.Sprint(userId),
		CartQuantity:  intQuantity,
		CartUpdatedAt: time.Now(),
	}
	rsCartItems, err := h.repo.UpdateCartItems(modifyCart)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, helper.BuildResponse(-1, "error when find cart items", err.Error()))
		return
	}

	ctx.JSON(http.StatusOK, helper.BuildResponse(1, "update cart items successfully!", rsCartItems))
}
