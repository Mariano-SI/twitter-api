package user

import (
	"github.com/Mariano-SI/twitter-api/internal/service/user"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	api         gin.IRouter
	userService user.UserService
}

func NewHandler(api gin.IRouter, userService user.UserService) *Handler {
	return &Handler{api: api, userService: userService}
}

func (h *Handler) RouteList() {
	auth := h.api.Group("/auth")
	auth.POST("/register", h.Register)
	auth.POST("/login", h.Login)
}
