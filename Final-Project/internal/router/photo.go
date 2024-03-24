package router

import (
	"github.com/gin-gonic/gin"
	"github.com/rahmatadlin/Go-Digitalent-2024/Final-Project/internal/handler"
	"github.com/rahmatadlin/Go-Digitalent-2024/Final-Project/internal/middleware"
)

type PhotoRouter interface {
	Mount()
}

type photoRouterImpl struct {
	v       *gin.RouterGroup
	handler handler.PhotoHandler
	auth    middleware.Authorization
}

func NewPhotoRouter(v *gin.RouterGroup, handler handler.PhotoHandler, auth middleware.Authorization) PhotoRouter {
	return &photoRouterImpl{v: v, handler: handler, auth: auth}
}

func (p *photoRouterImpl) Mount() {
	p.v.Use(p.auth.CheckAuth)
	p.v.POST("", p.handler.PostPhoto)
	p.v.GET("", p.handler.GetAllPhotosByUserId)
	p.v.PUT("/:id", p.handler.UpdatePhoto)
	p.v.DELETE("/:id", p.handler.DeletePhoto)
}