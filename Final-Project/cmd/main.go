package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	_ "github.com/rahmatadlin/Go-Digitalent-2024/Final-Project/cmd/docs"
	"github.com/rahmatadlin/Go-Digitalent-2024/Final-Project/internal/handler"
	"github.com/rahmatadlin/Go-Digitalent-2024/Final-Project/internal/infrastructure"
	"github.com/rahmatadlin/Go-Digitalent-2024/Final-Project/internal/middleware"
	"github.com/rahmatadlin/Go-Digitalent-2024/Final-Project/internal/repository"
	"github.com/rahmatadlin/Go-Digitalent-2024/Final-Project/internal/router"
	"github.com/rahmatadlin/Go-Digitalent-2024/Final-Project/internal/service"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)


func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalln("Cannot load env: ", err)
	}

	g := gin.Default()

	gorm := infrastructure.NewGormPostgres()

	userRouteGroup := g.Group("/v1/users")
	userRepo := repository.NewUserRepository(gorm)
	userService := service.NewUserService(userRepo)
	userHandler := handler.NewUserHandler(userService)

	auth := middleware.NewAuthorization(userService)

	userRouter := router.NewUserRouter(userRouteGroup, userHandler, auth)
	userRouter.Mount()

	photoRouteGroup := g.Group("/v1/photos")
	photoRepo := repository.NewPhotoRepository(gorm)
	photoService := service.NewPhotoService(photoRepo)
	photoHandler := handler.NewPhotoHandler(photoService)
	photoRouter := router.NewPhotoRouter(photoRouteGroup, photoHandler, auth)
	photoRouter.Mount()

	commentRouteGroup := g.Group("/v1/comments")
	commentRepo := repository.NewCommentRepository(gorm)
	commentService := service.NewCommentService(commentRepo)
	commentHandler := handler.NewCommentHandler(commentService, photoService)
	commentRouter := router.NewCommentRouter(commentRouteGroup, commentHandler, auth)
	commentRouter.Mount()

	socialMediaRouteGroup := g.Group("/v1/socialmedias")
	socialMediaRepo := repository.NewSocialMediaRepository(gorm)
	socialMediaService := service.NewSocialMediaService(socialMediaRepo)
	socialMediaHandler := handler.NewSocialMediaHandler(socialMediaService)
	socialMediaRouter := router.NewSocialMediaRouter(socialMediaRouteGroup, socialMediaHandler, auth)
	socialMediaRouter.Mount()

	g.GET("/ping", func(ctx *gin.Context) {
		ctx.Writer.Write([]byte("Server online"))
	})

	g.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	g.Run("127.0.0.1:3000")
}