package handler

import (
	"net/http"
	"strconv"

	"github.com/rahmatadlin/Go-Digitalent-2024/Final-Project/internal/model"
	"github.com/rahmatadlin/Go-Digitalent-2024/Final-Project/internal/service"
	"github.com/rahmatadlin/Go-Digitalent-2024/Final-Project/pkg/response"
	"github.com/rahmatadlin/Go-Digitalent-2024/Final-Project/pkg/helper"
	
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator"
)

type SocialMediaHandler interface {
	PostSocialMedia(ctx *gin.Context)
	GetAllSocialMediasByUserId(ctx *gin.Context)
	UpdateSocialMedia(ctx *gin.Context)
	DeleteSocialMedia(ctx *gin.Context)
}

type socialMediaHandlerImpl struct {
	svc service.SocialMediaService
}

func NewSocialMediaHandler(svc service.SocialMediaService) SocialMediaHandler {
	return &socialMediaHandlerImpl{svc: svc}
}

func (s *socialMediaHandlerImpl) PostSocialMedia(ctx *gin.Context) {
	userId, err := helper.GetUserIdFromGinCtx(ctx)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, response.ErrorResponse{Message: err.Error()})
		return
	}

	newSocial := model.NewSocialMedia{}
	err = ctx.ShouldBindJSON(&newSocial)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, response.ErrorResponse{Message: err.Error()})
		return
	}

	validate := validator.New()
	err = validate.Struct(newSocial)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, response.ErrorResponse{Message: err.Error()})
		return
	}

	socialMediaRes, err := s.svc.PostSocial(ctx, userId, newSocial)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, response.ErrorResponse{Message: err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, socialMediaRes)
}

func (s *socialMediaHandlerImpl) GetAllSocialMediasByUserId(ctx *gin.Context) {
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

	socials, err := s.svc.GetAllSocialMediasByUserId(ctx, uint32(userId))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, response.ErrorResponse{Message: err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, socials)
}

func (s *socialMediaHandlerImpl) UpdateSocialMedia(ctx *gin.Context) {
	socialId, err := strconv.Atoi(ctx.Param("id"))
	if socialId == 0 || err != nil {
		ctx.JSON(http.StatusBadRequest, response.ErrorResponse{Message: err.Error()})
		return
	}

	social, err := s.svc.GetSocialById(ctx, uint32(socialId))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, response.ErrorResponse{Message: err.Error()})
		return
	}

	if social.ID == 0 {
		ctx.JSON(http.StatusNotFound, response.ErrorResponse{Message: "User Social Media data did not exist"})
		return
	}

	userId, err := helper.GetUserIdFromGinCtx(ctx)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, response.ErrorResponse{Message: err.Error()})
		return
	}

	if userId != uint32(social.UserId) {
		ctx.JSON(http.StatusUnauthorized, response.ErrorResponse{Message: "unauthorized to do this request"})
		return
	}

	socialUpdateData := model.NewSocialMedia{}
	err = ctx.ShouldBindJSON(&socialUpdateData)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, response.ErrorResponse{Message: err.Error()})
		return
	}

	validate := validator.New()
	err = validate.Struct(socialUpdateData)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, response.ErrorResponse{Message: err.Error()})
		return
	}

	social.Name = socialUpdateData.Name
	social.SocialMediaUrl = socialUpdateData.SocialMediaUrl

	socialMediaRes, err := s.svc.UpdateSocial(ctx, *social)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, response.ErrorResponse{Message: err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, socialMediaRes)
}

func (s *socialMediaHandlerImpl) DeleteSocialMedia(ctx *gin.Context) {
	socialId, err := strconv.Atoi(ctx.Param("id"))
	if socialId == 0 || err != nil {
		ctx.JSON(http.StatusBadRequest, response.ErrorResponse{Message: err.Error()})
		return
	}

	social, err := s.svc.GetSocialById(ctx, uint32(socialId))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, response.ErrorResponse{Message: err.Error()})
		return
	}

	if social.ID == 0 {
		ctx.JSON(http.StatusNotFound, response.ErrorResponse{Message: "User Social Media data did not exist"})
		return
	}

	userId, err := helper.GetUserIdFromGinCtx(ctx)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, response.ErrorResponse{Message: err.Error()})
		return
	}

	if userId != uint32(social.UserId) {
		ctx.JSON(http.StatusUnauthorized, response.ErrorResponse{Message: "unauthorized to do this request"})
		return
	}

	err = s.svc.DeleteSocial(ctx, uint32(socialId))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, response.ErrorResponse{Message: err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, response.SuccessResponse{Message: "Your social media has been successfully deleted"})
}