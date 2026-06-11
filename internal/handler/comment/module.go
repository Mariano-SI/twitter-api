package comment

import (
	"github.com/Mariano-SI/twitter-api/internal/app"
	commentRepository "github.com/Mariano-SI/twitter-api/internal/repository/comment"
	commentImageRepository "github.com/Mariano-SI/twitter-api/internal/repository/comment_image"
	postRepository "github.com/Mariano-SI/twitter-api/internal/repository/post"
	commentService "github.com/Mariano-SI/twitter-api/internal/service/comment"
	"github.com/gin-gonic/gin"
)

// Register wires the comment repositories, service and handler, and registers
// its routes on the given router group.
func Register(rg gin.IRouter, deps *app.Deps) {
	commentRepo := commentRepository.NewRepository(deps.DB)
	commentImageRepo := commentImageRepository.NewRepository(deps.DB)
	postRepo := postRepository.NewRepository(deps.DB)

	svc := commentService.NewService(deps.Transactor, commentRepo, commentImageRepo, postRepo, deps.Storage)

	NewHandler(rg, svc).RouteList(deps.Config.JwtSecret)
}
