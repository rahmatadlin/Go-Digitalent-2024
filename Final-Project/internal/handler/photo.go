package handler

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator"
	"github.com/rahmatadlin/Go-Digitalent-2024/Final-Project/internal/model"
	"github.com/rahmatadlin/Go-Digitalent-2024/Final-Project/internal/service"
	"github.com/rahmatadlin/Go-Digitalent-2024/Final-Project/pkg/response"
	"github.com/rahmatadlin/Go-Digitalent-2024/Final-Project/pkg/helper"
)

type PhotoHandler interface {
	PostPhoto(ctx *gin.Context)
	GetAllPhotosByUserId(ctx *gin.Context)
	UpdatePhoto(ctx *gin.Context)
	DeletePhoto(ctx *gin.Context)
}

type photoHandlerImpl struct {
	svc service.PhotoService
}

func NewPhotoHandler(svc service.PhotoService) PhotoHandler {
	return &photoHandlerImpl{svc: svc}
}

func (p *photoHandlerImpl) PostPhoto(ctx *gin.Context) {
	userId, err := helper.GetUserIdFromGinCtx(ctx)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, response.ErrorResponse{Message: err.Error()})
		return
	}

	photoData := model.PhotoCreate{}
	err = ctx.ShouldBindJSON(&photoData)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, response.ErrorResponse{Message: err.Error()})
		return
	}

	validate := validator.New()
	err = validate.Struct(photoData)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, response.ErrorResponse{Message: err.Error()})
		return
	}

	photo := model.Photo{}
	photo.UserId = uint32(userId)
	photo.Caption = photoData.Caption
	photo.Title = photoData.Title
	photo.PhotoUrl = photoData.PhotoUrl

	photoRes, err := p.svc.PostPhoto(ctx, photo)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, response.ErrorResponse{Message: err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, photoRes)
}

func (p *photoHandlerImpl) GetAllPhotosByUserId(ctx *gin.Context) {
	userIdStr := ctx.Request.URL.Query().Get("userId")
	if userIdStr == "" {
		ctx.JSON(http.StatusBadRequest, response.ErrorResponse{Message: "Missing User id in query"})
		return
	}
	userId, err := strconv.Atoi(userIdStr)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, response.ErrorResponse{Message: err.Error()})
		return
	}

	photos, err := p.svc.GetAllPhotosByUserId(ctx, uint32(userId))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, response.ErrorResponse{Message: err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, photos)
}

func (p *photoHandlerImpl) UpdatePhoto(ctx *gin.Context) {
	photoId, err := strconv.Atoi(ctx.Param("id"))
	if photoId == 0 || err != nil {
		ctx.JSON(http.StatusBadRequest, response.ErrorResponse{Message: err.Error()})
		return
	}

	photo, err := p.svc.GetPhotoById(ctx, uint32(photoId))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, response.ErrorResponse{Message: err.Error()})
		return
	}

	if photo.ID == 0 {
		ctx.JSON(http.StatusNotFound, response.ErrorResponse{Message: "Photo did not exist"})
		return
	}

	userId, err := helper.GetUserIdFromGinCtx(ctx)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, response.ErrorResponse{Message: err.Error()})
		return
	}

	if userId != uint32(photo.UserId) {
		ctx.JSON(http.StatusUnauthorized, response.ErrorResponse{Message: "unauthorized to do this request"})
		return
	}

	photoEditData := model.UpdatePhoto{}
	err = ctx.ShouldBindJSON(&photoEditData)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, response.ErrorResponse{Message: err.Error()})
		return
	}

	validate := validator.New()
	err = validate.Struct(photoEditData)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, response.ErrorResponse{Message: err.Error()})
		return
	}

	photoUpdate := model.Photo{}
	photoUpdate.ID = uint32(photoId)
	photoUpdate.Title = photoEditData.Title
	photoUpdate.Caption = photoEditData.Caption
	photoUpdate.PhotoUrl = photoEditData.PhotoUrl

	photoRes, err := p.svc.UpdatePhoto(ctx, photoUpdate)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, response.ErrorResponse{Message: err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, photoRes)
}

func (p *photoHandlerImpl) DeletePhoto(ctx *gin.Context) {
	photoId, err := strconv.Atoi(ctx.Param("id"))
	if photoId == 0 || err != nil {
		ctx.JSON(http.StatusBadRequest, response.ErrorResponse{Message: err.Error()})
		return
	}

	photo, err := p.svc.GetPhotoById(ctx, uint32(photoId))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, response.ErrorResponse{Message: err.Error()})
		return
	}

	if photo.ID == 0 {
		ctx.JSON(http.StatusNotFound, response.ErrorResponse{Message: "Photo did not exist"})
		return
	}

	userId, err := helper.GetUserIdFromGinCtx(ctx)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, response.ErrorResponse{Message: err.Error()})
		return
	}

	if userId != uint32(photo.UserId) {
		ctx.JSON(http.StatusUnauthorized, response.ErrorResponse{Message: "unauthorized to do this request"})
		return
	}

	err = p.svc.DeletePhoto(ctx, uint32(photoId))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, response.ErrorResponse{Message: err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, response.SuccessResponse{Message: "Your photo has been successfully deleted"})
}