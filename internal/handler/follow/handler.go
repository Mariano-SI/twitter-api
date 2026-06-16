package follow

import (
	"github.com/Mariano-SI/twitter-api/internal/middleware"
	followService "github.com/Mariano-SI/twitter-api/internal/service/follow"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	api           gin.IRouter
	followService followService.FollowService
}

func NewHandler(api gin.IRouter, followService followService.FollowService) *Handler {
	return &Handler{
		api:           api,
		followService: followService,
	}
}

func (h *Handler) RouteList(secretKey string) {
	users := h.api.Group("/users")
	users.Use(middleware.AuthMiddleware(secretKey))
	users.POST("/:id/follow", h.FollowOrUnfollow)
}
