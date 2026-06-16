package user

import (
	"github.com/Mariano-SI/twitter-api/internal/middleware"
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

func (h *Handler) RouteList(secretKey string) {
	auth := h.api.Group("/auth")
	auth.POST("/register", h.Register)
	auth.POST("/login", h.Login)

	refreshRoute := h.api.Group("/auth")
	refreshRoute.Use(middleware.AuthRefreshTokenMiddleware(secretKey))
	refreshRoute.POST("/refresh", h.RefreshToken)

	users := h.api.Group("/users")
	users.Use(middleware.AuthMiddleware(secretKey))
	users.GET("/me", h.GetMyProfile)
	users.PATCH("/me", h.UpdateProfile)
	users.GET("/:id", h.GetProfile)
}
