package handler

import (
	"atro/internal/helper"
	"atro/internal/model"
	"atro/internal/repository"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

type AdminHandler interface {
	SignInUser(*gin.Context)
	AddUser(*gin.Context)
}

type adminHandler struct {
	repo repository.AdminRepository
}

//NewUserHandler --> returns new user handler
func NewAdminHandler() AdminHandler { // interface mac dinh la kieu con tro -> tra ve dia chi

	a := adminHandler{
		repo: repository.NewAdminRepository(),
	}

	return &a
}

func (h *adminHandler) SignInUser(ctx *gin.Context) {
	var admin model.Admin

	// check valid input
	if err := ctx.ShouldBindJSON(&admin); err != nil {
		ctx.JSON(http.StatusBadRequest, helper.BuildResponse(-1, "invalid format", err.Error()))
	}

	// lấy ra user tương ứng với tên đăng nhập
	dbUser, err := h.repo.GetByUserName(admin.Username)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, helper.BuildResponse(-1, "invalid user", err.Error()))
		return

	}

	// check pass, trả về token
	if isTrue := comparePassword(dbUser.Password, admin.Password); isTrue {
		fmt.Println("user before", dbUser.ID)
		token := GenerateToken(dbUser.ID)
		ctx.JSON(http.StatusOK, helper.BuildResponse(1, "login successfully", token))
		return
	}
	ctx.JSON(http.StatusInternalServerError, helper.BuildResponse(-1, "internal err", "k biết lỗi đâu mà k chạy, đừng khóc"))

}

func (h *adminHandler) AddUser(ctx *gin.Context) {
	var adminUser model.Admin

	if err := ctx.ShouldBindJSON(&adminUser); err != nil {
		ctx.JSON(http.StatusBadRequest, helper.BuildResponse(-1, "invalid format", err.Error()))
		return
	}

	hashPassword(&adminUser.Password)
	user, err := h.repo.AddUser(adminUser)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, helper.BuildResponse(-1, " lỗi khi add cmnr", err.Error()))
		return

	}
	user.Password = ""
	ctx.JSON(http.StatusOK, helper.BuildResponse(1, "create user rồi nhé", user))

}

func hashPassword(pass *string) {
	bytePass := []byte(*pass)
	hPass, _ := bcrypt.GenerateFromPassword(bytePass, bcrypt.DefaultCost)
	*pass = string(hPass)
}

func comparePassword(dbPass, pass string) bool {
	return bcrypt.CompareHashAndPassword([]byte(dbPass), []byte(pass)) == nil
}
