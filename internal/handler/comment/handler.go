package comment

import (
	"github.com/Mariano-SI/twitter-api/internal/middleware"
	commentService "github.com/Mariano-SI/twitter-api/internal/service/comment"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	api            gin.IRouter
	commentService commentService.CommentService
}

func NewHandler(api gin.IRouter, commentService commentService.CommentService) *Handler {
	return &Handler{
		api:            api,
		commentService: commentService,
	}
}

func (h *Handler) RouteList(secretKey string) {
	posts := h.api.Group("/posts")
	posts.Use(middleware.AuthMiddleware(secretKey))
	posts.POST("/:id/comments", h.Create)
}
