package post

import (
	"net/http"

	"github.com/Mariano-SI/twitter-api/internal/dto/post"
	"github.com/Mariano-SI/twitter-api/internal/handler/httperror"
	"github.com/gin-gonic/gin"
)

func (h *Handler) Create(c *gin.Context) {
	ctx := c.Request.Context()

	form, err := c.MultipartForm()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "invalid multipart form",
		})
		return
	}

	req := post.CreatePostDto{
		Content: c.PostForm("content"),
		Images:  form.File["images"],
	}

	if err := req.Validate(); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}

	userId := c.GetString("userId")

	post, err := h.postService.Create(ctx, req, userId)

	if err != nil {
		status, msg := httperror.FromError(err)
		c.JSON(status, gin.H{
			"message": msg,
		})
		return
	}

	c.JSON(http.StatusCreated, post)
}
