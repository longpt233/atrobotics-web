package handler

import (
	"atro/internal/helper"
	"atro/internal/model"
	"atro/internal/model/request"
	"atro/internal/repository"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type BannerHandler interface {
	GetBanner(*gin.Context)
	AddBanner(*gin.Context)
	UpdateBanner(*gin.Context)
	DeleteBanner(*gin.Context)
}

type bannerHandler struct {
	repo repository.BannerRepository
}

func NewBannerHandler() BannerHandler {
	return &bannerHandler{
		repo: repository.NewBannerRepository(),
	}
}
func (h *bannerHandler) GetBanner(ctx *gin.Context) {
	id := ctx.Param("id")
	banner, err := h.repo.GetBanner((id))
	if err != nil{
		ctx.JSON(http.StatusInternalServerError, helper.BuildResponse(-1, "error when find banner", err.Error()))
		return
	}
	ctx.JSON(http.StatusOK, helper.BuildResponse(1, "get product successfully!", banner))

}
func(h *bannerHandler) AddBanner(ctx *gin.Context){
	var newBanner request.BannerRequest
	if err := ctx.ShouldBindJSON(&newBanner); err != nil {
		ctx.JSON(http.StatusBadRequest, helper.BuildResponse(-1, "invalid request body", err.Error()))
		return
	}
	rsBanner := model.Banner{
		BannerId: uuid.NewString(),
		BannerProductId: newBanner.BannerProductId,
		BannerImage: newBanner.BannerImage,
		BannerCreateAt: time.Now(),
		BannerUpdateAt: time.Now(),
	}
	banner, err := h.repo.AddBanner(rsBanner)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, helper.BuildResponse(-1, "error when add banner", err.Error()))
		return
	}
	ctx.JSON(http.StatusOK, helper.BuildResponse(1, "add banner successfully!", banner))
}
func (h *bannerHandler) UpdateBanner(ctx *gin.Context) {
	var bannerRequest request.BannerRequest
	if err := ctx.ShouldBindJSON(&bannerRequest); err != nil{
		ctx.JSON(http.StatusBadRequest, helper.BuildResponse(-1, "Invalid request body", err.Error()))
		return
	}
	id := ctx.Param("id")
	rsBanner := model.Banner{
		BannerId: id,
		BannerProductId: bannerRequest.BannerProductId,
		BannerImage: bannerRequest.BannerImage,
		BannerUpdateAt: time.Now(),
	}
	updateBanner, err := h.repo.UpdateBanner(rsBanner)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, helper.BuildResponse(-1, "error when find Banner", err.Error()))
		return
	}

	ctx.JSON(http.StatusOK, helper.BuildResponse(1, "update banner successfully!", updateBanner))
}
func (h *bannerHandler) DeleteBanner(ctx *gin.Context) {
	id := ctx.Param("id")
	banner, err := h.repo.DeleteBanner(id)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, helper.BuildResponse(-1, "error when find banner", err.Error()))
		return
	}
	ctx.JSON(http.StatusOK, helper.BuildResponse(1, "delete banner successfully!", banner))

}
