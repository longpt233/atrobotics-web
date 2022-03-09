package handler

import (
	"atro/internal/helper"
	"atro/internal/model"
	"atro/internal/model/request"
	"atro/internal/repository"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

type UserHandler interface {
	GetUser(*gin.Context)
	AddUser(*gin.Context)
	SignInUser(*gin.Context)
}

type userHandler struct {
	repo repository.UserRepository
}

//NewUserHandler --> returns new user handler
func NewUserHandler() UserHandler { // interface mac dinh la kieu con tro -> tra ve dia chi

	a := userHandler{
		repo: repository.NewUserRepository(),
	}

	return &a // tra ve dia chi cua 1 struct userHandler , cai struct nay phai implement het cua interface UserHandler
}
func (h *userHandler) GetUser(ctx *gin.Context) {
	id := ctx.Query("id")
	intID, err := strconv.Atoi(id)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	user, err := h.repo.GetUser(intID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return

	}
	ctx.JSON(http.StatusOK, helper.BuildResponse(1,"get user information successfully", user))

}

func (h *userHandler) AddUser(ctx *gin.Context) {
	var registerUser request.RegisterForm

	if err := ctx.ShouldBindJSON(&registerUser); err != nil {
		ctx.JSON(http.StatusBadRequest, helper.BuildResponse(-1, "invalid user format", err.Error()))
		return
	}

	userRole, err := repository.NewRoleRepository().GetRoleByName("USER")
	if err != nil {
		ctx.JSON(http.StatusBadRequest, helper.BuildResponse(-1, "error when find USER role", err.Error()))
		return
	} else {
		user := model.User{
			UserRoleID: userRole.RoleID,
			UserEmail: registerUser.Email,
			UserPassword: registerUser.Password,
			UserFirstName: registerUser.FirstName,
			UserLastName: registerUser.LastName,
			UserPhone: registerUser.Phone,
			UserAddress: registerUser.Address,
		}
		hashPass(&user.UserPassword)
		fmt.Print("user register: ", user)
		registerUser, err := h.repo.AddUser(user)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, helper.BuildResponse(-1, "error when add user", err.Error()))
			return
		} else {
			registerUser.UserPassword = ""
			ctx.JSON(http.StatusOK, helper.BuildResponse(1, "create user successfully!", registerUser))
		}
	}
}

func (h *userHandler) SignInUser(ctx *gin.Context) {
	var loginForm request.LoginForm

	//check valid body
	if err := ctx.ShouldBindJSON(&loginForm); err != nil {
		ctx.JSON(http.StatusBadRequest, helper.BuildResponse(-1,"invalid login form format", err.Error()))
		return
	}
	loginUser, err := h.repo.GetUserByEmail(loginForm.Email)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, helper.BuildResponse(-1, "error when find user by email", err.Error()))
		return
	} else {
		isTrue := comparePass(loginUser.UserPassword, loginForm.Password);
		if isTrue {
			fmt.Println("Login with: ", loginUser.UserID)
			token := GenerateToken(uint(loginUser.UserID))
			ctx.JSON(http.StatusOK, helper.BuildResponse(1, "login successfully!", token))
			return
		} else {
			ctx.JSON(http.StatusUnauthorized, helper.BuildResponse(-1,"error when login", "Password not match!"))
			return
		}
	}

}

func hashPass(pass *string) {
	bytePass := []byte(*pass)
	hPass, _ := bcrypt.GenerateFromPassword(bytePass, bcrypt.DefaultCost)
	*pass = string(hPass)
}

func comparePass(dbPass, pass string) bool {
	return bcrypt.CompareHashAndPassword([]byte(dbPass), []byte(pass)) == nil
}
