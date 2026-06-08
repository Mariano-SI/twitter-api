package post

import (
	"net/http"

	postDto "github.com/Mariano-SI/twitter-api/internal/dto/post"
	"github.com/Mariano-SI/twitter-api/internal/handler/httperror"
	"github.com/gin-gonic/gin"
)

func (h *Handler) Delete(c *gin.Context) {
	ctx := c.Request.Context()

	req := postDto.DeletePostDto{
		Id: c.Param("id"),
	}

	if err := req.Validate(); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	userId := c.GetString("userId")

	response, err := h.postService.Delete(ctx, req, userId)
	if err != nil {
		status, msg := httperror.FromError(err)
		c.JSON(status, gin.H{"message": msg})
		return
	}

	c.JSON(http.StatusNoContent, response)
}
