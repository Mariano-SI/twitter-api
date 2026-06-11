package user

import (
	"github.com/Mariano-SI/twitter-api/internal/app"
	refreshTokenRepository "github.com/Mariano-SI/twitter-api/internal/repository/refresh_token"
	userRepository "github.com/Mariano-SI/twitter-api/internal/repository/user"
	userService "github.com/Mariano-SI/twitter-api/internal/service/user"
	"github.com/gin-gonic/gin"
)

// Register wires the user repositories, service and handler, and registers
// its routes on the given router group.
func Register(rg gin.IRouter, deps *app.Deps) {
	userRepo := userRepository.NewRepository(deps.DB)
	refreshRepo := refreshTokenRepository.NewRepository(deps.DB)

	svc := userService.NewService(deps.Config, userRepo, refreshRepo)

	NewHandler(rg, svc).RouteList(deps.Config.JwtSecret)
}
