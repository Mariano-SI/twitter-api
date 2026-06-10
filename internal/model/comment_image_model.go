package model

import (
	"time"

	"github.com/google/uuid"
)

type CommentImageModel struct {
	ID        string
	CommentID string
	ImageUrl  string
	Position  uint8
	CreatedAt time.Time
}

func NewCommentImageModel(comment_id, image_url string, position uint8) *CommentImageModel {
	return &CommentImageModel{
		ID:        uuid.NewString(),
		CommentID: comment_id,
		ImageUrl:  image_url,
		Position:  position,
		CreatedAt: time.Now(),
	}
}
