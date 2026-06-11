package postlike

import (
	"github.com/Mariano-SI/twitter-api/internal/app"
	postRepository "github.com/Mariano-SI/twitter-api/internal/repository/post"
	postLikeRepository "github.com/Mariano-SI/twitter-api/internal/repository/post_like"
	postLikeService "github.com/Mariano-SI/twitter-api/internal/service/post_like"
	"github.com/gin-gonic/gin"
)

// Register wires the post_like repositories, service and handler, and
// registers its routes on the given router group.
func Register(rg gin.IRouter, deps *app.Deps) {
	postLikeRepo := postLikeRepository.NewRepository(deps.DB)
	postRepo := postRepository.NewRepository(deps.DB)

	svc := postLikeService.NewService(postLikeRepo, postRepo)

	NewHandler(rg, svc).RouteList(deps.Config.JwtSecret)
}
