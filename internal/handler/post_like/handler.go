package postlike

import (
	"github.com/Mariano-SI/twitter-api/internal/middleware"
	postLikeService "github.com/Mariano-SI/twitter-api/internal/service/post_like"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	api             gin.IRouter
	postLikeService postLikeService.PostLikeService
}

func NewHandler(api gin.IRouter, postLikeService postLikeService.PostLikeService) *Handler {
	return &Handler{
		api:             api,
		postLikeService: postLikeService,
	}
}

func (h *Handler) RouteList(secretKey string) {
	posts := h.api.Group("/posts")
	posts.Use(middleware.AuthMiddleware(secretKey))
	posts.POST("/:id/likes", h.LikeOrUnlike)
}
