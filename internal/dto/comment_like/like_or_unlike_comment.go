package commentlike

import (
	"errors"
	"strings"

	"github.com/google/uuid"
)

type LikeOrUnlikeCommentDto struct {
	CommentId string
}

func (c *LikeOrUnlikeCommentDto) Validate() error {
	if strings.TrimSpace(c.CommentId) ==""{
		return errors.New("comment_id is required")
	}

	return uuid.Validate(c.CommentId)
}

type LikeOrUnlikeCommentResponseDto struct{
	Message string `json:"message"`
}