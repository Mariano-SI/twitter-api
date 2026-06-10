package model

import (
	"time"

	"github.com/google/uuid"
)

type CommentLikeModel struct {
	ID        string
	CommentID string
	UserID    string
	CreatedAt time.Time
}

func NewCommentLikeModel(commentId, userId string) *CommentLikeModel {
	return &CommentLikeModel{
		ID:        uuid.NewString(),
		CommentID: commentId,
		UserID:    userId,
		CreatedAt: time.Now(),
	}
}
