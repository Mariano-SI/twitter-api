package post

import (
	"errors"
	"strings"
)

type CreatePostDto struct {
	Content string `json:"content"`
}

func (c *CreatePostDto) Validate() error {
	if strings.TrimSpace(c.Content) == "" {
		return errors.New("content is required")
	}
	if len(c.Content) > 280 {
		return errors.New("content exceeds 280 characters")
	}

	return nil
}

type CreatePostResponseDto struct {
	Id string `json:"id"`
}
