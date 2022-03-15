package handler

import (
	"atro/internal/helper"
	"atro/internal/model"
	"atro/internal/model/request"
	"atro/internal/repository"
	"encoding/json"
	"fmt"
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
	repo        repository.OrderRepository
	repoProduct repository.ProductRepository
}

//NewOrderHandler --> return new Order Handler
func NewOrderHandler() OrderHandler {
	return &orderHandler{
		repo:        repository.NewOrderRepository(),
		repoProduct: repository.NewProductRepository(),
	}
}

func (h *orderHandler) OrderProduct(ctx *gin.Context) {

	// lấy thông tin order từ request
	var orderForm request.OrderRequest
	if err := ctx.ShouldBindJSON(&orderForm); err != nil {
		ctx.JSON(http.StatusBadRequest, helper.BuildResponse(-1, "invalid input", err.Error()))
		return
	}

	// lây thông tin user đã được set từ middleware
	id, _ := ctx.Get("userID")
	if id == nil {
		ctx.JSON(http.StatusInternalServerError, helper.BuildResponse(-1, "lấy id user bị lỗ rồi khóc đi ", ""))
		return
	}

	// chuyển đổi từ đơn về một ban ghi để lưu db
	order, err := h.OrderRequestToOrder(&orderForm, int(id.(float64))) // parse asset id to int syntax
	if err != nil {
		ctx.JSON(http.StatusBadRequest, helper.BuildResponse(-1, "không thể  tạo đơn", err.Error()))
		return
	}

	// lưu vô db
	rsOrder, err := h.repo.OrderProduct(order)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, helper.BuildResponse(-1, "error when add product (có thể là id bị sai)", err.Error()))
		return
	}
	ctx.JSON(http.StatusOK, helper.BuildResponse(1, "create order successfully!", rsOrder))

}

func (h *orderHandler) GetOrderProduct(ctx *gin.Context) {
}

func (h *orderHandler) GetAllOrderProduct(ctx *gin.Context) {
}

func (h *orderHandler) UpdateOrderProduct(ctx *gin.Context) {
}

func (h *orderHandler) OrderRequestToOrder(orderForm *request.OrderRequest, userId int) (model.Order, error) {

	// từ cái json gủi lên tính tiền
	total := 0.0
	for _, element := range orderForm.ProductOrders {
		if productDetail, err := h.repoProduct.GetProduct(element.ProductId); err == nil {
			fmt.Println("product detail ", productDetail)
			total = total + productDetail.ProductPrice*float64(element.Quantity)

		} else { 
			return model.Order{}, err
		}

	}

	// parse để lưu vô db
	productOrdersInfo, err := json.Marshal(orderForm)
	if err != nil {
		return model.Order{}, err
	}

	order := model.Order{
		UserId:         userId,
		OrderDetail:    string(productOrdersInfo),
		OrderPrice:     float32(total),
		OrderCreatedAt: time.Now(),
		OrderStatus:    orderForm.TypeOrder,
	}
	return order, nil
}
