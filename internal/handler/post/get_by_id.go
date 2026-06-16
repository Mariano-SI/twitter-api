package post

import (
	"net/http"

	postDto "github.com/Mariano-SI/twitter-api/internal/dto/post"
	"github.com/Mariano-SI/twitter-api/internal/handler/httperror"
	"github.com/gin-gonic/gin"
)

func (h *Handler) GetById(c *gin.Context) {
	ctx := c.Request.Context()

	req := postDto.GetPostByIdDto{
		Id: c.Param("id"),
	}

	if err := req.Validate(); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	response, err := h.postService.GetById(ctx, req)
	if err != nil {
		status, msg := httperror.FromError(err)
		c.JSON(status, gin.H{"message": msg})
		return
	}

	c.JSON(http.StatusOK, response)
}
