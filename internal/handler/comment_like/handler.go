package commentlike

import (
	"github.com/Mariano-SI/twitter-api/internal/middleware"
	commentLikeService "github.com/Mariano-SI/twitter-api/internal/service/comment_like"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	api            gin.IRouter
	commentLikeService commentLikeService.CommentLikeService
}

func NewHandler(api gin.IRouter, commentLikeService commentLikeService.CommentLikeService) *Handler {
	return &Handler{
		api:            api,
		commentLikeService: commentLikeService,
	}
}

func (h *Handler) RouteList(secretKey string) {
	posts := h.api.Group("/comments")
	posts.Use(middleware.AuthMiddleware(secretKey))
	posts.POST("/:id/likes", h.LikeOrUnlike)
}
