package post

import (
	"github.com/Mariano-SI/twitter-api/internal/middleware"
	"github.com/Mariano-SI/twitter-api/internal/service/post"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	api         gin.IRouter
	postService post.PostService
}

func NewHandler(api gin.IRouter, postService post.PostService) *Handler {
	return &Handler{
		api:         api,
		postService: postService,
	}
}

func (h *Handler) RouteList(secretKey string) {
	posts := h.api.Group("/posts")
	posts.Use(middleware.AuthMiddleware(secretKey))
	posts.POST("/", h.Create)
	posts.DELETE("/:id", h.Delete)
	posts.GET("/:id", h.GetById)
}
