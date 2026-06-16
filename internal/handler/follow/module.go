package follow

import (
	"github.com/Mariano-SI/twitter-api/internal/app"
	followRepository "github.com/Mariano-SI/twitter-api/internal/repository/follow"
	userRepository "github.com/Mariano-SI/twitter-api/internal/repository/user"
	followService "github.com/Mariano-SI/twitter-api/internal/service/follow"
	"github.com/gin-gonic/gin"
)

func Register(rg gin.IRouter, deps *app.Deps) {
	followRepo := followRepository.NewRepository(deps.DB)
	userRepo := userRepository.NewRepository(deps.DB)

	svc := followService.NewService(followRepo, userRepo)

	NewHandler(rg, svc).RouteList(deps.Config.JwtSecret)
}
