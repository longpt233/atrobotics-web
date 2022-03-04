package handler

import (
	"atro/internal/repository"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type UserHandler interface {
	GetUser(*gin.Context)
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
	ctx.JSON(http.StatusOK, user)

}
