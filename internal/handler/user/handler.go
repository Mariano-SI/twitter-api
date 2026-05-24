package user

import (
	"github.com/Mariano-SI/twitter-api/internal/service/user"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	api         gin.IRoutes
	userService user.UserService
}

func NewHandler(api gin.IRoutes, userService user.UserService) *Handler {
	return &Handler{api: api, userService: userService}
}

func (h *Handler) RouteList() {
	h.api.POST("/users", h.Register)
}
