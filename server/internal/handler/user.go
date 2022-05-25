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
	"golang.org/x/crypto/bcrypt"
)

type UserHandler interface {
	AddUser(*gin.Context)
	SignInUser(*gin.Context)
	GetUser(*gin.Context)
	UpdateUser(*gin.Context)
	ChangePassword(*gin.Context)
	GetAllUser(*gin.Context)
	ForgotPassword(*gin.Context)
	VerifyOtp(*gin.Context)
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

func (h *userHandler) AddUser(ctx *gin.Context) {
	var registerUser request.RegisterForm

	if err := ctx.ShouldBindJSON(&registerUser); err != nil {
		ctx.JSON(http.StatusBadRequest, helper.BuildResponse(-1, "invalid user format", err.Error()))
		return
	}

	_, err := h.repo.GetUserByEmail(registerUser.Email)
	if err == nil {
		ctx.JSON(http.StatusBadRequest, helper.BuildResponse(-1, "Email has already existed", ""))
		return
	}
	userRole, err := repository.NewRoleRepository().GetRoleByName("USER")
	if err != nil {
		ctx.JSON(http.StatusBadRequest, helper.BuildResponse(-1, "error when find USER role", err.Error()))
		return
	} else {
		user := model.User{
			UserID:        uuid.NewString(),
			UserRoleID:    userRole.RoleID,
			UserEmail:     registerUser.Email,
			UserPassword:  registerUser.Password,
			UserFirstName: registerUser.FirstName,
			UserLastName:  registerUser.LastName,
			UserPhone:     registerUser.Phone,
			UserAddress:   registerUser.Address,
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
		ctx.JSON(http.StatusBadRequest, helper.BuildResponse(-1, "invalid login form format", err.Error()))
		return
	}
	loginUser, err := h.repo.GetUserByEmail(loginForm.Email)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, helper.BuildResponse(-1, "error when find user by email", err.Error()))
		return
	} else {
		isTrue := comparePass(loginUser.UserPassword, loginForm.Password)
		if isTrue {
			fmt.Println("Login with: ", loginUser.UserID)
			token := GenerateToken(loginUser.UserID)
			ctx.JSON(http.StatusOK, helper.BuildResponse(1, "login successfully!", token))
			return
		} else {
			ctx.JSON(http.StatusUnauthorized, helper.BuildResponse(-1, "error when login", "Password not match!"))
			return
		}
	}
}

func (h *userHandler) GetUser(ctx *gin.Context) {

	if userID, isExist := ctx.Get("userID"); isExist {
		checkUser, err := repository.NewUserRepository().GetUser(fmt.Sprint(userID))
		if err == nil {
			ctx.JSON(http.StatusOK, helper.BuildResponse(1, "get user information successfully!", checkUser))
			return 
		} else {
			ctx.AbortWithStatusJSON(http.StatusInternalServerError, helper.BuildResponse(-1, "Error when find USER", err.Error()))
		}
	} else {
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, helper.BuildResponse(-1, "Not Exist session", ""))
	}
}

func (h *userHandler) UpdateUser(ctx *gin.Context) {

}

