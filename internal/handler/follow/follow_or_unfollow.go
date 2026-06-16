package follow

import (
	"net/http"

	followDto "github.com/Mariano-SI/twitter-api/internal/dto/follow"
	"github.com/Mariano-SI/twitter-api/internal/handler/httperror"
	"github.com/gin-gonic/gin"
)

func (h *Handler) FollowOrUnfollow(c *gin.Context) {
	ctx := c.Request.Context()

	req := followDto.FollowOrUnfollowDto{
		FollowedId: c.Param("id"),
	}

	if err := req.Validate(); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	followerId := c.GetString("userId")

	response, err := h.followService.FollowOrUnfollow(ctx, req, followerId)
	if err != nil {
		status, msg := httperror.FromError(err)
		c.JSON(status, gin.H{"message": msg})
		return
	}

	statusCode := http.StatusOK
	if response.Following {
		statusCode = http.StatusCreated
	}

	c.JSON(statusCode, response)
}
