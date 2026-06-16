package post

import (
	"github.com/Mariano-SI/twitter-api/internal/app"
	commentRepository "github.com/Mariano-SI/twitter-api/internal/repository/comment"
	commentImageRepository "github.com/Mariano-SI/twitter-api/internal/repository/comment_image"
	postRepository "github.com/Mariano-SI/twitter-api/internal/repository/post"
	postImageRepository "github.com/Mariano-SI/twitter-api/internal/repository/post_image"
	postService "github.com/Mariano-SI/twitter-api/internal/service/post"
	"github.com/gin-gonic/gin"
)

func Register(rg gin.IRouter, deps *app.Deps) {
	postRepo := postRepository.NewRepository(deps.DB)
	postImageRepo := postImageRepository.NewRepository(deps.DB)
	commentRepo := commentRepository.NewRepository(deps.DB)
	commentImageRepo := commentImageRepository.NewRepository(deps.DB)

	svc := postService.NewService(deps.Transactor, postRepo, postImageRepo, commentRepo, commentImageRepo, deps.Storage)

	NewHandler(rg, svc).RouteList(deps.Config.JwtSecret)
}