func (h *userHandler) ChangePassword(ctx *gin.Context) {

	if userID, isExist := ctx.Get("userID"); isExist {
		checkUser, err := repository.NewUserRepository().GetUser(fmt.Sprint(userID))
		if err == nil {
			role, err := repository.NewRoleRepository().GetRole(checkUser.UserRoleID)
			if err != nil {
				ctx.AbortWithStatusJSON(http.StatusInternalServerError, helper.BuildResponse(-1, "Error when find ROLE", err.Error()))
				return
			} else {
				if role.RoleName == "USER" {
					var changePassForm request.ChangePasswordForm
					if err := ctx.ShouldBindJSON(&changePassForm); err != nil {
						ctx.JSON(http.StatusBadRequest, helper.BuildResponse(-1, "invalid input format", err.Error()))
						return
					}
					fmt.Print(checkUser.UserPassword)
					isAuth := comparePass(checkUser.UserPassword, changePassForm.OldPassword)
					if !isAuth {
						ctx.JSON(http.StatusBadRequest, helper.BuildResponse(-1, "old password not match", err.Error()))
						return
					} else {
						var requestUser model.User
						requestUser.UserID = checkUser.UserID
						requestUser.UserPassword = changePassForm.NewPassword
						hashPass(&requestUser.UserPassword)
						updateUser, err := h.repo.UpdateUser(requestUser)
						if err != nil {
							ctx.JSON(http.StatusInternalServerError, helper.BuildResponse(-1, "error when update password for user", err.Error()))
							return
						}
						updateUser.UserPassword = ""
						ctx.JSON(http.StatusOK, helper.BuildResponse(1, "update password successfully", updateUser))
					}

				} else {
					ctx.AbortWithStatusJSON(http.StatusForbidden, helper.BuildResponse(-1, "only with USER role", ""))
					return
				}
			}
		} else {
			ctx.AbortWithStatusJSON(http.StatusInternalServerError, helper.BuildResponse(-1, "Error when find USER", err.Error()))
		}
	} else {
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, helper.BuildResponse(-1, "Not Exist session", ""))
	}
}
func (h *userHandler) GetAllUser(ctx *gin.Context) {
	userRole, err := repository.NewRoleRepository().GetRoleByName("USER")
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, helper.BuildResponse(-1, "Error when find USER role", err.Error()))
		return
	}
	listUser, err := h.repo.GetAllUser(userRole.RoleID)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, helper.BuildResponse(-1, "Error when get all users", err.Error()))
		return
	}
	ctx.JSON(http.StatusOK, helper.BuildResponse(1, "get list user successfully", listUser))
}
func (h *userHandler) ForgotPassword(ctx *gin.Context) {
	userEmail := ctx.Query("email")
	user, err := h.repo.GetUserByEmail(userEmail)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, helper.BuildResponse(-1, "Email is not exist", err.Error()))
		return
	}
	//if user is exist => generate new Otp
	otpValue := helper.GenerateOtp()

	//send email
	var to []string
	to = append(to, userEmail)
	sendErr := helper.SendEmailForgotPassword(to, otpValue)

	if sendErr != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, helper.BuildResponse(-1, "Error when send email reset password", sendErr.Error()))
		return
	}

	newOtp := model.Otp{
		OtpId:        uuid.NewString(),
		UserId: user.UserID,
		OtpCreateAt:  time.Now(),
		OtpUpdateAt:  time.Now(),
		OtpUsedOk:    0,
		OtpValue:     otpValue,
		OtpTimeStart: time.Now(),
		OtpTimeEnd:   time.Now().Add(120e9),
	}
	_, otpErr := repository.NewOtpRepository().AddOtp(newOtp)
	if otpErr != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, helper.BuildResponse(-1, "Error when create otp", otpErr.Error()))
		return
	}

	ctx.JSON(http.StatusOK, helper.BuildResponse(1, "send email verify otp success", user.UserID))

}

func (h *userHandler) VerifyOtp(ctx *gin.Context) {
	var otpRequest request.VerifyOtpRequest
	if err := ctx.ShouldBindJSON(&otpRequest); err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, helper.BuildResponse(-1, "Invalid Params", err.Error()))
		return
	}
	_ , err := h.repo.GetUser(otpRequest.UserId)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, helper.BuildResponse(-1, "User is not existed", err.Error()))
		return
	}
	otp, err := repository.NewOtpRepository().GetOtpByUserId(otpRequest.UserId)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, helper.BuildResponse(-1, "Error when find otp", err.Error()))
		return
	}
	if otp.OtpValue == otpRequest.OtpValue {
		if time.Now().Before(otp.OtpTimeEnd) {
			otp.OtpUsedOk = 1
			otp.OtpUpdateAt = time.Now()
			_, err = repository.NewOtpRepository().UpdateOtpStatus(otp)
			if err != nil {
				ctx.AbortWithStatusJSON(http.StatusInternalServerError, helper.BuildResponse(-1, "Error when verify otp", err.Error()))
				return
			}
			ctx.JSON(http.StatusOK, helper.BuildResponse(1, "verify otp success", ""))
			return
		}else {
			ctx.AbortWithStatusJSON(http.StatusInternalServerError, helper.BuildResponse(-1, "Otp is expired", err.Error()))
			return
		}

	} else {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, helper.BuildResponse(-1, "Invalid otp value", err.Error()))
		return
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
