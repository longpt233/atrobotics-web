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

type DeliveryAddressHandler interface {
	AddDeliveryAddress(*gin.Context)
	GetDeliveryAddressByUser(*gin.Context)
	UpdateDeliveryAddress(*gin.Context)
	DeleteDeliveryAddress(*gin.Context)
}

type deliveryAddressHandler struct {
	repo repository.DeliveryAddressRepository
}

func NewDeliveryAddressHandler() DeliveryAddressHandler{
	return &deliveryAddressHandler{
		repo: repository.NewDeliveryAddressRepository(),
	}
}

func (h *deliveryAddressHandler) AddDeliveryAddress(ctx *gin.Context){
	userId, isExist := ctx.Get("userID")
	if !isExist {
		ctx.JSON(http.StatusBadRequest, helper.BuildResponse(-1, "unauthorized", "Invalid token"))
		return
	}
	var newAddress request.DeliveryAddressRequest
	if err := ctx.ShouldBindJSON(&newAddress); err != nil {
		ctx.JSON(http.StatusBadRequest, helper.BuildResponse(-1, "invalid request body", err.Error()))
		return
	}

	strUserId := fmt.Sprint(userId)
	newDeliveryAddress := model.DeliveryAddress{
		DeliveryAddressId: uuid.NewString(),
		UserId: strUserId,
		Fullname: newAddress.Fullname,
		Phone: newAddress.Phone,
		City: newAddress.City,
		District: newAddress.District,
		Ward: newAddress.Ward,
		DetailAddress: newAddress.DetailAddress,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
	rsAddress, err := h.repo.AddDeliveryAddress(newDeliveryAddress)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, helper.BuildResponse(-1, "error when add delivery address", err.Error()))
			return
	}
	ctx.JSON(http.StatusOK, helper.BuildResponse(1, "Add delivery successfully!", rsAddress))

}
func (h *deliveryAddressHandler) GetDeliveryAddressByUser(ctx *gin.Context){
	userId, isExist := ctx.Get("userID")
	if !isExist {
		ctx.JSON(http.StatusBadRequest, helper.BuildResponse(-1, "unauthorized", "Invalid token"))
		return
	}
	listAddress, err := h.repo.GetDeliveryAddressByUser(fmt.Sprint(userId))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, helper.BuildResponse(-1, "Error when get list address", err.Error()))
		return
	}
	ctx.JSON(http.StatusOK, helper.BuildResponse(1, "get list address successfully!", listAddress))

}
func (h *deliveryAddressHandler) UpdateDeliveryAddress(ctx *gin.Context){
	userId, isExist := ctx.Get("userID")
	if !isExist {
		ctx.JSON(http.StatusBadRequest, helper.BuildResponse(-1, "unauthorized", "Invalid token"))
		return
	}
	var modifyAddress request.DeliveryAddressRequest
	if err := ctx.ShouldBindJSON(&modifyAddress); err != nil {
		ctx.JSON(http.StatusBadRequest, helper.BuildResponse(-1, "invalid request body", err.Error()))
		return
	}

	addressId := ctx.Param("id")
	strUserId := fmt.Sprint(userId)
	modifyDeliveryAddress := model.DeliveryAddress{
		DeliveryAddressId: addressId,
		UserId: strUserId,
		Fullname: modifyAddress.Fullname,
		Phone: modifyAddress.Phone,
		City: modifyAddress.City,
		District: modifyAddress.District,
		Ward: modifyAddress.Ward,
		DetailAddress: modifyAddress.DetailAddress,
		UpdatedAt: time.Now(),
	}
	rsAddress, err := h.repo.UpdateDeliveryAddress(modifyDeliveryAddress)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, helper.BuildResponse(-1, "error when update delivery address", err.Error()))
			return
	}
	ctx.JSON(http.StatusOK, helper.BuildResponse(1, "Update delivery successfully!", rsAddress))
}
func (h *deliveryAddressHandler) DeleteDeliveryAddress(ctx *gin.Context){
	addressId := ctx.Param("id")

	address, err := h.repo.DeleteDeliveryAddress(addressId)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, helper.BuildResponse(-1, "Address Id not found", err.Error()))
		return
	}
	ctx.JSON(http.StatusOK, helper.BuildResponse(1, "delete address successfully!", address))
}
