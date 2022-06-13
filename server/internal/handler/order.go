package handler

import (
	"atro/internal/helper"
	"atro/internal/model"
	"atro/internal/model/request"
	"atro/internal/model/response"
	"atro/internal/repository"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

//OrderHandler --> Handler for Order Entity
type OrderHandler interface {
	OrderProduct(*gin.Context)
	GetOrderProduct(*gin.Context)
	GetAllOrderProduct(*gin.Context)
	UpdateOrderStatus(*gin.Context)
	GetAllOrderByStatus(*gin.Context)
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

func (h *orderHandler) OrderProduct(ctx *gin.Context) { // TODO: transaction?
	userId, isExist := ctx.Get("userID")
	if !isExist {
		ctx.JSON(http.StatusBadRequest, helper.BuildResponse(-1, "unauthorized", "Invalid token"))
		return
	}
	addressId := ctx.Query("addressId")
	address, err := repository.NewDeliveryAddressRepository().GetDeliveryAddressById(addressId)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, helper.BuildResponse(-1, "error when get address", err))
		return
	}
	orderAddress := address.Phone + ", " + address.Fullname + ", " + address.DetailAddress + ", " + address.Ward + ", " + address.District + ", " + address.City
	listCartItems, err := repository.NewCartItemsRepository().GetCartItemsByUserId(fmt.Sprint(userId))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, helper.BuildResponse(-1, "error when get list cart items", err))
		return
	}
	if len(listCartItems) == 0 {
		ctx.JSON(http.StatusBadRequest, helper.BuildResponse(-1, "Cart items is empty, please choose product to order", err))
		return
	}
	var orderDetails []model.OrderProduct
	for _, cItems := range listCartItems {
		product, err := h.repoProduct.GetProduct(cItems.CartProductId)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, helper.BuildResponse(-1, "product is not exist", err))
			return
		}
		var productImgs []string
		err = json.Unmarshal([]byte(product.ProductImages), &productImgs)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, helper.BuildResponse(-1, "Error when get first product images", err))
		}
		orderProduct := model.OrderProduct{
			ProductId:        product.ProductID,
			ProductName:      product.ProductName,
			ProductImage:     productImgs[0],
			CurrentPrice:     product.ProductPrice,
			Quantity:         cItems.CartQuantity,
			ShortDescription: product.ProductShortDesc,
			Colors:           cItems.CartColor,
		}
		orderDetails = append(orderDetails, orderProduct)
	}
	orderDetailsStr, err := json.Marshal(orderDetails)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, helper.BuildResponse(-1, "Error when parse to json", err))
		return
	}
	order := model.Order{
		OrderId:        uuid.NewString(),
		OrderItems:     string(orderDetailsStr),
		OrderCreatedAt: time.Now(),
		OrderStatus:    1,
		OrderCode:      helper.GenerateOrderCode(),
		OrderAddress:   orderAddress,
		UserId:         fmt.Sprint(userId),
	}
	createOrder, err := h.repo.OrderProduct(order)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, helper.BuildResponse(-1, "Error when add order", err))
		return
	}
	for _, cItems := range listCartItems {
		_, err := repository.NewCartItemsRepository().DeleteCartItems(cItems.CartId)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, helper.BuildResponse(-1, "Error when clear cart", err))
			return
		}
	}
	orderResponse := response.OrderResponse{
		OrderId:        createOrder.OrderId,
		OrderItems:     orderDetails,
		OrderCreatedAt: createOrder.OrderCreatedAt,
		OrderStatus:    createOrder.OrderStatus,
		OrderCode:      createOrder.OrderCode,
		OrderAddress:   createOrder.OrderAddress,
		UserId:         createOrder.UserId,
	}
	ctx.JSON(http.StatusOK, helper.BuildResponse(1, "order product successfully!", orderResponse))
}

func (h *orderHandler) GetOrderProduct(ctx *gin.Context) {
	orderId := ctx.Param("id")
	userId, isExist := ctx.Get("userID")
	if !isExist {
		ctx.JSON(http.StatusBadRequest, helper.BuildResponse(-1, "unauthorized", "Invalid token"))
		return
	}
	order, err := h.repo.GetOrderById(fmt.Sprint(userId), orderId)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, helper.BuildResponse(-1, "error when find order", err.Error()))
		return
	}

	orderRes, err := orderToOrderResponse(order)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, helper.BuildResponse(-1, "error when parse json to list order product", err.Error()))
			return
		}
	ctx.JSON(http.StatusOK, helper.BuildResponse(1, "get list order by user successfully!", orderRes))

}

