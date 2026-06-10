package commentlike

import (
	"net/http"

	"github.com/Mariano-SI/twitter-api/internal/handler/httperror"
	"github.com/gin-gonic/gin"
	commentLikeDto "github.com/Mariano-SI/twitter-api/internal/dto/comment_like"
)

func (h *Handler) LikeOrUnlike(c *gin.Context) {
	ctx := c.Request.Context()

	req := commentLikeDto.LikeOrUnlikeCommentDto{
		CommentId: c.Param("id"),
	}

	if err := req.Validate(); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	userId := c.GetString("userId")

	response, err := h.commentLikeService.LikeOrUnlikeComment(ctx, req, userId)
	if err != nil {
		status, msg := httperror.FromError(err)
		c.JSON(status, gin.H{"message": msg})
		return
	}

	c.JSON(http.StatusOK, response)
}
