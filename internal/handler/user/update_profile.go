package user

import (
	"net/http"
	"strconv"

	userDto "github.com/Mariano-SI/twitter-api/internal/dto/user"
	"github.com/Mariano-SI/twitter-api/internal/handler/httperror"
	"github.com/gin-gonic/gin"
)

func (h *Handler) UpdateProfile(c *gin.Context) {
	ctx := c.Request.Context()

	form, err := c.MultipartForm()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "invalid multipart form"})
		return
	}

	req := userDto.UpdateProfileDto{}

	if vals, ok := form.Value["description"]; ok {
		desc := vals[0]
		req.Description = &desc
	}

	if vals, ok := form.Value["remove_profile_picture"]; ok {
		b, err := strconv.ParseBool(vals[0])
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"message": "remove_profile_picture must be a boolean"})
			return
		}
		req.RemoveProfilePicture = &b
	}

	if files := form.File["profile_picture"]; len(files) > 0 {
		req.ProfilePicture = files[0]
	}

	if err := req.Validate(); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	userId := c.GetString("userId")

	response, err := h.userService.UpdateProfile(ctx, req, userId)
	if err != nil {
		status, msg := httperror.FromError(err)
		c.JSON(status, gin.H{"message": msg})
		return
	}

	c.JSON(http.StatusOK, response)
}
