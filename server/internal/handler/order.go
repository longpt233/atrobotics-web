package handler

import (
	"atro/internal/helper"
	"atro/internal/model"
	"atro/internal/model/request"
	"atro/internal/repository"
	"encoding/json"
	"net/http"
	"time"

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

	// lấy thông tin order từ request
	var orderForm request.OrderRequest
	if err := ctx.ShouldBindJSON(&orderForm); err != nil {
		ctx.JSON(http.StatusBadRequest, helper.BuildResponse(-1, "invalid id input", err.Error()))
		return
	}

	// lây thông tin user từ token

	// chuyển đổi từ đơn về một ban ghi để lưu db
	order, err := OrderRequestToOrder(&orderForm)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, helper.BuildResponse(-1, "không thể  tạo đơn", err.Error()))
		return
	}

	// lưu vô db
	product, err := h.repo.OrderProduct(order)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, helper.BuildResponse(-1, "error when add product", err.Error()))
		return
	}
	ctx.JSON(http.StatusOK, helper.BuildResponse(1, "add product successfully!", product))

}

func (h *orderHandler) GetOrderProduct(ctx *gin.Context) {
}

func (h *orderHandler) GetAllOrderProduct(ctx *gin.Context) {
}

func (h *orderHandler) UpdateOrderProduct(ctx *gin.Context) {
}

func OrderRequestToOrder(orderForm *request.OrderRequest) (model.Order, error) {

	// parse
	productOrdersInfo, err := json.Marshal(orderForm.ProductOrders)
	if err != nil {
		return model.Order{}, err
	}
	order := model.Order{
		UserId:         1,
		OrderDetail:    string(productOrdersInfo),
		OrderPrice:     0,
		OrderCreatedAt: time.Now(),
		OrderStatus:    1,
	}
	return order, nil
}
