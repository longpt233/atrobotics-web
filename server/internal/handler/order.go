package handler

import (
	"atro/internal/helper"
	"atro/internal/repository"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

//OrderHandler --> Handler for Order Entity
type OrderHandler interface {
	OrderProduct(*gin.Context)
	GetOrderProduct(*gin.Context)
	GetAllOrderProduct(*gin.Context) 
	UpdateOrderProduct(*gin.Context)
}

type orderHandler struct {
	repo repository.OrderRepository
}

//NewOrderHandler --> return new Order Handler
func NewOrderHandler() OrderHandler {
	return &orderHandler{
		repo: repository.NewOrderRepository(),
	}
}

func (h *orderHandler) OrderProduct(ctx *gin.Context) {
	prodIDStr := ctx.Param("product")
	quantityIDStr := ctx.Param("quantity")
		if quantityID, err := strconv.Atoi(quantityIDStr); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
		} else {
			userID, isExist := ctx.Get("userID")
			if isExist {
				if err := h.repo.OrderProduct(fmt.Sprint(userID), prodIDStr, quantityID); err != nil {
					ctx.JSON(http.StatusBadRequest, gin.H{
						"error": err.Error(),
					})
				} else {
					ctx.String(http.StatusOK, "Product Successfully ordered")
				}
			}else {
				ctx.AbortWithStatusJSON(http.StatusUnauthorized, helper.BuildResponse(-1, "Not Exist session", ""))
			}
			
		}

}

func (h *orderHandler) GetOrderProduct(ctx *gin.Context) {
}

func (h *orderHandler) GetAllOrderProduct(ctx *gin.Context) {
}

func (h *orderHandler) UpdateOrderProduct(ctx *gin.Context) {
}
