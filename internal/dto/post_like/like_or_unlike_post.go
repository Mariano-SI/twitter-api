package postlike

import (
	"errors"
	"strings"

	"github.com/google/uuid"
)

type LikeOrUnlikePostDto struct {
	PostId string `json:"post_id"`
}

func (l *LikeOrUnlikePostDto) Validate() error {
	if strings.TrimSpace(l.PostId) == "" {
		return errors.New("post_id is required")
	}

	return uuid.Validate(l.PostId)
}

type LikeOrUnlikePostResponseDto struct {
	Message string `json:"message"`
}