func (h *orderHandler) GetAllOrderProduct(ctx *gin.Context) {
	userId, isExist := ctx.Get("userID")
	if !isExist {
		ctx.JSON(http.StatusBadRequest, helper.BuildResponse(-1, "unauthorized", "Invalid token"))
		return
	}

	// tạo query sort
	sortBy := ctx.Query("sort") //sort params: asc, desc
	if sortBy == "" {
		sortBy = "DESC" // sortBy is default DESC with column order_created_at
	} else {
		sortBy = strings.ToUpper(sortBy)
	}

	// tao query limit
	strLimit := ctx.Query("limit")
	fmt.Println("param limit", strLimit)
	limit := -1 // with a value as -1 for gorms Limit method, we'll get a request without limit as default
	if strLimit != "" {
		limit, err := strconv.Atoi(strLimit)
		if err != nil || limit < -1 {
			ctx.JSON(http.StatusBadRequest, helper.BuildResponse(-1, "limit query parameter is no valid number", err.Error()))
			return
		}
	}

	// tạo query offset
	strOffset := ctx.Query("offset")
	offset := -1
	if strOffset != "" {
		offset, err := strconv.Atoi(strOffset)
		if err != nil || offset < -1 {
			ctx.JSON(http.StatusBadRequest, helper.BuildResponse(-1, "offset query parameter is no valid number", err.Error()))
			return
		}
	}

	// gửi query
	rsOrders, err := h.repo.GetAllOrderOptions(fmt.Sprint(userId), limit, offset, sortBy)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, helper.BuildResponse(-1, "not found !", err.Error()))
		return
	}

	var orderResponses []response.OrderResponse
	for _, o := range rsOrders {
		orderRes, err := orderToOrderResponse(o)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, helper.BuildResponse(-1, "error when parse json to list order product", err.Error()))
			return
		}
		orderResponses = append(orderResponses, orderRes)
	}
	ctx.JSON(http.StatusOK, helper.BuildResponse(1, "get list products successfully!", orderResponses))
}

func (h *orderHandler) UpdateOrderStatus(ctx *gin.Context) {
	orderId := ctx.Query("orderId")
	status := ctx.Query("status")
	intStatus, err := strconv.Atoi(status)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, helper.BuildResponse(-1, "Invalid status params", err.Error()))
		return
	}
	order, err := h.repo.UpdateOrderStatus(orderId, intStatus)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, helper.BuildResponse(-1, "Error when update order status", err.Error()))
		return
	}
	orderResponse, err := orderToOrderResponse(order)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, helper.BuildResponse(-1, "error when parse json to list order product", err.Error()))
		return
	}
	ctx.JSON(http.StatusOK, helper.BuildResponse(1, "update order status success fully!", orderResponse))

}

func (h *orderHandler) OrderRequestToOrder(orderForm *request.OrderRequest, userId string) (model.Order, error) {

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
		OrderItems:     string(productOrdersInfo),
		OrderCode:      "",
		OrderCreatedAt: time.Now(),
		OrderStatus:    orderForm.TypeOrder,
	}
	return order, nil
}
func orderToOrderResponse(order model.Order) (response.OrderResponse, error) {
	var orderItems []model.OrderProduct
	err := json.Unmarshal([]byte(order.OrderItems), &orderItems)
	if err != nil {
		return response.OrderResponse{}, err
	}
	orderResponse := response.OrderResponse{
		OrderId:        order.OrderId,
		OrderItems:     orderItems,
		OrderCreatedAt: order.OrderCreatedAt,
		OrderStatus:    order.OrderStatus,
		OrderCode:      order.OrderCode,
		OrderAddress:   order.OrderAddress,
		UserId:         order.UserId,
	}
	return orderResponse, nil
}
func(h *orderHandler) GetAllOrderByStatus(ctx *gin.Context) {
	status := ctx.Query("status")
	intStatus, err := strconv.Atoi(status)
	if err != nil || intStatus < 0 || intStatus > 4 {
		ctx.JSON(http.StatusBadRequest, helper.BuildResponse(-1, "Invalid status params", err.Error()))
		return
	}
	listOrder, err := h.repo.GetAllOrderByStatus(intStatus)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, helper.BuildResponse(-1, "Error when get list order by status !", err.Error()))
		return
	}
	var orderResponses []response.OrderResponse
	for _, o := range listOrder {
		orderRes, err := orderToOrderResponse(o)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, helper.BuildResponse(-1, "error when parse json to list order product", err.Error()))
			return
		}
		orderResponses = append(orderResponses, orderRes)
	}
	ctx.JSON(http.StatusOK, helper.BuildResponse(1, "get list orders by status successfully!", orderResponses))
}
