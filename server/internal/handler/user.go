package handler

import "github.com/gin-gonic/gin"

type UserHandler interface {
	GetUser(*gin.Context)
}

type userHandler struct {
	repo repository.UserRepository
}
