package user

import (
	"net/http"

	"github.com/Mariano-SI/twitter-api/internal/handler/httperror"
	"github.com/gin-gonic/gin"
)

func (h *Handler) GetMyProfile(c *gin.Context) {
	h.getProfile(c, c.GetString("userId"))
}

func (h *Handler) GetProfile(c *gin.Context) {
	h.getProfile(c, c.Param("id"))
}

func (h *Handler) getProfile(c *gin.Context, userId string) {
	response, err := h.userService.GetProfile(c.Request.Context(), userId)
	if err != nil {
		status, msg := httperror.FromError(err)
		c.JSON(status, gin.H{"message": msg})
		return
	}

	c.JSON(http.StatusOK, response)
}
