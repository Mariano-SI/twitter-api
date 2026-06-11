package commentlike

import (
	"github.com/Mariano-SI/twitter-api/internal/app"
	commentRepository "github.com/Mariano-SI/twitter-api/internal/repository/comment"
	commentLikeRepository "github.com/Mariano-SI/twitter-api/internal/repository/comment_like"
	commentLikeService "github.com/Mariano-SI/twitter-api/internal/service/comment_like"
	"github.com/gin-gonic/gin"
)

// Register wires the comment_like repositories, service and handler, and
// registers its routes on the given router group.
func Register(rg gin.IRouter, deps *app.Deps) {
	commentLikeRepo := commentLikeRepository.NewRepository(deps.DB)
	commentRepo := commentRepository.NewRepository(deps.DB)

	svc := commentLikeService.NewService(commentLikeRepo, commentRepo)

	NewHandler(rg, svc).RouteList(deps.Config.JwtSecret)
}
