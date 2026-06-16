package user

import (
	"github.com/Mariano-SI/twitter-api/internal/app"
	refreshTokenRepository "github.com/Mariano-SI/twitter-api/internal/repository/refresh_token"
	userRepository "github.com/Mariano-SI/twitter-api/internal/repository/user"
	userService "github.com/Mariano-SI/twitter-api/internal/service/user"
	"github.com/gin-gonic/gin"
)

func Register(rg gin.IRouter, deps *app.Deps) {
	userRepo := userRepository.NewRepository(deps.DB)
	refreshRepo := refreshTokenRepository.NewRepository(deps.DB)

	svc := userService.NewService(deps.Config, userRepo, refreshRepo, deps.Storage)

	NewHandler(rg, svc).RouteList(deps.Config.JwtSecret)
}
