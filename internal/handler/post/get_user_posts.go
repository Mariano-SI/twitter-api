package post

import (
	"net/http"
	"strconv"

	postDto "github.com/Mariano-SI/twitter-api/internal/dto/post"
	"github.com/Mariano-SI/twitter-api/internal/handler/httperror"
	"github.com/gin-gonic/gin"
)

func (h *Handler) GetUserPosts(c *gin.Context) {
	ctx := c.Request.Context()

	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	limit, _ := strconv.Atoi(c.DefaultQuery("limit", "20"))

	req := postDto.GetMyPostsDto{
		Page:  page,
		Limit: limit,
	}

	userId := c.Param("id")

	response, err := h.postService.GetUserPosts(ctx, req, userId)
	if err != nil {
		status, msg := httperror.FromError(err)
		c.JSON(status, gin.H{"message": msg})
		return
	}

	c.JSON(http.StatusOK, response)
}
