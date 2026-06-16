package post

import (
	"errors"
	"strings"
	"time"

	"github.com/google/uuid"
)

type GetPostByIdDto struct {
	Id string
}

func (d *GetPostByIdDto) Validate() error {
	if strings.TrimSpace(d.Id) == "" {
		return errors.New("id is required")
	}

	return uuid.Validate(d.Id)
}

type GetPostByIdCommentResponseDto struct {
	Id        string    `json:"id"`
	Username  string    `json:"username"`
	Content   string    `json:"content"`
	LikeCount int       `json:"like_count"`
	Images    []string  `json:"images"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type GetPostByIdResponseDto struct {
	Id        string                          `json:"id"`
	Username  string                          `json:"username"`                       
	Content   string                          `json:"content"`
	Images    []string                        `json:"images"`
	LikeCount int                             `json:"like_count"`
	Comments  []GetPostByIdCommentResponseDto `json:"comments"`
	CreatedAt time.Time                       `json:"created_at"`
	UpdatedAt time.Time                       `json:"updated_at"`
}
