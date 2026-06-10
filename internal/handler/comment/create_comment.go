package comment

import (
	"net/http"

	commentDto "github.com/Mariano-SI/twitter-api/internal/dto/comment"
	"github.com/Mariano-SI/twitter-api/internal/handler/httperror"
	"github.com/gin-gonic/gin"
)

func (h *Handler) Create(c *gin.Context) {
	ctx := c.Request.Context()

	form, err := c.MultipartForm()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "invalid multipart form"})
		return
	}

	req := commentDto.CreateCommentDto{
		PostId:  c.Param("id"),
		Content: c.PostForm("content"),
		Images:  form.File["images"],
	}

	if err := req.Validate(); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	userId := c.GetString("userId")

	response, err := h.commentService.CreateComment(ctx, req, userId)
	if err != nil {
		status, msg := httperror.FromError(err)
		c.JSON(status, gin.H{"message": msg})
		return
	}

	c.JSON(http.StatusCreated, response)
}
