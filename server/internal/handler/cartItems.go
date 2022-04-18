package handler

import (
	"atro/internal/helper"
	"atro/internal/model"
	"atro/internal/model/request"
	"atro/internal/repository"
	"fmt"
	"net/http"
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
	var rsCartItems model.CardItems
	if err != nil {
		rsCartItems = model.CardItems{
			CartId:        uuid.NewString(),
			CartUserId:    strUserId,
			CartProductId: newCartItems.CartProductId,
			CartQuantity:  newCartItems.CartQuantity,
			CartCreatedAt: time.Now(),
			CartUpdatedAt: time.Now(),
			CartState:     0, //init state is in cart
		}
		cartItems, err := h.repo.AddCartItems(rsCartItems)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, helper.BuildResponse(-1, "error when add cart items", err.Error()))
			return
		}
		ctx.JSON(http.StatusOK, helper.BuildResponse(1, "add cart items successfully!", cartItems))
	} else {
		rsCartItems = model.CardItems{
			CartId:        checkExistCart.CartId,
			CartUserId:    strUserId,
			CartProductId: newCartItems.CartProductId,
			CartQuantity:  newCartItems.CartQuantity + checkExistCart.CartQuantity,
			CartCreatedAt: checkExistCart.CartCreatedAt,
			CartUpdatedAt: time.Now(),
			CartState:     0, //init state is in cart
		}
		cartItems, err := h.repo.UpdateCartItems(rsCartItems)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, helper.BuildResponse(-1, "error when update cart items", err.Error()))
			return
		}
		ctx.JSON(http.StatusOK, helper.BuildResponse(1, "update cart items successfully!", cartItems))
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
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, helper.BuildResponse(-1, "error when find cart items", err.Error()))
		return
	}
	ctx.JSON(http.StatusOK, helper.BuildResponse(1, "get list cartItems successfully!", listCart))
}

func (h *cartItemsHandler) UpdateCartItems(ctx *gin.Context) {
	var requestCart request.CartItemsRequest
	userId, isExist := ctx.Get("userID")
	if !isExist {
		ctx.JSON(http.StatusBadRequest, helper.BuildResponse(-1, "unauthorized", "Invalid token"))
		return
	}
	_, errProduct := repository.NewProductRepository().GetProduct(requestCart.CartProductId)
	if errProduct != nil {
		ctx.JSON(http.StatusBadRequest, helper.BuildResponse(-1, "product id not found", errProduct.Error()))
		return
	}
	if err := ctx.ShouldBindJSON(&requestCart); err != nil {
		ctx.JSON(http.StatusBadRequest, helper.BuildResponse(-1, "Invalid request body", err.Error()))
		return
	}
	id := ctx.Param("id")
	modifyCart := model.CardItems{
		CartId:        id,
		CartUserId:    fmt.Sprint(userId),
		CartProductId: requestCart.CartProductId,
		CartQuantity:  requestCart.CartQuantity,
		CartUpdatedAt: time.Now(),
	}
	rsCartItems, err := h.repo.UpdateCartItems(modifyCart)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, helper.BuildResponse(-1, "error when find cart items", err.Error()))
		return
	}

	ctx.JSON(http.StatusOK, helper.BuildResponse(1, "update cart items successfully!", rsCartItems))
}
