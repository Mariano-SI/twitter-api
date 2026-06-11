package post

import (
	"github.com/Mariano-SI/twitter-api/internal/app"
	postRepository "github.com/Mariano-SI/twitter-api/internal/repository/post"
	postImageRepository "github.com/Mariano-SI/twitter-api/internal/repository/post_image"
	postService "github.com/Mariano-SI/twitter-api/internal/service/post"
	"github.com/gin-gonic/gin"
)

// Register wires the post repositories, service and handler, and registers
// its routes on the given router group.
func Register(rg gin.IRouter, deps *app.Deps) {
	postRepo := postRepository.NewRepository(deps.DB)
	postImageRepo := postImageRepository.NewRepository(deps.DB)

	svc := postService.NewService(deps.Transactor, postRepo, postImageRepo, deps.Storage)

	NewHandler(rg, svc).RouteList(deps.Config.JwtSecret)
}
